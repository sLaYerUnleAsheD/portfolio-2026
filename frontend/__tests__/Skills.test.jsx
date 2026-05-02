import { render, screen } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import Skills from '../src/components/Skills';

describe('Skills', () => {
  it('renders the skills section heading', () => {
    render(<Skills />);
    expect(screen.getByText('Skills & Expertise')).toBeInTheDocument();
  });

  it('renders all three skill categories', () => {
    render(<Skills />);
    expect(screen.getByText('Languages')).toBeInTheDocument();
    expect(screen.getByText('Frameworks & Libraries')).toBeInTheDocument();
    expect(screen.getByText('Tools & Platforms')).toBeInTheDocument();
  });

  it('renders individual skills from the data file', () => {
    render(<Skills />);
    // Check a few skills from each category
    expect(screen.getByText('JavaScript')).toBeInTheDocument();
    expect(screen.getByText('React')).toBeInTheDocument();
    expect(screen.getByText('Docker')).toBeInTheDocument();
  });

  it('renders skill proficiency percentages', () => {
    render(<Skills />);
    // JavaScript is 90% in the data file — multiple skills may share this value
    const matches = screen.getAllByText('90%');
    expect(matches.length).toBeGreaterThan(0);
  });

  it('renders progress bars with correct aria attributes', () => {
    render(<Skills />);
    const progressBars = screen.getAllByRole('progressbar');
    expect(progressBars.length).toBeGreaterThan(0);

    // Check that each progress bar has the required aria attributes
    progressBars.forEach((bar) => {
      expect(bar).toHaveAttribute('aria-valuemin', '0');
      expect(bar).toHaveAttribute('aria-valuemax', '100');
      expect(bar).toHaveAttribute('aria-valuenow');
    });
  });
});
