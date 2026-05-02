import { useState } from 'react';
import { useApi } from '../hooks/useApi';

/**
 * MusicBooth — Interactive AI music generation section.
 * Users select a genre tag, which triggers a POST to /api/generate-music.
 * The response (track name, artist, genre) is displayed with a spinning
 * vinyl record animation.
 */

const GENRES = [
  { id: 'lofi',      label: 'Lo-Fi',      emoji: '🎧' },
  { id: 'ambient',   label: 'Ambient',    emoji: '🌙' },
  { id: 'jazz',      label: 'Jazz',       emoji: '🎷' },
  { id: 'chill',     label: 'Chill',      emoji: '🌊' },
  { id: 'classical', label: 'Classical',  emoji: '🎻' },
  { id: 'electronic',label: 'Electronic', emoji: '⚡' },
];

export default function MusicBooth() {
  const [selectedGenre, setSelectedGenre] = useState(null);
  const [track, setTrack]                 = useState(null);
  const [spinning, setSpinning]           = useState(false);

  const { loading, error, execute } = useApi('/api/generate-music', {
    method: 'POST',
    immediate: false,
  });

  /**
   * Handle genre tag click — fires the API call and displays results.
   */
  const handleGenreClick = async (genre) => {
    setSelectedGenre(genre.id);
    setSpinning(true);
    setTrack(null);

    const result = await execute({ genre: genre.id });

    if (result) {
      setTrack(result);
    }
    /* Keep spinning while there's a track (visual effect) */
  };

  return (
    <section id="music" className="relative">
      <div className="section-container">
        <div className="text-center mb-14">
          <h2 className="section-title">AI Music Booth</h2>
          <p className="section-subtitle">
            Pick a genre and let the AI compose something special 🎶
          </p>
        </div>

        <div className="max-w-2xl mx-auto">
          {/* ─── Genre tags ─── */}
          <div
            id="genre-tags"
            className="flex flex-wrap items-center justify-center gap-3 mb-12"
          >
            {GENRES.map((genre) => (
              <button
                key={genre.id}
                id={`genre-${genre.id}`}
                onClick={() => handleGenreClick(genre)}
                disabled={loading}
                className={`btn-tag flex items-center gap-2 ${
                  selectedGenre === genre.id ? 'btn-tag--active' : ''
                } ${loading ? 'opacity-50 cursor-not-allowed' : ''}`}
              >
                <span>{genre.emoji}</span>
                <span>{genre.label}</span>
              </button>
            ))}
          </div>

          {/* ─── Vinyl + Result Display ─── */}
          <div className="flex flex-col items-center">
            {/* Vinyl record */}
            <div
              className={`relative w-48 h-48 sm:w-56 sm:h-56 rounded-full vinyl-record shadow-cozy mb-8 ${
                spinning && track ? 'animate-vinyl-spin' : ''
              }`}
            >
              {/* Center label */}
              <div className="absolute inset-0 flex items-center justify-center">
                <div className="w-16 h-16 sm:w-20 sm:h-20 rounded-full bg-cream border-4 border-sand/40 flex items-center justify-center">
                  {loading ? (
                    <div className="w-6 h-6 border-2 border-terracotta/30 border-t-terracotta rounded-full animate-spin" />
                  ) : track ? (
                    <span className="text-2xl">🎵</span>
                  ) : (
                    <span className="text-2xl">🎶</span>
                  )}
                </div>
              </div>
              {/* Groove rings */}
              <div className="absolute inset-6 rounded-full border border-white/5" />
              <div className="absolute inset-10 rounded-full border border-white/5" />
              <div className="absolute inset-14 rounded-full border border-white/5" />
            </div>

            {/* Loading state */}
            {loading && (
              <p className="text-sm text-cocoa/60 animate-pulse-soft">
                Composing your track…
              </p>
            )}

            {/* Error state */}
            {error && !loading && (
              <div id="music-error" className="text-center">
                <p className="text-cocoa font-medium">Oops, couldn't generate a track</p>
                <p className="text-sm text-cocoa/60 mt-1">{error}</p>
              </div>
            )}

            {/* Track result */}
            {track && !loading && (
              <div
                id="music-result"
                className="glass text-center px-8 py-6 animate-fade-in-up w-full max-w-md"
              >
                <p className="text-xs uppercase tracking-widest text-cocoa/50 mb-2">Now Playing</p>
                <h3 className="font-display text-xl font-bold text-charcoal mb-1">
                  {track.trackName}
                </h3>
                <p className="text-sm text-cocoa/70 mb-3">{track.artist}</p>
                <span className="inline-block px-3 py-1 rounded-full bg-terracotta/10 text-terracotta text-xs font-semibold">
                  {track.genre}
                </span>
              </div>
            )}

            {/* Initial empty state */}
            {!track && !loading && !error && (
              <p className="text-sm text-cocoa/40 italic">
                Select a genre above to generate a track ☝️
              </p>
            )}
          </div>
        </div>
      </div>
    </section>
  );
}
