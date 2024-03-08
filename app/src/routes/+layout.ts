import { init, locale, register } from 'svelte-i18n';
import Cookies from 'js-cookie';

export const ssr = false;
export const prerender = "auto";

const getUserLocale = () => {
	const localeCookie = Cookies.get("locale");
	const navigatorLocale = navigator.language.split("-")[0];
	return localeCookie || navigatorLocale;
}

const setupLocale = () => {
	const languages = ['en', 'pl', 'es', 'fr', 'de'];
	languages.forEach(lang => {
		register(lang, () => import(`$lib/i18n/${lang}.json`));
	})

	init({
		fallbackLocale: 'en',
		initialLocale: getUserLocale(),
	});

	locale.subscribe(v => {
		if(!v) {
			Cookies.remove('locale')
			return;
		}
		Cookies.set('locale', v);
	})
}

setupLocale()