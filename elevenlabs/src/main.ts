import { firestore } from 'firebase-admin';
import { randomUUID } from 'node:crypto';
import { ElementHandle, Page } from 'puppeteer';
import puppeteer from 'puppeteer-extra';
import StealthPlugin from 'puppeteer-extra-plugin-stealth';
import { getFirestore } from 'firebase-admin/firestore';
import { initializeApp, applicationDefault } from 'firebase-admin/app';
import FieldValue = firestore.FieldValue;


const db = getFirestore(initializeApp({
	credential: applicationDefault()
}));

puppeteer.use(StealthPlugin());

const createBrowser = async () => {
	return await puppeteer.launch({
		headless: true,
		ignoreDefaultArgs: [
			"--disable-extensions",
			"--enable-automation"
		],
		args: [
			'--disable-gpu-sandbox',
			'--disable-software-rasterizer',
			'--disable-background-timer-throttling',
			'--disable-backgrounding-occluded-windows',
			'--disable-renderer-backgrounding',
			'--disable-infobars',
			'--disable-breakpad',
			'--disable-canvas-aa',
			'--disable-2d-canvas-clip-aa',
			'--disable-gl-drawing-for-tests',

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

			'--disable-extensions-except=./captcha-solver',
			'--load-extension=./captcha-solver',

			'--proxy-server=socks5://p.webshare.io:9999'
		]
	});

}

let numInstances = 5;
const retries = 10;
let success = 0;
let failed = 0;

const printStats = () => {
	console.log(`Success: ${success}, Failed: ${failed}, Left: ${numInstances-success-failed}`);
}

const saveKey = async (i: number, key: string) => {
	const docRef = await db.collection('api-keys').doc('elevenlabs').set({
		keys: FieldValue.arrayUnion(key)
	}, { merge: true })
	console.log(i, 'wrote key to db @', docRef.writeTime);
	success++;
	printStats()
}

const randomDelay = async (min: number, max: number) => {
	const delay = Math.floor(Math.random() * (max - min) + min)
	await new Promise(resolve => setTimeout(resolve, delay))
}

const typeWithDelay = async (element: ElementHandle, text: string) => {
	for(let i = 0; i < text.length; i++) {
		await element.type(text[i], { delay: Math.floor(Math.random() * (100 - 50) + 50) })
	}
}

