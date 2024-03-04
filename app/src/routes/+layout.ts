import { init, register } from 'svelte-i18n';

export const ssr = false;
export const prerender = "auto";

const setupLocale = () => {
	register('en', () => import('$lib/i18n/en.json'));

	init({
		fallbackLocale: 'en',
		initialLocale: 'en'
	});
}

setupLocale()