import '@testing-library/jest-dom';

// ─── jsdom polyfills ────────────────────────────────────────────
// IntersectionObserver is not available in jsdom. Provide a minimal
// mock so components that use it (e.g., Skills) can render in tests.
class IntersectionObserverMock {
  constructor(callback) {
    this.callback = callback;
  }
  observe() { return null; }
  unobserve() { return null; }
  disconnect() { return null; }
}

window.IntersectionObserver = IntersectionObserverMock;
