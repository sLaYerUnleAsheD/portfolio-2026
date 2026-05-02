// Package services — gemini_image.go implements the ImageGenerator interface
// using Google's Gemini API (Imagen model) for real AI image generation.
//
// This generates a 2D anime-style cat image daily. The generated image is
// returned as a base64 data URL so no external storage is needed.
package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/genai"
)

// catPrompts provides varied prompts for daily cat art so each day feels unique.
var catPrompts = []string{
	"A cute 2D anime-style cat sitting on a windowsill watching rain, warm cozy lighting, soft pastel colors, studio ghibli inspired, detailed illustration",
	"A playful 2D anime kitten chasing a butterfly in a flower meadow, warm golden hour light, soft watercolor style, kawaii aesthetic",
	"A serene 2D anime cat curled up next to a steaming cup of coffee, cozy room with bookshelves, warm earth tones, lo-fi aesthetic illustration",
	"A curious 2D anime cat wearing a tiny scarf, sitting under cherry blossom trees, soft pink and cream palette, detailed manga style",
	"A majestic 2D anime cat sitting regally on a stack of old books, warm library setting, golden candlelight, cottagecore aesthetic",
	"A sleepy 2D anime cat napping in a sunbeam on a comfortable armchair, warm afternoon light, soft shadows, cozy illustration style",
	"A cheerful 2D anime cat peeking out of a cardboard box, playful expression, warm pastel background, kawaii style illustration",
}

// catCaptions provides matching captions for the generated art.
var catCaptions = []string{
	"A cozy cat watching the rain from a warm windowsill",
	"A playful kitten chasing butterflies in golden light",
	"A serene cat enjoying a quiet moment with coffee",
	"A curious cat exploring under the cherry blossoms",
	"A majestic cat presiding over a library of old books",
	"A sleepy cat bathing in warm afternoon sunbeams",
	"A cheerful cat discovering the joy of cardboard boxes",
}

// GeminiImageGenerator implements ImageGenerator using Google's Gemini Imagen API.
type GeminiImageGenerator struct {
	client *genai.Client
}

// NewGeminiImageGenerator creates a new GeminiImageGenerator with the provided API key.
func NewGeminiImageGenerator(apiKey string) (*GeminiImageGenerator, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &GeminiImageGenerator{client: client}, nil
}

// GenerateCatArt generates a 2D anime-style cat image using Imagen.
// The prompt is deterministic per day (based on day-of-year) so the same
// image concept is requested for the entire day. Combined with the 24h cache
// in the handler, only one API call is made per day.
func (g *GeminiImageGenerator) GenerateCatArt(ctx context.Context) (ImageResult, error) {
	// Pick a deterministic prompt based on the current day
	dayOfYear := time.Now().YearDay()
	promptIdx := dayOfYear % len(catPrompts)
	prompt := catPrompts[promptIdx]
	caption := catCaptions[promptIdx]

	log.Printf("[gemini-image] generating art with prompt index %d", promptIdx)

	// Call Imagen API
	config := &genai.GenerateImagesConfig{
		NumberOfImages: 1,
	}

	response, err := g.client.Models.GenerateImages(ctx, "imagen-3.0-generate-001", prompt, config)
	if err != nil {
		return ImageResult{}, fmt.Errorf("Imagen API call failed: %w", err)
	}

	if len(response.GeneratedImages) == 0 {
		return ImageResult{}, fmt.Errorf("no images returned from Imagen API")
	}

	// Convert image bytes to a base64 data URL
	img := response.GeneratedImages[0].Image
	mimeType := "image/png"
	if img.MIMEType != "" {
		mimeType = img.MIMEType
	}

	dataURL := fmt.Sprintf("data:%s;base64,%s",
		mimeType,
		base64.StdEncoding.EncodeToString(img.ImageBytes),
	)

	// Add a fun emoji to the caption
	emojis := []string{"✨", "🎨", "🐱", "💫", "🌟", "🐾", "💕"}
	r := rand.New(rand.NewSource(int64(dayOfYear)))
	emoji := emojis[r.Intn(len(emojis))]

	return ImageResult{
		URL:     dataURL,
		Caption: fmt.Sprintf("%s %s", caption, emoji),
	}, nil
}
