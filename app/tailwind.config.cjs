/** @type {import('tailwindcss').Config}*/
const config = {
	content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],

	plugins: [require('flowbite/plugin')],

	darkMode: 'class',

	theme: {
		extend: {
			zIndex: {
				'100': '100',
				'150': '150',
				'200': '200',
			},
			colors: {
				// flowbite-svelte
				primary: {
					50: '#f2f2ff',
					100: '#eef3ff',
					200: '#dee8ff',
					300: '#ccdeff',
					400: '#adc7ff',
					500: '#5d90fe',
					600: '#2f72ef',
					700: '#276ceb',
					800: '#224acc',
					900: '#1b3ba5'
				}
			}
		}
	}
};

module.exports = config