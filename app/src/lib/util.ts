import {t} from 'svelte-i18n';
import { get } from 'svelte/store';

export const pbEnumToItems = (enumType: any): any[] => {
	const entries = JSON.parse(JSON.stringify(enumType))
	let items = [];
	for(const [k,v] of Object.entries(entries)) {
		const n = Number.parseInt(k);
		if (n != k) continue;

		let name = v as string;
		const formatter = get(t)
		items.push({
			value: n,
			name: formatter(name.toLowerCase())
		})
	}
	return items;
}