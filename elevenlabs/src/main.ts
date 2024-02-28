import { createBrowser } from './browser';
import { handleExtension, PlanType } from './ext';
import { completeElevenlabsRegistration } from './flow';
import { generatePassword, getEmail } from './util';

export const numInstances = 2;
export const headless = false;
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
		const browser = await createBrowser(extension)
		const page = await browser.newPage()

		try {
			const email = await getEmail()
			const password = generatePassword()
			await completeElevenlabsRegistration(i, page, email, password)
			console.log(i, 'Completed flow successfully')
		} catch(e) {
			printStats(InvocationType.FAIL)
		  console.error(i, "Error, retrying", e)
		} finally {
			console.log(i, 'Closing browser')
			await browser.close()
		}
	}
	console.log(i, 'Finished work');
}

const main = async () => {
	const ext = await handleExtension("pkulik0-d5225369-c4cc-45b4-000d-4fda0b44a3c7", PlanType.PRO)
	console.log("Starting workers")
	await Promise.all(Array.from({ length: numInstances }).map((_, i) => doWork(i, ext)))
	console.log("All workers finished, exiting...")
}

main()