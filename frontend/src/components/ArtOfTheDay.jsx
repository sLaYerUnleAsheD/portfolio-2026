import { useApi } from '../hooks/useApi';

/**
 * ArtOfTheDay — Displays a daily AI-generated 2D anime-style cat image.
 * Fetches from GET /api/art-of-the-day on mount. Shows a skeleton
 * loader while the request is in flight, and a friendly error state
 * if the request fails.
 */
export default function ArtOfTheDay() {
  const { data, loading, error } = useApi('/api/art-of-the-day');

  return (
    <section id="art" className="relative bg-parchment/40">
      <div className="section-container">
        <div className="text-center mb-14">
          <h2 className="section-title">Art of the Day</h2>
          <p className="section-subtitle">
            A daily AI-generated anime cat — refreshed every 24 hours 🎨
          </p>
        </div>

        <div className="max-w-lg mx-auto">
          {/* ─── Loading skeleton ─── */}
          {loading && (
            <div id="art-skeleton" className="card flex flex-col items-center gap-4">
              <div className="skeleton w-full aspect-square" />
              <div className="skeleton h-4 w-3/4 rounded-lg" />
            </div>
          )}

          {/* ─── Error state ─── */}
          {error && !loading && (
            <div id="art-error" className="card text-center py-12">
              <div className="text-4xl mb-4">😿</div>
              <p className="text-cocoa font-medium mb-1">Couldn't fetch today's art</p>
              <p className="text-sm text-cocoa/60">{error}</p>
            </div>
          )}

          {/* ─── Art display ─── */}
          {data && !loading && (
            <div id="art-display" className="card overflow-hidden group animate-fade-in">
              {/* Gallery frame */}
              <div className="relative rounded-xl overflow-hidden mb-4 bg-sand/30">
                <img
                  src={data.url}
                  alt={data.caption || 'AI-generated anime cat'}
                  className="w-full aspect-square object-cover transition-transform duration-700 group-hover:scale-105"
                  loading="lazy"
                />
                {/* Hover overlay */}
                <div className="absolute inset-0 bg-gradient-to-t from-charcoal/30 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500" />
              </div>

              {/* Caption */}
              <div className="flex items-center justify-between">
                <p className="text-sm text-cocoa/80 italic">
                  {data.caption || 'A cozy anime cat, just for today'}
                </p>
                <span className="text-xs text-cocoa/40 font-medium uppercase tracking-wider">
                  AI Art
                </span>
              </div>
            </div>
          )}
        </div>
      </div>
    </section>
  );
}
