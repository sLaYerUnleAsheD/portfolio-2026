# ═══════════════════════════════════════════════════════════════════
# Multi-stage Dockerfile for vibe-portfolio
# Optimized for Google Cloud Run deployment
#
# Stage 1: Build React frontend → static assets
# Stage 2: Build Go backend (embeds frontend via go:embed)
# Stage 3: Minimal distroless runtime image
# ═══════════════════════════════════════════════════════════════════

# ─── Stage 1: Frontend Build ─────────────────────────────────────
FROM node:22-alpine AS frontend-builder

WORKDIR /app/frontend

# Install dependencies first (cached layer)
COPY frontend/package.json frontend/package-lock.json* ./
RUN npm ci --production=false

# Copy source and build
COPY frontend/ ./
RUN npm run build

# ─── Stage 2: Backend Build ──────────────────────────────────────
FROM golang:1.23-alpine AS backend-builder

WORKDIR /app/backend

# Install dependencies first (cached layer)
COPY backend/go.mod ./
RUN go mod download github.com/go-chi/chi/v5@v5.2.1

# Copy backend source
COPY backend/ ./
RUN go mod tidy

# Copy the frontend build output into the Go embed directory.
# main.go expects //go:embed dist/* so we place it at cmd/server/dist/
COPY --from=frontend-builder /app/frontend/dist ./cmd/server/dist/

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /server ./cmd/server

# ─── Stage 3: Minimal Runtime ───────────────────────────────────
FROM gcr.io/distroless/static-debian12:nonroot

# Copy the compiled binary
COPY --from=backend-builder /server /server

# Cloud Run uses PORT env var (default 8080)
ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
