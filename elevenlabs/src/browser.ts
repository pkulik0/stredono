import puppeteer from 'puppeteer-extra';
import StealthPlugin from 'puppeteer-extra-plugin-stealth';
import { headless } from './main';

puppeteer.use(StealthPlugin());

export const createBrowser = async (extension: string) => {
	return await puppeteer.launch({
		headless: headless,
		ignoreDefaultArgs: [
			"--disable-extensions",
			"--enable-automation"
		],
		args: [
			'--disable-blink-features=AutomationControlled',
			'--mute-audio',
			'--no-zygote',
			'--no-xshm',
			'--no-first-run',
			'--no-default-browser-check',
			'--disable-dev-shm-usage',
			'--disable-gpu',
			'--enable-webgl',

			'--window-size=1920,1500',
			'--lang=en-US,en;q=0.9',

			`--disable-extensions-except=./${extension}`,
			`--load-extension=./${extension}`,

			'--proxy-server=socks5://p.webshare.io:9999'
		]
	});

}