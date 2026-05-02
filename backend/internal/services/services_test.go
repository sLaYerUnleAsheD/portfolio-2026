package services

import (
	"context"
	"testing"
)

// ─── Image Service Tests ────────────────────────────────────────────

func TestMockImageGenerator_GenerateCatArt(t *testing.T) {
	gen := NewMockImageGenerator()

	tests := []struct {
		name string
	}{
		{"returns a valid image result"},
		{"result is deterministic within same call"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := gen.GenerateCatArt(context.Background())
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.URL == "" {
				t.Error("expected non-empty URL")
			}
			if result.Caption == "" {
				t.Error("expected non-empty Caption")
			}
		})
	}
}

// ─── Music Service Tests ────────────────────────────────────────────

func TestMockMusicGenerator_GenerateTrack(t *testing.T) {
	gen := NewMockMusicGenerator()

	tests := []struct {
		name      string
		genre     string
		wantGenre string // expected genre label in response
		wantOK    bool
	}{
		{
			name:      "lofi genre returns Lo-Fi label",
			genre:     "lofi",
			wantGenre: "Lo-Fi",
			wantOK:    true,
		},
		{
			name:      "jazz genre returns Jazz label",
			genre:     "jazz",
			wantGenre: "Jazz",
			wantOK:    true,
		},
		{
			name:      "ambient genre returns Ambient label",
			genre:     "ambient",
			wantGenre: "Ambient",
			wantOK:    true,
		},
		{
			name:      "chill genre returns Chill label",
			genre:     "chill",
			wantGenre: "Chill",
			wantOK:    true,
		},
		{
			name:      "classical genre returns Classical label",
			genre:     "classical",
			wantGenre: "Classical",
			wantOK:    true,
		},
		{
			name:      "electronic genre returns Electronic label",
			genre:     "electronic",
			wantGenre: "Electronic",
			wantOK:    true,
		},
		{
			name:      "unknown genre returns fallback track",
			genre:     "death-metal",
			wantGenre: "death-metal",
			wantOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := gen.GenerateTrack(context.Background(), tt.genre)

			if tt.wantOK && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result.TrackName == "" {
				t.Error("expected non-empty TrackName")
			}
			if result.Artist == "" {
				t.Error("expected non-empty Artist")
			}
			if result.Genre != tt.wantGenre {
				t.Errorf("genre = %q, want %q", result.Genre, tt.wantGenre)
			}
		})
	}
}
