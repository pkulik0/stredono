import { Page } from 'puppeteer';
import { saveKey } from './db';
import { getMessage, getMessages } from './email';
import { getRandomIndex, randomDelay, typeWithDelay } from './util';

export const completeElevenlabsRegistration = async (i: number, page: Page, email: string, password: string) => {
	let logIndex = 0;
	const print = (msg: string) => {
		console.log(i, `[${logIndex++}]`, msg)
	}

	await page.setViewport({ width: 1920, height: 1200 });

	print("Navigating to Elevenlabs")
	await page.goto('https://elevenlabs.io/sign-up', { waitUntil: 'load' });

	print('Looking for captcha iframe');
	const iframe = await page.waitForSelector('iframe[tabindex="0"]')
	if(!iframe) {
		throw new Error('Failed to find iframe')
	}

	print('Waiting for captcha to be solved');
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
	await randomDelay(500, 1500)

	print('Looking for registration email input');
	const emailInput = await page.waitForSelector('input[name="email"]');
	if(!emailInput) {
		throw new Error('Failed to find email input')
	}
	print('Typing in registration email');
	await typeWithDelay(emailInput, email)
	await randomDelay(500, 2000)

	print('Looking for registration password input');
	const passwordInput = await page.waitForSelector('input[name="password"]');
	if(!passwordInput) {
		throw new Error('Failed to find password input')
	}
	print('Typing in registration password');
	await typeWithDelay(passwordInput, password)
	await randomDelay(500, 2000)

	print('Accepting TOS');
	// TOS Checkbox, selecting it and clicking doesn't work, so it's a workaround
	await page.mouse.click(1406, 426)
	await randomDelay(200, 1000)

	print('Submitting registration form')
	const submitBtn = await page.waitForSelector('button[type="submit"]')
	if(!submitBtn) {
		throw new Error('Failed to find submit button')
	}
	await submitBtn.click()

	print('Waiting for registration to complete and email confirmation');
	await randomDelay(10000, 15000)

	print('Getting email messages');
	const msgs = await getMessages(email)
	if(msgs.length === 0) {
		throw new Error('Registration failed or email not received')
	}

	print('Getting message content');
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

	print('Navigating to email confirmation link')
	await page.goto(matches[0], { waitUntil: 'load', timeout: 15000 })

	print('Waiting for close button');
	const closeBtn = await page.waitForSelector('button[tabindex="0"]')
	if(!closeBtn) {
		throw new Error('Failed to find close button')
	}
	await closeBtn.click()
	await randomDelay(500, 2000)

	print('Looking for email input on sign in page');
	const emailInput2 = await page.waitForSelector('input[type="email"]')
	if(!emailInput2) {
		throw new Error('Failed to find email input')
	}
	print('Typing email on sign in page');
	await typeWithDelay(emailInput2, email)
	await randomDelay(500, 2000)

	print('Looking for password input on sign in page');
	const passwordInput2 = await page.waitForSelector('input[type="password"]')
	if(!passwordInput2) {
		throw new Error('Failed to find password input')
	}
	print('Typing password on sign in page');
	await typeWithDelay(passwordInput2, password)
	await randomDelay(500, 2000)

	print('Submitting sign in form');
	const submitBtn2 = await page.waitForSelector('button[type="submit"]')
	if(!submitBtn2) {
		throw new Error('Failed to find submit button')
	}
	await submitBtn2.click()
	await page.waitForNavigation({ waitUntil: 'load', timeout: 15000 })
	await randomDelay(3000, 5000)

	// Login successful

	const firstNames = ["Alice", "Bob", "Charlie", "David", "Emily", "Adam", "Eva", "Ron", "Seth", "Ivan", "John", "Jane", "Michael", "Michelle", "Olivia", "Oscar", "Peter", "Paul", "Quentin", "Rachel", "Ralph", "Sam", "Samantha", "Tom", "Tina", "Ursula", "Ulysses", "Victor", "Victoria", "Walter", "Wendy", "Xander", "Xena", "Yvonne", "Yannick", "Zach", "Zoe"]

	print('Finding name input (onboarding)');
	const nameInput = await page.waitForSelector('input[name="name"]')
	if(!nameInput) {
		throw new Error('Failed to find name input')
	}
	await typeWithDelay(nameInput, firstNames[getRandomIndex(firstNames.length)])
	await randomDelay(500, 2000)

	print('Finding listbox');
	const selectButton = await page.waitForSelector('button[aria-haspopup="listbox"]')
	if(!selectButton) {
		throw new Error('Failed to find select button')
	}
	await selectButton.click()
	await randomDelay(500, 2000)

	print('Selecting random option');
	const listElements = await page.$$('li[tabindex="-1"]')
	if(!listElements) {
		throw new Error('Failed to find list elements')
	}
	const option = listElements[getRandomIndex(listElements.length)]
	await option.scrollIntoView()
	await option.click()
	await randomDelay(500, 2000)

	print('Clicking next button');
	await page.mouse.click(720, 423)
	await randomDelay(500, 2000)

	const pos = {
		x: 300,
		y: 420
	}
	print('Clicking on use case option');
	await page.mouse.click(pos.x, pos.y)
	await randomDelay(500, 2000)

	print('Answering second use case question');
	await page.mouse.click(pos.x, pos.y)
	await randomDelay(1000, 3000)

	// Main page

	await page.goto('https://elevenlabs.io/api', { waitUntil: 'load', timeout: 15000 })

	// Sometimes after the first reload the user is logged out, so we reload again
	await page.reload({ waitUntil: 'load', timeout: 15000 })

	print('Clicking on user menu');
	const menuBtn = await page.waitForSelector('button[data-testid="user-menu-button"]')
	if(!menuBtn) {
		throw new Error('Failed to find menu button')
	}
	await menuBtn.click()
	await randomDelay(500, 2000)

	print('Clicking on profile and api keys');
	const menuItems = await page.$$('div[role="menuitem"]')
	if(!menuItems) {
		throw new Error('Failed to find profile button')
	}
	await menuItems[0].click()
	await randomDelay(500, 2000)

	print('Finding the api key input');
	const apiKeyInput = await page.waitForSelector('input[type="password"]', { timeout: 150000 })
	if(!apiKeyInput) {
		throw new Error('Failed to find api key input')
	}
	await apiKeyInput.click()
	await randomDelay(500, 2000)

	print('Getting the api key value');
	const value = await page.$eval('input[type="password"]', (el) => el.getAttribute('value'))
	if(!value) {
		throw new Error('Failed to get api key value')
	}

	print('Saving the api key');
	await saveKey(i, value)
}