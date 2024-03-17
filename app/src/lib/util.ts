import { Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
import { Currency } from '$lib/pb/enums_pb';
import { EventType } from '$lib/pb/event_pb';
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

export const currencyToSymbol = (currency: Currency): string => {
	switch(currency) {
		case Currency.PLN:
			return "zł";
		default:
			return "?";
	}
}