const completeElevenlabsRegistration = async (i: number, page: Page, email: string, password: string) => {
	await page.setViewport({ width: 1920, height: 1200 });

	console.log(i, 'Navigating to Elevenlabs');
	await page.goto('https://elevenlabs.io/sign-up', { waitUntil: 'networkidle2' });

	console.log(i, 'Accept cookies');
	const acceptCookies = await page.waitForSelector('a[id="CybotCookiebotDialogBodyButtonAccept"]')
	if(!acceptCookies) {
		throw new Error('Failed to find accept cookies button')
	}
	await acceptCookies.click()
	await randomDelay(500, 2000)

	console.log(i, 'Waiting for email input');
	const emailInput = await page.waitForSelector('input[name="email"]');
	if(!emailInput) {
		throw new Error('Failed to find email input')
	}
	console.log(i, 'Typing email');
	await typeWithDelay(emailInput, email)
	await randomDelay(500, 2000)

	console.log(i, 'Waiting for password input');
	const passwordInput = await page.waitForSelector('input[name="password"]');
	if(!passwordInput) {
		throw new Error('Failed to find password input')
	}
	console.log(i, 'Typing password');
	await typeWithDelay(passwordInput, password)
	await randomDelay(500, 2000)

	console.log(i, 'Looking for captcha iframe');
	const iframe = await page.waitForSelector('iframe[tabindex="0"]')
	if(!iframe) {
		throw new Error('Failed to find iframe')
	}

	console.log(i, 'Waiting for captcha to be solved');
	const frame = await iframe.contentFrame()
	const res = await frame.waitForFunction(() => {
		const checkbox = document.getElementById("checkbox")
		if(!checkbox) {
			return false
		}
		return checkbox.getAttribute('aria-checked') === 'true'
	}, { timeout: 60000 })
	if(!res) {
		throw new Error('Failed to solve captcha')
	}
	await randomDelay(200, 1000)


	console.log(i, 'Accepting TOS');
	// TOS Checkbox, selecting it and clicking doesn't work, so it's a workaround
	await page.mouse.click(1406, 426)
	await randomDelay(200, 1000)

	console.log(i, "Submitting form")
	const submitBtn = await page.waitForSelector('button[type="submit"]')
	if(!submitBtn) {
		throw new Error('Failed to find submit button')
	}
	await submitBtn.click()

	console.log(i, 'Waiting for registration to complete and email confirmation');
	await randomDelay(10000, 15000)

	console.log(i, 'Getting email messages');
	const msgs = await getMessages(email)
	if(msgs.length === 0) {
		throw new Error('Registration failed or email not received')
	}

	console.log(i, 'Getting message');
	const msg = await getMessage(email, msgs[0].id)
	if(!msg) {
		throw new Error('Failed to get message')
	}

	const extractPattern = "https?:\\/\\/beta[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b(?:[-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)"
	const matches = msg.match(extractPattern)
	if(!matches) {
		throw new Error('Failed to extract link')
	}

	// Back to the page

	console.log(i, 'Navigating to email confirmation link')
	await page.goto(matches[0], { waitUntil: 'networkidle2', timeout: 15000 })

	console.log(i, 'Waiting for close button');
	const closeBtn = await page.waitForSelector('button[tabindex="0"]')
	if(!closeBtn) {
		throw new Error('Failed to find close button')
	}
	await closeBtn.click()
	await randomDelay(500, 2000)

	console.log(i, "Entering email");
	const emailInput2 = await page.waitForSelector('input[type="email"]')
	if(!emailInput2) {
		throw new Error('Failed to find email input')
	}
	await typeWithDelay(emailInput2, email)
	await randomDelay(500, 2000)

	console.log(i, "Entering password");
	const passwordInput2 = await page.waitForSelector('input[type="password"]')
	if(!passwordInput2) {
		throw new Error('Failed to find password input')
	}
	await typeWithDelay(passwordInput2, password)
	await randomDelay(500, 2000)

	console.log(i, "Submitting form");
	const submitBtn2 = await page.waitForSelector('button[type="submit"]')
	if(!submitBtn2) {
		throw new Error('Failed to find submit button')
	}
	await submitBtn2.click()
	await page.waitForNavigation({ waitUntil: 'networkidle2', timeout: 15000 })
	await randomDelay(3000, 5000)

	// Login successful

	const firstNames = ["Alice", "Bob", "Charlie", "David", "Emily", "Adam", "Eva", "Ron", "Seth", "Ivan", "John", "Jane", "Michael", "Michelle", "Olivia", "Oscar", "Peter", "Paul", "Quentin", "Rachel", "Ralph", "Sam", "Samantha", "Tom", "Tina", "Ursula", "Ulysses", "Victor", "Victoria", "Walter", "Wendy", "Xander", "Xena", "Yvonne", "Yannick", "Zach", "Zoe"]

	console.log(i, 'Finding the name input');
	const nameInput = await page.waitForSelector('input[name="name"]')
	if(!nameInput) {
		throw new Error('Failed to find name input')
	}
	await typeWithDelay(nameInput, firstNames[getRandomIndex(firstNames.length)])
	await randomDelay(500, 2000)

	console.log(i, 'Finding listbox');
	const selectButton = await page.waitForSelector('button[aria-haspopup="listbox"]')
	if(!selectButton) {
		throw new Error('Failed to find select button')
	}
	await selectButton.click()
	await randomDelay(500, 2000)

	console.log(i, 'Selecting random option');
	const listElements = await page.$$('li[tabindex="-1"]')
	if(!listElements) {
		throw new Error('Failed to find list elements')
	}
	const option = listElements[getRandomIndex(listElements.length)]
	await option.scrollIntoView()
	await option.click()
	await randomDelay(500, 2000)

	console.log(i, 'Clicking next button');
	await page.mouse.click(720, 423)
	await randomDelay(500, 2000)

	const pos = {
		x: 300,
		y: 420
	}
	console.log(i, 'Clicking on use case option');
	await page.mouse.click(pos.x, pos.y)
	await randomDelay(500, 2000)

	console.log(i, 'Answering second use case question');
	await page.mouse.click(pos.x, pos.y)
	await randomDelay(1000, 3000)

	// Main page

	await page.goto('https://elevenlabs.io/api', { waitUntil: 'networkidle2', timeout: 15000 })

	// Sometimes after the first reload the user is logged out, so we reload again
	await page.reload({ waitUntil: 'networkidle2', timeout: 15000 })

	console.log(i, 'Clicking on menu');
	const menuBtn = await page.waitForSelector('button[data-testid="user-menu-button"]')
	if(!menuBtn) {
		throw new Error('Failed to find menu button')
	}
	await menuBtn.click()
	await randomDelay(500, 2000)

	console.log(i, 'Clicking on profile and api keys');
	const menuItems = await page.$$('div[role="menuitem"]')
	if(!menuItems) {
		throw new Error('Failed to find profile button')
	}
	await menuItems[0].click()
	await randomDelay(500, 2000)

	console.log(i, 'Finding the api key input');
	const apiKeyInput = await page.waitForSelector('input[type="password"]', { timeout: 150000 })
	if(!apiKeyInput) {
		throw new Error('Failed to find api key input')
	}
	await apiKeyInput.click()
	await randomDelay(500, 2000)

	const value = await page.$eval('input[type="password"]', (el) => el.getAttribute('value'))
	if(!value) {
		throw new Error('Failed to get api key value')
	}
	await saveKey(i, value)
}

