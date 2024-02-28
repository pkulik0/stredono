import { env } from '$env/dynamic/public';
import TerraformOutput from './terraform_output.json';

export let terraformOutput = TerraformOutput;

export let isLocal = env.PUBLIC_IS_LOCAL === 'true';
export let isDev = isLocal || env.PUBLIC_IS_DEV === 'true';
if(isDev) {
	const mapKv = new Map(Object.entries(terraformOutput.FunctionUrls));
	for(let [name, url] of mapKv) {
		mapKv.set(name, `http://localhost:8080/${name}`);
	}
	terraformOutput.FunctionUrls = Object.fromEntries(mapKv);
	console.log("Functions URLs set to localhost - dev mode");
}