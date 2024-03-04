import { Browser } from 'puppeteer';
import { createBrowser } from './browser';
import { handleExtension, PlanType } from './ext';
import { completeElevenlabsRegistration } from './flow';
import { generatePassword, getEmail } from './util';

export const numInstances = 5;
export const headless = true;
let success = 0;
let fails = 0;

export enum InvocationType {
	NONE, SUCCESS, FAIL
}
export const printStats = (type: InvocationType) => {
  switch(type) {
		case InvocationType.SUCCESS:
			success++
			break;
		case InvocationType.FAIL:
			fails++
			break;
		case InvocationType.NONE:
			break;
	}
	console.log(`Completed: ${success}, Fails: ${fails} NumInstances: ${numInstances}`)
}

let continueRunning = true;

const doWork = async (i: number, extension: string) => {
	console.log(i, 'Starting work');
	while(continueRunning) {
		console.log(i, 'Starting browser');
		let browser: Browser|undefined;

		try {
			browser = await createBrowser(extension)
			await completeElevenlabsRegistration(i, await browser.newPage(), getEmail(), generatePassword())
			console.log(i, 'Completed flow successfully')
		} catch(e) {
			printStats(InvocationType.FAIL)
			if(e instanceof Error) {
				console.error(i, "Retrying, error message:", e.message)
			} else {
				console.error(i, "Retrying, error:", e)
			}
		} finally {
			console.log(i, 'Closing browser')
			if(browser) {
				await browser.close()
			}
		}
	}
	console.log(i, 'Finished work');
}

const main = async () => {
	const ext = await handleExtension("pkulik-fd0871c0-39a5-46a8-9a06-8cd28d90e5b7", PlanType.PRO)
	console.log("Starting workers")
	await Promise.all(Array.from({ length: numInstances }).map((_, i) => doWork(i, ext)))
	console.log("All workers finished, exiting...")
}

main()