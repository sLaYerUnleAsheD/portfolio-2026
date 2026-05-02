import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import HiddenCat from '../src/components/HiddenCat';

describe('HiddenCat', () => {
  it('renders the cat element', () => {
    render(<HiddenCat />);
    const cat = document.getElementById('hidden-cat');
    expect(cat).toBeInTheDocument();
  });

  it('is initially partially hidden (translated down)', () => {
    render(<HiddenCat />);
    // The cat should have translate-y-[70%] class when not peeking
    const cat = document.getElementById('hidden-cat');
    const svgContainer = cat.querySelector('div');
    expect(svgContainer.className).toContain('translate-y-[70%]');
  });

  it('shows a speech bubble with a meow message on click', () => {
    render(<HiddenCat />);
    const cat = document.getElementById('hidden-cat');

    fireEvent.click(cat);

    const bubble = document.getElementById('cat-speech-bubble');
    expect(bubble).toBeInTheDocument();
    expect(bubble.textContent).toBeTruthy();
  });

  it('has proper accessibility attributes', () => {
    render(<HiddenCat />);
    const cat = document.getElementById('hidden-cat');
    expect(cat).toHaveAttribute('role', 'button');
    expect(cat).toHaveAttribute('tabindex', '0');
    expect(cat).toHaveAttribute('aria-label', 'Hidden interactive cat');
  });
});
