# 🎨 Vibe Portfolio

A warm, minimal, and cozy interactive portfolio website built with **React + Tailwind CSS** (frontend) and **Go** (backend API gateway). Features AI-powered daily art generation and an interactive music booth.

## ✨ Features

- **Hidden Interactive Cat** — Easter egg with hover/click animations
- **Skills Matrix** — Data-driven, easily updatable via `skills.json`
- **Art of the Day** — AI-generated anime cat art (cached 24h)
- **AI Music Booth** — Genre-based music track generation
- **Responsive** — Mobile-first design with warm earth tones
- **Dockerized** — Multi-stage build optimized for Cloud Run

## 🏗️ Architecture

```
┌─────────────┐     ┌──────────────────────────────┐
│   Browser   │────▶│  Go Server (:8080)           │
│  React SPA  │     │  ├── /api/art-of-the-day     │
│             │◀────│  ├── /api/generate-music     │
│             │     │  └── /* (serves frontend)    │
└─────────────┘     └──────────────────────────────┘
                              │
                    ┌─────────┴──────────┐
                    │  Service Layer     │
                    │  (Interface-based) │
                    │  ├── ImageGenerator│
                    │  └── MusicGenerator│
                    └────────────────────┘
```

## 🚀 Quick Start

### Local Development (with Docker)

```bash
docker compose up
```
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

### Local Development (without Docker)

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```

**Backend (requires Go 1.23+):**
```bash
cd backend
go run ./cmd/server
```

### Run Tests

**Frontend:**
```bash
cd frontend
npm test
```

**Backend:**
```bash
cd backend
go test ./...
```

## 🌐 Deploy to Cloud Run

```bash
gcloud run deploy vibe-portfolio \
  --source . \
  --project vibe-code-495106 \
  --region us-central1 \
  --allow-unauthenticated
```

## 🔧 Configuration

Copy `.env.example` to `.env` and fill in your API keys:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `USE_MOCKS` | Use mock AI services | `true` |
| `IMAGE_API_KEY` | Nano Banana API key | — |
| `MUSIC_API_KEY` | Music generation API key | — |

## 📁 Updating Content

- **Skills**: Edit `frontend/src/data/skills.json`
- **Social links**: Edit the `SOCIAL_LINKS` array in `frontend/src/components/Footer.jsx`
- **Hero text**: Edit `frontend/src/components/Hero.jsx`

## 📄 License

MIT
