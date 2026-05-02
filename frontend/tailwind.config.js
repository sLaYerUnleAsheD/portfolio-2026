/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        cream:      '#FFF8F0',
        sand:       '#E8DCC8',
        terracotta: '#C4755B',
        forest:     '#5B7B5E',
        charcoal:   '#3A3A3A',
        cocoa:      '#6B4F3F',
        blush:      '#F2CFC2',
        latte:      '#D4B89C',
        moss:       '#8FA98F',
        parchment:  '#F5EFE6',
      },
      fontFamily: {
        sans:    ['Inter', 'system-ui', 'sans-serif'],
        display: ['Outfit', 'system-ui', 'sans-serif'],
      },
      boxShadow: {
        'soft':    '0 2px 15px rgba(107, 79, 63, 0.08)',
        'warm':    '0 4px 25px rgba(196, 117, 91, 0.12)',
        'cozy':    '0 8px 40px rgba(107, 79, 63, 0.15)',
      },
      borderRadius: {
        'xl':  '1rem',
        '2xl': '1.5rem',
        '3xl': '2rem',
      },
      animation: {
        'float':       'float 6s ease-in-out infinite',
        'fade-in':     'fadeIn 0.8s ease-out forwards',
        'fade-in-up':  'fadeInUp 0.8s ease-out forwards',
        'slide-up':    'slideUp 0.5s ease-out forwards',
        'spin-slow':   'spin 8s linear infinite',
        'bounce-soft': 'bounceSoft 2s ease-in-out infinite',
        'peek':        'peek 0.4s cubic-bezier(0.34, 1.56, 0.64, 1) forwards',
        'wiggle':      'wiggle 0.5s ease-in-out',
        'vinyl-spin':  'spin 3s linear infinite',
        'pulse-soft':  'pulseSoft 2s ease-in-out infinite',
        'skill-fill':  'skillFill 1.2s ease-out forwards',
      },
      keyframes: {
        float: {
          '0%, 100%': { transform: 'translateY(0px)' },
          '50%':      { transform: 'translateY(-20px)' },
        },
        fadeIn: {
          '0%':   { opacity: '0' },
          '100%': { opacity: '1' },
        },
        fadeInUp: {
          '0%':   { opacity: '0', transform: 'translateY(30px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        slideUp: {
          '0%':   { opacity: '0', transform: 'translateY(20px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        bounceSoft: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%':      { transform: 'translateY(-8px)' },
        },
        peek: {
          '0%':   { transform: 'translateY(100%)' },
          '100%': { transform: 'translateY(0)' },
        },
        wiggle: {
          '0%, 100%': { transform: 'rotate(0deg)' },
          '25%':      { transform: 'rotate(-5deg)' },
          '75%':      { transform: 'rotate(5deg)' },
        },
        pulseSoft: {
          '0%, 100%': { opacity: '1' },
          '50%':      { opacity: '0.7' },
        },
        skillFill: {
          '0%':   { width: '0%' },
          '100%': { width: 'var(--skill-level)' },
        },
      },
    },
  },
  plugins: [],
};
