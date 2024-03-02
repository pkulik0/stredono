import AdmZip from 'adm-zip';
import { writeFile, rm, readFile } from 'node:fs/promises';
import path from 'node:path';

const getExtension = async () => {
	const url = "https://api.github.com/repos/noCaptchaAi/chrome/releases/latest"
	const response = await fetch(url)
	if (!response.ok) {
		throw new Error(`HTTP error! status: ${response.status}`);
	}

	const data = await response.json()
	const assets = data.assets

	const extension = assets.find((asset: any) => asset.name.includes(".zip") && !asset.name.includes(".crx"))
	if (!extension) {
		throw new Error("No extension found")
	}

	const extensionUrl = extension.browser_download_url
	const name = extension.name.split(".zip")[0]
		console.log("Downloading extension from", extensionUrl)
	const extensionResponse = await fetch(extensionUrl, {
		headers: {
			"Accept": "application/blob"
		}
	})

	if (!extensionResponse.ok) {
		throw new Error(`HTTP error! status: ${extensionResponse.status}`);
	}

	const extensionBuffer = await extensionResponse.blob()
	const arrayBuffer = await extensionBuffer.arrayBuffer();
	await writeFile(path.join(process.cwd(), "nocaptchaai.zip"), Buffer.from(arrayBuffer))

	return name
}

const unpackExtension = async () => {
	const zip = new AdmZip(path.join(process.cwd(), "nocaptchaai.zip"))
	zip.extractAllTo(process.cwd(), true)
}

export enum PlanType {
	FREE = "free",
	PRO = "pro"
}

const editConfig = async (version: string, apiKey: string, plan: PlanType) => {
	const configPath = path.join(process.cwd(), version, "src", "background.js")
	const fileContent = await readFile(configPath, "utf-8")
	const newContent = fileContent.replace('APIKEY: ""', `APIKEY: "${apiKey}"`).replace('PLANTYPE: "pro"', `PLANTYPE: "${plan}"`)
	await writeFile(configPath, newContent)
}

export const handleExtension = async (apiKey: string, plan: PlanType): Promise<string> => {
	console.log("Removing old extension if exists")
	await rm(path.join(process.cwd(), "nocaptcha"), { recursive: true, force: true })
	console.log("Downloading noCaptchaAi extension")
	const version = await getExtension()
	console.log("Unpacking extension zip")
	await unpackExtension()
	await rm(path.join(process.cwd(), "nocaptchaai.zip"))
	console.log("Editing extension config")
	await editConfig(version, apiKey, plan)
	console.log("Extension setup complete")
	return version
}