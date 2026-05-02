// Package services — music.go defines the MusicGenerator interface and
// its mock implementation for AI music generation.
//
// Implement MusicGenerator to integrate with real music generation APIs.
package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// TrackResult represents the response from a music generation call.
type TrackResult struct {
	TrackName string `json:"trackName"`
	Artist    string `json:"artist"`
	Genre     string `json:"genre"`
}

// MusicGenerator is the interface for AI music generation services.
// Implement this to integrate with real providers.
type MusicGenerator interface {
	GenerateTrack(ctx context.Context, genre string) (TrackResult, error)
}

// ─── Mock Track Database ────────────────────────────────────────────

// mockTracks maps genre IDs to curated track lists that give the app
// a polished feel even without a real AI backend.
var mockTracks = map[string][]struct {
	TrackName string
	Artist    string
}{
	"lofi": {
		{"Midnight Coffee Drip", "Sleepy Beans"},
		{"Rainy Window Study", "Lo-Fi Cat"},
		{"Soft Focus Sunrise", "Mellow Waves"},
		{"Paper Crane Dreams", "Chill Paws"},
		{"Late Night Ramen", "Cozy Beats"},
	},
	"ambient": {
		{"Stellar Drift", "Cosmo Fauna"},
		{"Morning Fog Waltz", "Ethereal Tides"},
		{"Glacier Whisper", "Arctic Bloom"},
		{"Deep Forest Hymn", "Moss & Fern"},
		{"Ocean Floor Lullaby", "Abyssal Light"},
	},
	"jazz": {
		{"Velvet Espresso", "The Blue Note Cats"},
		{"Swing Through Tokyo", "Neon Quartet"},
		{"Moonlit Serenade", "Charlie Whiskers"},
		{"Smoky Lounge Dreams", "The Mellow Four"},
		{"Bebop Sunrise", "Jazz Paws Trio"},
	},
	"chill": {
		{"Sunbeam Siesta", "Warm Drift"},
		{"Golden Hour Breeze", "Horizon Haze"},
		{"Lazy Sunday Clouds", "Chill Garden"},
		{"Pastel Daydream", "Soft Spectrum"},
		{"Hammock Lullaby", "Breeze Collective"},
	},
	"classical": {
		{"Sonata for a Curious Cat", "A. Vivalpaw"},
		{"Nocturne in Cream Minor", "F. Chopin Cat"},
		{"The Four Seasons: Spring (Cat Remix)", "Vivaldi Ensemble"},
		{"Moonlight Purrnata", "L. van Beethoven Cat"},
		{"Canon in Paw Major", "Pachelbel's Kittens"},
	},
	"electronic": {
		{"Neon Pulse", "Synthwave Cat"},
		{"Digital Sunrise", "Pixel Drift"},
		{"Circuit Bloom", "Volt & Whiskers"},
		{"Hyperspace Groove", "Quantum Beats"},
		{"Binary Lullaby", "Code Wave"},
	},
}

// fallbackTracks are returned when the requested genre doesn't match any known genre.
var fallbackTracks = []struct {
	TrackName string
	Artist    string
}{
	{"Unknown Frequency", "Mystery Cat"},
	{"Random Melody #42", "AI Composer"},
	{"Untitled Vibe", "The Unknowns"},
}

// MockMusicGenerator is a mock implementation of MusicGenerator.
type MockMusicGenerator struct{}

// NewMockMusicGenerator creates a new MockMusicGenerator.
func NewMockMusicGenerator() *MockMusicGenerator {
	return &MockMusicGenerator{}
}

// GenerateTrack returns a random mock track for the given genre.
// If the genre is unknown, a fallback track is returned.
func (m *MockMusicGenerator) GenerateTrack(ctx context.Context, genre string) (TrackResult, error) {
	tracks, ok := mockTracks[genre]
	if !ok {
		// Use fallback for unknown genres
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		fb := fallbackTracks[r.Intn(len(fallbackTracks))]
		return TrackResult{
			TrackName: fb.TrackName,
			Artist:    fb.Artist,
			Genre:     genre,
		}, nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selected := tracks[r.Intn(len(tracks))]

	// Format the genre label nicely
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

	return TrackResult{
		TrackName: selected.TrackName,
		Artist:    selected.Artist,
		Genre:     fmt.Sprintf("%s", label),
	}, nil
}
