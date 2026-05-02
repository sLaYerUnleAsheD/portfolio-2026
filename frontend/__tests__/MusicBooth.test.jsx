import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import MusicBooth from '../src/components/MusicBooth';

// Mock the useApi hook
vi.mock('../src/hooks/useApi', () => ({
  useApi: vi.fn(),
}));

import { useApi } from '../src/hooks/useApi';

describe('MusicBooth', () => {
  let mockExecute;

  beforeEach(() => {
    vi.clearAllMocks();
    mockExecute = vi.fn().mockResolvedValue({
      trackName: 'Velvet Espresso',
      artist: 'The Blue Note Cats',
      genre: 'Jazz',
    });

    useApi.mockReturnValue({
      data: null,
      loading: false,
      error: null,
      execute: mockExecute,
    });
  });

  it('renders the section title', () => {
    render(<MusicBooth />);
    expect(screen.getByText('AI Music Booth')).toBeInTheDocument();
  });

  it('renders all genre tags', () => {
    render(<MusicBooth />);

    expect(screen.getByText('Lo-Fi')).toBeInTheDocument();
    expect(screen.getByText('Ambient')).toBeInTheDocument();
    expect(screen.getByText('Jazz')).toBeInTheDocument();
    expect(screen.getByText('Chill')).toBeInTheDocument();
    expect(screen.getByText('Classical')).toBeInTheDocument();
    expect(screen.getByText('Electronic')).toBeInTheDocument();
  });

  it('calls execute with genre when a tag is clicked', async () => {
    render(<MusicBooth />);

    const jazzButton = screen.getByText('Jazz').closest('button');
    fireEvent.click(jazzButton);

    expect(mockExecute).toHaveBeenCalledWith({ genre: 'jazz' });
  });

  it('shows the empty state message initially', () => {
    render(<MusicBooth />);
    expect(screen.getByText(/select a genre/i)).toBeInTheDocument();
  });

  it('each genre button has a unique id', () => {
    render(<MusicBooth />);

    const genres = ['lofi', 'ambient', 'jazz', 'chill', 'classical', 'electronic'];
    genres.forEach((genre) => {
      expect(document.getElementById(`genre-${genre}`)).toBeInTheDocument();
    });
  });
});
