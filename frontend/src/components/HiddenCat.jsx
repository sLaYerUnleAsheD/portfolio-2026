import { useState } from 'react';

/**
 * HiddenCat — A playful Easter-egg cat component hidden at the
 * bottom-right corner of the page. On hover it peeks up; on click
 * it shows a speech bubble with a random meow message.
 *
 * Renders as absolute-positioned so it doesn't affect page flow.
 */

const MEOW_MESSAGES = [
  'Meow! 🐱',
  'Purrr~ 😻',
  'Nyaa~! ✨',
  '*head bonk* 💕',
  'Feed me code! 🐾',
  '*stretches* 😸',
  'Did you find me? 🎉',
  '*purring intensifies*',
];

export default function HiddenCat() {
  const [isPeeking, setIsPeeking] = useState(false);
  const [message, setMessage]     = useState(null);
  const [wiggle, setWiggle]       = useState(false);

  /**
   * Pick a random meow message and display it in a speech bubble.
   * Clears itself after 2.5 seconds.
   */
  const handleClick = () => {
    const randomMsg = MEOW_MESSAGES[Math.floor(Math.random() * MEOW_MESSAGES.length)];
    setMessage(randomMsg);
    setWiggle(true);

    setTimeout(() => setWiggle(false), 500);
    setTimeout(() => setMessage(null), 2500);
  };

  return (
    <div
      id="hidden-cat"
      className="fixed bottom-0 right-6 z-40 cursor-pointer select-none"
      onMouseEnter={() => setIsPeeking(true)}
      onMouseLeave={() => {
        setIsPeeking(false);
        setMessage(null);
      }}
      onClick={handleClick}
      role="button"
      tabIndex={0}
      aria-label="Hidden interactive cat"
      onKeyDown={(e) => e.key === 'Enter' && handleClick()}
    >
      {/* Speech bubble */}
      {message && (
        <div
          id="cat-speech-bubble"
          className="absolute -top-14 right-0 bg-white px-4 py-2 rounded-2xl rounded-br-sm shadow-warm text-sm font-medium text-charcoal whitespace-nowrap animate-fade-in"
        >
          {message}
        </div>
      )}

      {/* Cat body */}
      <div
        className={`transition-transform duration-400 ${
          isPeeking ? 'translate-y-0' : 'translate-y-[70%]'
        } ${wiggle ? 'animate-wiggle' : ''}`}
      >
        {/* Cat SVG — a cute minimal cat face */}
        <svg
          width="64"
          height="72"
          viewBox="0 0 64 72"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
          className="drop-shadow-lg"
        >
          {/* Ears */}
          <path d="M8 28L16 4L24 24" fill="#C4755B" />
          <path d="M40 24L48 4L56 28" fill="#C4755B" />
          {/* Inner ears */}
          <path d="M12 24L16 10L20 22" fill="#F2CFC2" />
          <path d="M44 22L48 10L52 24" fill="#F2CFC2" />
          {/* Head */}
          <ellipse cx="32" cy="44" rx="26" ry="24" fill="#C4755B" />
          {/* Face lighter area */}
          <ellipse cx="32" cy="48" rx="18" ry="16" fill="#F2CFC2" />
          {/* Eyes */}
          <ellipse cx="24" cy="42" rx="3" ry={isPeeking ? 4 : 1} fill="#3A3A3A">
            {!isPeeking && (
              <animate attributeName="ry" values="1;4;1" dur="3s" repeatCount="indefinite" />
            )}
          </ellipse>
          <ellipse cx="40" cy="42" rx="3" ry={isPeeking ? 4 : 1} fill="#3A3A3A">
            {!isPeeking && (
              <animate attributeName="ry" values="1;4;1" dur="3s" repeatCount="indefinite" />
            )}
          </ellipse>
          {/* Eye shine */}
          {isPeeking && (
            <>
              <circle cx="25" cy="40" r="1.2" fill="white" />
              <circle cx="41" cy="40" r="1.2" fill="white" />
            </>
          )}
          {/* Nose */}
          <ellipse cx="32" cy="48" rx="2" ry="1.5" fill="#C4755B" />
          {/* Mouth */}
          <path d="M28 50 Q32 54 36 50" stroke="#C4755B" strokeWidth="1.5" fill="none" strokeLinecap="round" />
          {/* Whiskers */}
          <line x1="4" y1="44" x2="20" y2="46" stroke="#6B4F3F" strokeWidth="1" opacity="0.4" />
          <line x1="4" y1="48" x2="20" y2="48" stroke="#6B4F3F" strokeWidth="1" opacity="0.4" />
          <line x1="44" y1="46" x2="60" y2="44" stroke="#6B4F3F" strokeWidth="1" opacity="0.4" />
          <line x1="44" y1="48" x2="60" y2="48" stroke="#6B4F3F" strokeWidth="1" opacity="0.4" />
        </svg>
      </div>
    </div>
  );
}
