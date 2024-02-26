import { AnimationType } from '$lib/pb/stredono_pb';

export const pbEnumToItems = (enumType: any): any[] => {
	const entries = JSON.parse(JSON.stringify(AnimationType))
	let items = [];
	for(const [k,v] of Object.entries(entries)) {
		const n = Number.parseInt(k);
		if (n != k) continue;

		let name = v as string;
		name = name.split("_").map((s) => s.charAt(0).toUpperCase() + s.slice(1).toLowerCase()).join(" ");

		items.push({
			value: n as AnimationType,
			name: name
		})
	}
	return items;
}