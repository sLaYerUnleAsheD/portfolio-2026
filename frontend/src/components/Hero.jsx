/**
 * Hero — Full-viewport landing section with name, tagline,
 * and floating decorative blobs for a warm, cozy first impression.
 */
export default function Hero() {
  return (
    <section
      id="hero"
      className="relative min-h-screen flex items-center justify-center overflow-hidden"
    >
      {/* ─── Decorative floating blobs ─── */}
      <div className="absolute inset-0 pointer-events-none overflow-hidden">
        <div className="absolute -top-20 -left-20 w-72 h-72 bg-blush/30 rounded-full blur-3xl animate-float" />
        <div
          className="absolute top-1/3 -right-16 w-96 h-96 bg-latte/25 rounded-full blur-3xl animate-float"
          style={{ animationDelay: '2s' }}
        />
        <div
          className="absolute -bottom-24 left-1/4 w-80 h-80 bg-moss/15 rounded-full blur-3xl animate-float"
          style={{ animationDelay: '4s' }}
        />
      </div>

      {/* ─── Content ─── */}
      <div className="relative z-10 text-center px-4 max-w-3xl mx-auto">
        {/* Greeting pill */}
        <div className="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-parchment border border-sand/50 mb-8 animate-fade-in">
          <span className="text-sm">👋</span>
          <span className="text-sm font-medium text-cocoa">Hey there, welcome!</span>
        </div>

        {/* Name */}
        <h1
          className="font-display text-5xl sm:text-6xl lg:text-7xl font-extrabold text-charcoal leading-tight mb-6 animate-fade-in-up"
          style={{ animationDelay: '0.2s' }}
        >
          Mihir{' '}
          <span className="text-transparent bg-clip-text bg-gradient-to-r from-terracotta to-cocoa">
            Panchal
          </span>
        </h1>

        {/* Tagline */}
        <p
          className="text-lg sm:text-xl text-cocoa/70 max-w-xl mx-auto mb-10 animate-fade-in-up"
          style={{ animationDelay: '0.4s' }}
        >
          Full-stack developer crafting warm, thoughtful digital experiences
          with code, creativity, and a little bit of magic ✨
        </p>

        {/* CTA buttons */}
        <div
          className="flex flex-col sm:flex-row items-center justify-center gap-4 animate-fade-in-up"
          style={{ animationDelay: '0.6s' }}
        >
          <a href="#skills" className="btn-primary">
            <span>Explore my work</span>
            <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth="2">
              <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </a>
          <a
            href="#music"
            className="inline-flex items-center gap-2 px-6 py-3 rounded-xl font-medium text-sm text-cocoa border border-sand hover:border-terracotta hover:text-terracotta transition-all duration-300"
          >
            <span>🎵</span>
            <span>Try the Music Booth</span>
          </a>
        </div>
      </div>

      {/* ─── Scroll indicator ─── */}
      <div className="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-2 animate-bounce-soft">
        <span className="text-xs font-medium text-cocoa/40 uppercase tracking-widest">scroll</span>
        <div className="w-5 h-8 rounded-full border-2 border-cocoa/20 flex justify-center pt-1.5">
          <div className="w-1 h-2 rounded-full bg-terracotta/60 animate-bounce" />
        </div>
      </div>
    </section>
  );
}
