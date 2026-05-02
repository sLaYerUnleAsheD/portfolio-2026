import { useEffect, useRef } from 'react';
import skillsData from '../data/skills.json';

/**
 * Skills — A data-driven skills matrix that reads from skills.json.
 * Each category renders as a card with animated proficiency bars.
 * The bars animate on scroll into view using IntersectionObserver.
 *
 * To update skills, edit `src/data/skills.json` — no component changes needed.
 */
export default function Skills() {
  const sectionRef = useRef(null);

  /* Animate skill bars when section scrolls into view */
  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            entry.target.classList.add('is-visible');
          }
        });
      },
      { threshold: 0.15 }
    );

    const cards = sectionRef.current?.querySelectorAll('.observe-fade');
    cards?.forEach((card) => observer.observe(card));

    return () => observer.disconnect();
  }, []);

  return (
    <section id="skills" className="relative" ref={sectionRef}>
      {/* Soft gradient divider */}
      <div className="absolute top-0 inset-x-0 h-24 bg-gradient-to-b from-cream to-transparent pointer-events-none" />

      <div className="section-container">
        <div className="text-center mb-14">
          <h2 className="section-title">Skills &amp; Expertise</h2>
          <p className="section-subtitle">Technologies I work with and love</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {skillsData.categories.map((category, catIdx) => (
            <div
              key={category.name}
              className="card observe-fade"
              style={{ transitionDelay: `${catIdx * 0.15}s` }}
            >
              {/* Category header */}
              <div className="flex items-center gap-3 mb-6">
                <span className="text-2xl">{category.icon}</span>
                <h3 className="font-display text-lg font-semibold text-charcoal">
                  {category.name}
                </h3>
              </div>

              {/* Skill bars */}
              <ul className="space-y-4">
                {category.skills.map((skill, skillIdx) => (
                  <li key={skill.name}>
                    <div className="flex justify-between items-center mb-1.5">
                      <span className="text-sm font-medium text-charcoal">{skill.name}</span>
                      <span className="text-xs font-semibold text-cocoa/60">{skill.level}%</span>
                    </div>
                    <div className="w-full h-2 bg-sand/50 rounded-full overflow-hidden">
                      <div
                        className="skill-bar-fill h-full rounded-full bg-gradient-to-r from-terracotta to-cocoa"
                        style={{
                          '--skill-level': `${skill.level}%`,
                          '--delay': `${catIdx * 0.15 + skillIdx * 0.08}s`,
                        }}
                        role="progressbar"
                        aria-valuenow={skill.level}
                        aria-valuemin={0}
                        aria-valuemax={100}
                        aria-label={`${skill.name} proficiency: ${skill.level}%`}
                      />
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
