package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// SpotifyClient interacts with the Spotify Web API
type SpotifyClient struct {
	clientID     string
	clientSecret string
	
	tokenMu      sync.RWMutex
	accessToken  string
	tokenExpires time.Time
	
	httpClient   *http.Client
}

// NewSpotifyClient creates a new Spotify API client
func NewSpotifyClient(clientID, clientSecret string) *SpotifyClient {
	return &SpotifyClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   &http.Client{Timeout: 10 * time.Second},
	}
}

// getAccessToken returns a valid access token, refreshing if necessary
func (s *SpotifyClient) getAccessToken() (string, error) {
	s.tokenMu.RLock()
	if s.accessToken != "" && time.Now().Before(s.tokenExpires) {
		token := s.accessToken
		s.tokenMu.RUnlock()
		return token, nil
	}
	s.tokenMu.RUnlock()

	s.tokenMu.Lock()
	defer s.tokenMu.Unlock()
	
	// Double check after acquiring write lock
	if s.accessToken != "" && time.Now().Before(s.tokenExpires) {
		return s.accessToken, nil
	}

	authData := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", s.clientID, s.clientSecret)))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+authData)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get token, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	s.accessToken = result.AccessToken
	// Keep a 1-minute buffer before expiration
	s.tokenExpires = time.Now().Add(time.Duration(result.ExpiresIn-60) * time.Second)

	return s.accessToken, nil
}

// SearchTrack searches for a track by query and returns its Spotify URL
func (s *SpotifyClient) SearchTrack(query string) (string, error) {
	token, err := s.getAccessToken()
	if err != nil {
		return "", fmt.Errorf("failed to get access token: %w", err)
	}

	searchURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=1", url.QueryEscape(query))
	
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("search failed, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		Tracks struct {
			Items []struct {
				ExternalURLs struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
			} `json:"items"`
		} `json:"tracks"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Tracks.Items) == 0 {
		return "", fmt.Errorf("no tracks found for query: %s", query)
	}

	return result.Tracks.Items[0].ExternalURLs.Spotify, nil
}
