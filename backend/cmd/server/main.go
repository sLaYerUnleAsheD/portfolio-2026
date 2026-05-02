// Package main — entry point for the vibe-portfolio API server.
//
// This server:
//   - Mounts REST API endpoints under /api/
//   - Serves the React frontend as static files from the embedded dist/ folder
//   - Supports graceful shutdown
//   - Reads configuration from environment variables
//
// Environment variables:
//   PORT           — HTTP listen port (default: 8080, Cloud Run sets this)
//   USE_MOCKS      — "true" to use mock AI services (default: true)
//   IMAGE_API_KEY  — API key for real image generation (future use)
//   MUSIC_API_KEY  — API key for real music generation (future use)
package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/mihir-panchal/vibe-portfolio/internal/cache"
	"github.com/mihir-panchal/vibe-portfolio/internal/handlers"
	"github.com/mihir-panchal/vibe-portfolio/internal/middleware"
	"github.com/mihir-panchal/vibe-portfolio/internal/services"
)

// distFS embeds the frontend build output. In production, the Dockerfile
// copies the React dist/ folder here before compiling the Go binary.
// During local development, this will be empty and the Vite dev server
// is used instead (via the proxy in vite.config.js).
//
//go:embed dist/*
var distFS embed.FS

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("🚀 Starting vibe-portfolio server…")

	// ─── Configuration ──────────────────────────────────────────
	port := getEnv("PORT", "8080")

	// ─── Services ───────────────────────────────────────────────
	// Switch between mock and real Gemini-powered AI services based
	// on the USE_MOCKS environment variable.
	useMocks := getEnv("USE_MOCKS", "true")
	apiKey := getEnv("GEMINI_API_KEY", "")

	// Also check legacy env var names for backward compatibility
	if apiKey == "" {
		apiKey = getEnv("IMAGE_API_KEY", "")
	}

	var imageGen services.ImageGenerator
	var musicGen services.MusicGenerator

	if useMocks != "true" && apiKey != "" {
		log.Println("🤖 Using real Gemini AI services")

		var err error
		imageGen, err = services.NewGeminiImageGenerator(apiKey)
		if err != nil {
			log.Printf("⚠️  Failed to init Gemini image generator, falling back to mock: %v", err)
			imageGen = services.NewMockImageGenerator()
		}

		musicGen, err = services.NewGeminiMusicGenerator(apiKey)
		if err != nil {
			log.Printf("⚠️  Failed to init Gemini music generator, falling back to mock: %v", err)
			musicGen = services.NewMockMusicGenerator()
		}
	} else {
		log.Println("🎭 Using mock AI services (set USE_MOCKS=false and GEMINI_API_KEY to enable real AI)")
		imageGen = services.NewMockImageGenerator()
		musicGen = services.NewMockMusicGenerator()
	}

	// ─── Cache ──────────────────────────────────────────────────
	artCache := cache.New(24 * time.Hour)

	// ─── Router ─────────────────────────────────────────────────
	r := chi.NewRouter()

	// Middleware stack
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.Compress(5))
	r.Use(middleware.CORS("*")) // Allow all origins; tighten for production

	// API routes
	artHandler := handlers.NewArtHandler(imageGen, artCache)
	musicHandler := handlers.NewMusicHandler(musicGen)

	r.Route("/api", func(api chi.Router) {
		api.Handle("/art-of-the-day", artHandler)
		api.Handle("/generate-music", musicHandler)
	})

	// ─── Serve Frontend (Production) ────────────────────────────
	// Try to serve the embedded frontend dist. If the embed is empty
	// (local dev), this gracefully falls back to a 404, and the Vite
	// dev server handles frontend serving via proxy.
	distSubFS, err := fs.Sub(distFS, "dist")
	if err != nil {
		log.Printf("⚠️  No embedded frontend found (expected in dev mode): %v", err)
	} else {
		fileServer := http.FileServer(http.FS(distSubFS))
		// Serve SPA — for any non-API, non-file route, serve index.html
		r.Get("/*", func(w http.ResponseWriter, req *http.Request) {
			// Try to serve the file directly
			path := req.URL.Path
			if path == "/" {
				path = "/index.html"
			}
			// Check if file exists in the embedded FS
			if f, err := distSubFS.Open(path[1:]); err == nil {
				f.Close()
				fileServer.ServeHTTP(w, req)
				return
			}
			// Fallback to index.html for SPA client-side routing
			req.URL.Path = "/"
			fileServer.ServeHTTP(w, req)
		})
	}

	// ─── HTTP Server ────────────────────────────────────────────
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// ─── Graceful Shutdown ──────────────────────────────────────
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("✅ Server listening on http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	<-stop
	log.Println("🛑 Shutting down gracefully…")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Forced shutdown: %v", err)
	}

	log.Println("👋 Server stopped cleanly")
}

// getEnv reads an environment variable with a fallback default.
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
