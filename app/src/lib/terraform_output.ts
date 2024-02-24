import { PUBLIC_IS_LOCAL } from '$env/static/public';
import TerraformOutput from './terraform_output.json';

export let terraformOutput = TerraformOutput;

if(PUBLIC_IS_LOCAL) {
	const mapKv = new Map(Object.entries(terraformOutput.FunctionUrls));
	for(let [name, url] of mapKv) {
		mapKv.set(name, `http://localhost:8080/${name}`);
	}
	terraformOutput.FunctionUrls = Object.fromEntries(mapKv);
}