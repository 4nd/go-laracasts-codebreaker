/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./assets/**/*.{css,js,html}",
    "./views/templates/**/*.html",
  ],    theme: {
    extend: {
      colors: {
        'primary': '#0a2444'
      }
    },
  },
  plugins: [],
}

