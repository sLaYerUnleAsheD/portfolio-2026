// Package handlers — music.go implements the /api/generate-music endpoint.
//
// This handler accepts a genre parameter via POST and delegates to a
// MusicGenerator service to produce a track result.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mihir-panchal/vibe-portfolio/internal/services"
)

// musicRequest is the expected JSON body for POST /api/generate-music.
type musicRequest struct {
	Genre string `json:"genre"`
}

// validGenres defines the accepted genre identifiers.
var validGenres = map[string]bool{
	"lofi":       true,
	"ambient":    true,
	"jazz":       true,
	"chill":      true,
	"classical":  true,
	"electronic": true,
}

// MusicHandler handles requests for AI music generation.
type MusicHandler struct {
	generator services.MusicGenerator
}

// NewMusicHandler creates a new MusicHandler with the given music generator.
func NewMusicHandler(generator services.MusicGenerator) *MusicHandler {
	return &MusicHandler{
		generator: generator,
	}
}

// ServeHTTP handles POST /api/generate-music.
// It validates the genre, delegates to the MusicGenerator, and returns
// the generated track as JSON.
func (h *MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req musicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[music] error decoding request: %v", err)
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Normalize and validate genre
	genre := strings.TrimSpace(strings.ToLower(req.Genre))
	if genre == "" {
		http.Error(w, `{"error":"genre is required"}`, http.StatusBadRequest)
		return
	}

	if !validGenres[genre] {
		log.Printf("[music] unknown genre requested: %q (proceeding with fallback)", genre)
	}

	// Generate track
	log.Printf("[music] generating track for genre: %s", genre)
	result, err := h.generator.GenerateTrack(r.Context(), genre)
	if err != nil {
		log.Printf("[music] error generating track: %v", err)
		http.Error(w, `{"error":"failed to generate track"}`, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, result)
}
