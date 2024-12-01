/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{js,ts,jsx,tsx}',
    './index.html',
  ],
  theme: {
    extend: {
      // Add customizations if needed
      colors: {
        primary: {
          DEFAULT: '#1d4ed8',
          light: '#3b82f6',
          dark: '#1e3a8a',
        },
      },
    },
  },
  plugins: [
    require('@shadcn/ui'),
    require('@tailwindcss/forms'), // For better form styling
    require('@tailwindcss/typography'), // For rich-text formatting
  ],
};
