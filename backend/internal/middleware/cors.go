// Package middleware provides HTTP middleware for the API server.
package middleware

import (
	"net/http"
)

// CORS returns middleware that sets Cross-Origin Resource Sharing headers.
// In development, this allows the Vite dev server (port 5173) to call
// the Go API (port 8080). In production, requests come from the same
// origin since the Go server serves the frontend.
func CORS(allowedOrigins ...string) func(http.Handler) http.Handler {
	// Default to allowing the Vite dev server
	if len(allowedOrigins) == 0 {
		allowedOrigins = []string{"http://localhost:5173"}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Check if the origin is in the allowed list or allow all with "*"
			allowed := false
			for _, o := range allowedOrigins {
				if o == "*" || o == origin {
					allowed = true
					break
				}
			}

			if allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")

			// Handle preflight requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
