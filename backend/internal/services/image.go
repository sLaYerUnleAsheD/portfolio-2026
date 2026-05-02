// Package services defines the core business interfaces and their mock
// implementations for external AI integrations.
//
// ImageGenerator abstracts the image generation API (e.g., Nano Banana).
// Implement this interface to plug in real AI image generation.
package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// ImageResult represents the response from an image generation call.
type ImageResult struct {
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

// ImageGenerator is the interface for AI image generation services.
// Implement this to integrate with real providers like Nano Banana.
type ImageGenerator interface {
	GenerateCatArt(ctx context.Context) (ImageResult, error)
}

// ─── Mock Implementation ────────────────────────────────────────────

// mockCatImages provides curated placeholder cat art URLs and captions.
// These are used by the mock implementation so the app works out of the box.
var mockCatImages = []struct {
	URL     string
	Caption string
}{
	{
		URL:     "https://placecats.com/millie/400/400",
		Caption: "A serene tabby cat lounging in golden afternoon light",
	},
	{
		URL:     "https://placecats.com/neo/400/400",
		Caption: "A curious kitten exploring a garden of wildflowers",
	},
	{
		URL:     "https://placecats.com/bella/400/400",
		Caption: "A majestic cat sitting regally on a windowsill",
	},
	{
		URL:     "https://placecats.com/400/400",
		Caption: "A playful cat chasing butterflies in a meadow",
	},
	{
		URL:     "https://placecats.com/louie/400/400",
		Caption: "A cozy cat curled up with a warm cup of cocoa nearby",
	},
}

// MockImageGenerator is a mock implementation of ImageGenerator that returns
// a deterministic image based on the current date (one per day).
type MockImageGenerator struct{}

// NewMockImageGenerator creates a new MockImageGenerator.
func NewMockImageGenerator() *MockImageGenerator {
	return &MockImageGenerator{}
}

// GenerateCatArt returns a mock cat art image. The selection is deterministic
// based on the current date, so the same image is returned for the entire day.
func (m *MockImageGenerator) GenerateCatArt(ctx context.Context) (ImageResult, error) {
	// Use the day of year as seed for consistent daily selection
	dayOfYear := time.Now().YearDay()
	index := dayOfYear % len(mockCatImages)

	selected := mockCatImages[index]

	// Add a small variation to the caption
	suffix := []string{"✨", "🎨", "🐱", "💫", "🌟"}
	r := rand.New(rand.NewSource(int64(dayOfYear)))
	emoji := suffix[r.Intn(len(suffix))]

	return ImageResult{
		URL:     selected.URL,
		Caption: fmt.Sprintf("%s %s", selected.Caption, emoji),
	}, nil
}
