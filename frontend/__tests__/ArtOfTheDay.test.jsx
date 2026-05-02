import { render, screen } from '@testing-library/react';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import ArtOfTheDay from '../src/components/ArtOfTheDay';

// Mock the useApi hook
vi.mock('../src/hooks/useApi', () => ({
  useApi: vi.fn(),
}));

import { useApi } from '../src/hooks/useApi';

describe('ArtOfTheDay', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('shows skeleton loader while loading', () => {
    useApi.mockReturnValue({ data: null, loading: true, error: null });

    render(<ArtOfTheDay />);

    expect(document.getElementById('art-skeleton')).toBeInTheDocument();
  });

  it('renders the image when data is loaded', () => {
    useApi.mockReturnValue({
      data: {
        url: 'https://example.com/cat.png',
        caption: 'A cute anime cat',
      },
      loading: false,
      error: null,
    });

    render(<ArtOfTheDay />);

    const display = document.getElementById('art-display');
    expect(display).toBeInTheDocument();

    const img = screen.getByAltText('A cute anime cat');
    expect(img).toHaveAttribute('src', 'https://example.com/cat.png');
  });

  it('shows error state when fetch fails', () => {
    useApi.mockReturnValue({
      data: null,
      loading: false,
      error: 'Network error',
    });

    render(<ArtOfTheDay />);

    const errorEl = document.getElementById('art-error');
    expect(errorEl).toBeInTheDocument();
    expect(screen.getByText('Network error')).toBeInTheDocument();
  });

  it('renders section title and subtitle', () => {
    useApi.mockReturnValue({ data: null, loading: true, error: null });

    render(<ArtOfTheDay />);

    expect(screen.getByText('Art of the Day')).toBeInTheDocument();
  });
});
