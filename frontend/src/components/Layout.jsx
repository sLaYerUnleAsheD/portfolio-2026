import { useState, useEffect } from 'react';

/**
 * Layout — Page shell providing the sticky navigation bar, smooth-scroll
 * section links, and a subtle grain texture overlay.
 *
 * This component wraps all page content and provides the global page chrome.
 */

const NAV_LINKS = [
  { label: 'Home',   href: '#hero' },
  { label: 'Skills', href: '#skills' },
  { label: 'Art',    href: '#art' },
  { label: 'Music',  href: '#music' },
];

export default function Layout({ children }) {
  const [scrolled, setScrolled]     = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);

  /* Track scroll position for navbar backdrop effect */
  useEffect(() => {
    const onScroll = () => setScrolled(window.scrollY > 40);
    window.addEventListener('scroll', onScroll, { passive: true });
    return () => window.removeEventListener('scroll', onScroll);
  }, []);

  return (
    <div className="grain-overlay min-h-screen">
      {/* ─── Sticky Nav ─── */}
      <nav
        id="main-nav"
        className={`fixed top-0 inset-x-0 z-50 transition-all duration-500 ${
          scrolled
            ? 'bg-cream/80 backdrop-blur-lg shadow-soft py-3'
            : 'bg-transparent py-5'
        }`}
      >
        <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center justify-between">
          {/* Logo / Name */}
          <a
            href="#hero"
            className="font-display text-xl font-bold text-charcoal hover:text-terracotta transition-colors duration-300"
          >
            mihir<span className="text-terracotta">.</span>panchal
          </a>

          {/* Desktop nav links */}
          <ul className="hidden md:flex items-center gap-8">
            {NAV_LINKS.map(({ label, href }) => (
              <li key={href}>
                <a
                  href={href}
                  className="text-sm font-medium text-cocoa/80 hover:text-terracotta transition-colors duration-300 relative group"
                >
                  {label}
                  <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-terracotta rounded-full transition-all duration-300 group-hover:w-full" />
                </a>
              </li>
            ))}
          </ul>

          {/* Mobile hamburger */}
          <button
            id="mobile-menu-toggle"
            className="md:hidden p-2 text-charcoal hover:text-terracotta transition-colors"
            onClick={() => setMobileOpen((v) => !v)}
            aria-label="Toggle navigation menu"
          >
            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth="2">
              {mobileOpen ? (
                <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
              ) : (
                <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16M4 18h16" />
              )}
            </svg>
          </button>
        </div>

        {/* Mobile dropdown */}
        {mobileOpen && (
          <div className="md:hidden bg-cream/95 backdrop-blur-lg border-t border-sand/40 animate-slide-up">
            <ul className="flex flex-col items-center gap-4 py-6">
              {NAV_LINKS.map(({ label, href }) => (
                <li key={href}>
                  <a
                    href={href}
                    onClick={() => setMobileOpen(false)}
                    className="text-base font-medium text-cocoa hover:text-terracotta transition-colors"
                  >
                    {label}
                  </a>
                </li>
              ))}
            </ul>
          </div>
        )}
      </nav>

      {/* ─── Page Content ─── */}
      <main>{children}</main>
    </div>
  );
}