const getEmail = async (): Promise<string> => {
	const domains = [
		"txcct.com"
	]
	const login = randomUUID().split('-').slice(0, 3).join('')
	return login + '@' + domains[getRandomIndex(domains.length)]
}

interface EmailMessage {
	id: number
	from: string
	subject: string
	date: string
}

const getRandomIndex = (length: number) => {
	return Math.floor(Math.random() * length)
}

const generatePassword = () => {
	const lower = "abcdefghijklmnopqrstuvwxyz"
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const numbers = "0123456789"
	const special = "!?@#$%^&*"
	const charset = lower + upper + numbers + special

	const length = Math.random() * (20 - 8) + 8
	let password = ""

	password += lower[getRandomIndex(lower.length)]
	password += upper[getRandomIndex(upper.length)]
	for(let i = 4; i < length; i++) {
		password += charset[getRandomIndex(charset.length)]
	}
	password += numbers[getRandomIndex(numbers.length)]
	password += special[getRandomIndex(special.length)]

	return password
}

const getMessages = async (email: string): Promise<EmailMessage[]> => {
	const login = email.split('@')[0]
	const domain = email.split('@')[1]

	const res = await fetch(`https://www.1secmail.com/api/v1/?action=getMessages&login=${login}&domain=${domain}`)
	if(!res.ok) {
		throw new Error('Failed to get messages')
	}

	return await res.json() as EmailMessage[]
}

const getMessage = async (email: string, id: number): Promise<string> => {
	const login = email.split('@')[0]
	const domain = email.split('@')[1]

	const res = await fetch(`https://www.1secmail.com/api/v1/?action=readMessage&login=${login}&domain=${domain}&id=${id}`)
	if(!res.ok) {
		throw new Error('Failed to get message')
	}

	const data = await res.json()
	return data.body
}

const doWork = async (i: number) => {
	console.log(i, 'Starting browser');
	const browser = await createBrowser()
	const page = await browser.newPage()

	try {
		const email = await getEmail()
		const password = generatePassword()
		await completeElevenlabsRegistration(i, page, email, password)
	} catch(e) {
		await page.screenshot({ path: `screenshots/error-${i}-${Date.now()}.png` })
		throw e;
	} finally {
		console.log(i, 'Closing browser')
		await browser.close()
	}
}

const run = async (i: number, retry: number = 0) => {
	await new Promise(resolve => setTimeout(resolve, 1000 * Math.random() * 3))
	try  {
		await doWork(i)
	} catch(e) {
		if(retry < retries) {
			console.error(i, `Retrying ${retry+1} of ${retries}`)
			await run(i, retry + 1)
		} else {
			console.error(i, `Failed after ${retries} retries`)
			failed++
		}
	}
	await run(i, 0)
}

for(let i = 0; i < numInstances; i++) {
	run(i)
}