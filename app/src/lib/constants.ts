import { env } from '$env/dynamic/public';
import TerraformOutput from './terraform_output.json';

export let terraformOutput = TerraformOutput;

export let isLocal = env.PUBLIC_IS_LOCAL === 'true';
export let isDev = isLocal || env.PUBLIC_IS_DEV === 'true';
if(isDev) {
	terraformOutput.FunctionsUrl = `http://localhost:8080`;
}