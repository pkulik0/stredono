import { randomUUID } from 'node:crypto';
import { ElementHandle } from 'puppeteer';

export const getRandomIndex = (length: number) => {
	return Math.floor(Math.random() * length)
}

export const generatePassword = () => {
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


export const randomDelay = async (min: number, max: number) => {
	const delay = Math.floor(Math.random() * (max - min) + min)
	await new Promise(resolve => setTimeout(resolve, delay))
}

export const typeWithDelay = async (element: ElementHandle, text: string) => {
	for(let i = 0; i < text.length; i++) {
		await element.type(text[i], { delay: Math.floor(Math.random() * (100 - 50) + 50) })
	}
}

export const getEmail = async (): Promise<string> => {
	const domains = [
		"txcct.com"
	]
	const login = randomUUID().split('-').slice(0, 3).join('')
	const loginLength = Math.random() * (login.length - 8) + 8
	return login.substring(0, loginLength-1) + '@' + domains[getRandomIndex(domains.length)]
}