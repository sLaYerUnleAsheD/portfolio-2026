// Package services — gemini_music.go implements the MusicGenerator interface
// using Google's Gemini API for AI-powered creative track name generation.
//
// Since Gemini generates text (not audio), this produces creative, unique
// track names and artist names based on the selected genre — giving each
// user click a fresh, AI-generated result.
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"google.golang.org/genai"
)

// GeminiMusicGenerator implements MusicGenerator using Google's Gemini text API.
type GeminiMusicGenerator struct {
	client *genai.Client
}

// NewGeminiMusicGenerator creates a new GeminiMusicGenerator with the provided API key.
func NewGeminiMusicGenerator(apiKey string) (*GeminiMusicGenerator, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &GeminiMusicGenerator{client: client}, nil
}

// GenerateTrack uses Gemini to create a unique, creative track name and artist
// for the given genre. Returns a TrackResult with the generated metadata.
func (g *GeminiMusicGenerator) GenerateTrack(ctx context.Context, genre string) (TrackResult, error) {
	// Format genre label for display
	genreLabels := map[string]string{
		"lofi":       "Lo-Fi",
		"ambient":    "Ambient",
		"jazz":       "Jazz",
		"chill":      "Chill",
		"classical":  "Classical",
		"electronic": "Electronic",
	}
	label := genreLabels[genre]
	if label == "" {
		label = genre
	}

	prompt := fmt.Sprintf(`Generate a unique, creative music track for the "%s" genre. 
Return ONLY a JSON object with exactly these fields, no markdown formatting, no code blocks:
{"trackName": "creative track name", "artist": "creative artist/band name", "genre": "%s"}

Requirements:
- The track name should be evocative, poetic, and match the %s genre mood
- The artist name should sound like a real indie artist or band
- Be creative and unique — never repeat common names
- Return ONLY the JSON, nothing else`, label, label, label)

	parts := []*genai.Part{
		{Text: prompt},
	}

	log.Printf("[gemini-music] generating track for genre: %s", genre)

	result, err := g.client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		[]*genai.Content{{Parts: parts}},
		nil,
	)
	if err != nil {
		return TrackResult{}, fmt.Errorf("Gemini API call failed: %w", err)
	}

	// Extract the text response
	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return TrackResult{}, fmt.Errorf("empty response from Gemini API")
	}

	responseText := result.Candidates[0].Content.Parts[0].Text
	// Clean up potential markdown code block wrapping
	responseText = strings.TrimSpace(responseText)
	responseText = strings.TrimPrefix(responseText, "```json")
	responseText = strings.TrimPrefix(responseText, "```")
	responseText = strings.TrimSuffix(responseText, "```")
	responseText = strings.TrimSpace(responseText)

	log.Printf("[gemini-music] raw response: %s", responseText)

	// Parse the JSON response
	var track TrackResult
	if err := json.Unmarshal([]byte(responseText), &track); err != nil {
		log.Printf("[gemini-music] failed to parse response, falling back: %v", err)
		// Fallback: use the raw text as the track name
		return TrackResult{
			TrackName: responseText,
			Artist:    "AI Composer",
			Genre:     label,
		}, nil
	}

	// Ensure genre label is properly formatted
	track.Genre = label

	return track, nil
}
