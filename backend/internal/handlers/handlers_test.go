package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mihir-panchal/vibe-portfolio/internal/cache"
	"github.com/mihir-panchal/vibe-portfolio/internal/services"
	"time"
)

// ─── Stub implementations for testing ───────────────────────────────

// stubImageGenerator always returns a fixed result for deterministic testing.
type stubImageGenerator struct {
	result services.ImageResult
	err    error
}

func (s *stubImageGenerator) GenerateCatArt(ctx context.Context) (services.ImageResult, error) {
	return s.result, s.err
}

// stubMusicGenerator always returns a fixed result for deterministic testing.
type stubMusicGenerator struct {
	result services.TrackResult
	err    error
}

func (s *stubMusicGenerator) GenerateTrack(ctx context.Context, genre string) (services.TrackResult, error) {
	return s.result, s.err
}

// ─── Art Handler Tests ──────────────────────────────────────────────

func TestArtHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		generator      *stubImageGenerator
		preCacheResult *services.ImageResult // if non-nil, pre-populate cache
		wantStatus     int
		wantURL        string
	}{
		{
			name:   "GET returns generated art on cache miss",
			method: http.MethodGet,
			generator: &stubImageGenerator{
				result: services.ImageResult{
					URL:     "https://example.com/cat.png",
					Caption: "A cute cat",
				},
			},
			wantStatus: http.StatusOK,
			wantURL:    "https://example.com/cat.png",
		},
		{
			name:   "GET returns cached art on cache hit",
			method: http.MethodGet,
			generator: &stubImageGenerator{
				result: services.ImageResult{
					URL:     "https://example.com/should-not-be-called.png",
					Caption: "Should not appear",
				},
			},
			preCacheResult: &services.ImageResult{
				URL:     "https://example.com/cached-cat.png",
				Caption: "Cached cat",
			},
			wantStatus: http.StatusOK,
			wantURL:    "https://example.com/cached-cat.png",
		},
		{
			name:       "POST returns 405 Method Not Allowed",
			method:     http.MethodPost,
			generator:  &stubImageGenerator{},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "GET returns 500 on generator error",
			method: http.MethodGet,
			generator: &stubImageGenerator{
				err: context.DeadlineExceeded,
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			artCache := cache.New(24 * time.Hour)

			// Pre-populate cache if test requires it
			if tt.preCacheResult != nil {
				artCache.Set(artCacheKey, *tt.preCacheResult)
			}

			handler := NewArtHandler(tt.generator, artCache)

			req := httptest.NewRequest(tt.method, "/api/art-of-the-day", nil)
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}

			if tt.wantURL != "" {
				var result services.ImageResult
				if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if result.URL != tt.wantURL {
					t.Errorf("url = %q, want %q", result.URL, tt.wantURL)
				}
			}
		})
	}
}

// ─── Music Handler Tests ────────────────────────────────────────────

func TestMusicHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		body       interface{}
		generator  *stubMusicGenerator
		wantStatus int
		wantTrack  string
	}{
		{
			name:   "POST with valid genre returns track",
			method: http.MethodPost,
			body:   map[string]string{"genre": "jazz"},
			generator: &stubMusicGenerator{
				result: services.TrackResult{
					TrackName: "Velvet Espresso",
					Artist:    "The Blue Note Cats",
					Genre:     "Jazz",
				},
			},
			wantStatus: http.StatusOK,
			wantTrack:  "Velvet Espresso",
		},
		{
			name:       "GET returns 405 Method Not Allowed",
			method:     http.MethodGet,
			generator:  &stubMusicGenerator{},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "POST with empty genre returns 400",
			method:     http.MethodPost,
			body:       map[string]string{"genre": ""},
			generator:  &stubMusicGenerator{},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "POST with invalid JSON returns 400",
			method:     http.MethodPost,
			body:       nil, // will send empty body
			generator:  &stubMusicGenerator{},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:   "POST with unknown genre still returns a track",
			method: http.MethodPost,
			body:   map[string]string{"genre": "polka"},
			generator: &stubMusicGenerator{
				result: services.TrackResult{
					TrackName: "Unknown Frequency",
					Artist:    "Mystery Cat",
					Genre:     "polka",
				},
			},
			wantStatus: http.StatusOK,
			wantTrack:  "Unknown Frequency",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewMusicHandler(tt.generator)

			var body *bytes.Buffer
			if tt.body != nil {
				b, _ := json.Marshal(tt.body)
				body = bytes.NewBuffer(b)
			} else {
				body = bytes.NewBuffer(nil)
			}

			req := httptest.NewRequest(tt.method, "/api/generate-music", body)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}

			if tt.wantTrack != "" {
				var result services.TrackResult
				if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if result.TrackName != tt.wantTrack {
					t.Errorf("trackName = %q, want %q", result.TrackName, tt.wantTrack)
				}
			}
		})
	}
}
