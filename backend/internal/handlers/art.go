// Package handlers implements HTTP handlers for the portfolio API.
//
// ArtHandler manages the /api/art-of-the-day endpoint, which returns
// a daily AI-generated cat image. Results are cached for 24 hours.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mihir-panchal/vibe-portfolio/internal/cache"
	"github.com/mihir-panchal/vibe-portfolio/internal/services"
)

const artCacheKey = "art-of-the-day"

// ArtHandler handles requests for the daily art image.
type ArtHandler struct {
	generator services.ImageGenerator
	cache     *cache.Cache
}

// NewArtHandler creates a new ArtHandler with the given image generator and cache.
func NewArtHandler(generator services.ImageGenerator, artCache *cache.Cache) *ArtHandler {
	return &ArtHandler{
		generator: generator,
		cache:     artCache,
	}
}

// ServeHTTP handles GET /api/art-of-the-day.
// It first checks the cache; if a valid entry exists, it returns the cached
// result. Otherwise, it calls the image generator, caches the result for 24h,
// and returns the fresh image.
func (h *ArtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	// Check cache first
	if cached, ok := h.cache.Get(artCacheKey); ok {
		if result, valid := cached.(services.ImageResult); valid {
			log.Println("[art] cache hit — returning cached art-of-the-day")
			writeJSON(w, http.StatusOK, result)
			return
		}
	}

	// Cache miss — generate fresh art
	log.Println("[art] cache miss — generating new art-of-the-day")
	result, err := h.generator.GenerateCatArt(r.Context())
	if err != nil {
		log.Printf("[art] error generating art: %v", err)
		http.Error(w, `{"error":"failed to generate art"}`, http.StatusInternalServerError)
		return
	}

	// Cache for 24 hours
	h.cache.SetWithTTL(artCacheKey, result, 24*time.Hour)

	writeJSON(w, http.StatusOK, result)
}

// writeJSON is a helper to send JSON responses with proper headers.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("[http] error encoding JSON response: %v", err)
	}
}
