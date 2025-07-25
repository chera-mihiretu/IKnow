import React from 'react';
import Image from 'next/image';
import Link from 'next/link';
import { COLORS, FONT_FAMILY } from '../../utils/color';

const AboutSection: React.FC = () => (
  <section
    className="w-full max-w-full bg-white flex flex-col items-center justify-center py-16 pb-10 relative box-border overflow-hidden m-0 left-auto right-auto ml-0 mr-0"
  >
    {/* Subtle abstract background element (optional) */}
    <svg
      width="120"
      height="120"
      viewBox="0 0 120 120"
      fill="none"
      className="absolute top-6 left-8 opacity-5 z-0 pointer-events-none"
      aria-hidden="true"
    >
      {/* Graduation cap outline */}
      <path d="M20 60L60 40L100 60L60 80L20 60Z" stroke={COLORS.primary} strokeWidth="2.5" fill="none" />
      <circle cx="60" cy="60" r="56" stroke={COLORS.primary} strokeWidth="1.5" fill="none" />
    </svg>
    {/* Brand Logo */}
    <div className="z-[1] mb-4.5 flex flex-col items-center">
      <Image
        src="/logos/logo-real.png"
        alt="IKnow Logo"
        width={64}
        height={64}
        className="mb-2"
        priority
      />
    </div>
    {/* Mission Statement */}
    <div className="z-[1] max-w-[600px] text-center mx-auto px-6">
      <p
        style={{ fontFamily: FONT_FAMILY.poppins }}
        className="font-medium text-[1.35rem] text-[var(--foreground)] leading-[1.7] tracking-[0.01em] mb-0"
      >
        Empowering students with the tools and connections they need to <span style={{ color: COLORS.primary, fontWeight: 600 }}>succeed</span> in their academic journey and beyond.
      </p>
    </div>
    {/* Get Started Button */}
    <div className="z-[1] mt-8">
      <Link
        href="/auth/login"
        style={{
          background: COLORS.primary,
          fontFamily: FONT_FAMILY.poppins,
        }}
        className="text-white font-semibold text-[1.08rem] rounded-lg px-9 py-[0.7rem] no-underline shadow-[0_2px_8px_rgba(67,24,209,0.07)] transition-colors duration-200 tracking-[0.01em] inline-block mr-2"
      >
        Get Started
      </Link>
    </div>
    {/* Social Media Links (optional) */}
    <div className="flex gap-[18px] mt-9 justify-center z-[1]">
      <a
        href="https://twitter.com/"
        target="_blank"
        rel="noopener noreferrer"
        aria-label="Twitter"
        style={{ color: COLORS.primary, opacity: 0.7 }}
      >
        <svg width="24" height="24" fill="none" viewBox="0 0 24 24"><path d="M22 5.92a8.38 8.38 0 0 1-2.36.65A4.13 4.13 0 0 0 21.4 4.1a8.27 8.27 0 0 1-2.61 1A4.13 4.13 0 0 0 12 8.13c0 .32.04.64.1.94A11.7 11.7 0 0 1 3.1 4.6a4.13 4.13 0 0 0 1.28 5.5A4.07 4.07 0 0 1 2.8 9.1v.05a4.13 4.13 0 0 0 3.3 4.05c-.2.05-.4.08-.62.08-.15 0-.3-.01-.44-.04a4.13 4.13 0 0 0 3.85 2.86A8.3 8.3 0 0 1 2 19.54a11.7 11.7 0 0 0 6.29 1.84c7.55 0 11.68-6.26 11.68-11.68 0-.18-.01-.36-.02-.54A8.18 8.18 0 0 0 22 5.92z" fill="currentColor"/></svg>
      </a>
      <a
        href="https://facebook.com/"
        target="_blank"
        rel="noopener noreferrer"
        aria-label="Facebook"
        style={{ color: COLORS.primary, opacity: 0.7 }}
      >
        <svg width="24" height="24" fill="none" viewBox="0 0 24 24"><path d="M22.675 0h-21.35C.595 0 0 .592 0 1.326v21.348C0 23.408.595 24 1.325 24h11.495v-9.294H9.692v-3.622h3.128V8.413c0-3.1 1.893-4.788 4.659-4.788 1.325 0 2.463.099 2.797.143v3.24l-1.918.001c-1.504 0-1.797.715-1.797 1.763v2.313h3.587l-.467 3.622h-3.12V24h6.116C23.406 24 24 23.408 24 22.674V1.326C24 .592 23.406 0 22.675 0" fill="currentColor"/></svg>
      </a>
      <a
        href="https://linkedin.com/"
        target="_blank"
        rel="noopener noreferrer"
        aria-label="LinkedIn"
        style={{ color: COLORS.primary, opacity: 0.7 }}
      >
        <svg width="24" height="24" fill="none" viewBox="0 0 24 24"><path d="M20.447 20.452h-3.554v-5.569c0-1.327-.025-3.037-1.849-3.037-1.851 0-2.132 1.445-2.132 2.939v5.667H9.358V9h3.414v1.561h.049c.476-.899 1.637-1.849 3.37-1.849 3.602 0 4.267 2.369 4.267 5.455v6.285zM5.337 7.433a2.062 2.062 0 1 1 0-4.124 2.062 2.062 0 0 1 0 4.124zM7.119 20.452H3.554V9h3.565v11.452zM22.225 0H1.771C.792 0 0 .771 0 1.723v20.549C0 23.229.792 24 1.771 24h20.451C23.2 24 24 23.229 24 22.271V1.723C24 .771 23.2 0 22.225 0z" fill="currentColor"/></svg>
      </a>
    </div>
    {/* Minimal Footer Navigation (optional) */}
    <footer
      style={{ fontFamily: FONT_FAMILY.poppins }}
      className="mt-8 text-[1rem] text-[#b0b0b0] flex gap-6 justify-center z-[1]"
    >
      <a href="#" className="text-[#b0b0b0] no-underline transition-colors duration-200">Home</a>
      <a href="#" className="text-[#b0b0b0] no-underline transition-colors duration-200">Features</a>
      <a href="#" className="text-[#b0b0b0] no-underline transition-colors duration-200">Contact</a>
    </footer>
  </section>
);

export default AboutSection;
