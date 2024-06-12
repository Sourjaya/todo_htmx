/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.{html,js,templ,go}",
    "./templates/common/**/*.{html,js,templ,go}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'culpa': ['"Mea Culpa"', 'cursive'],
        'tiny': ['"Tiny5"', 'sans-seriff'],
        'exo2': ['"Exo 2"', 'sans-seriff']
      }
    },
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};