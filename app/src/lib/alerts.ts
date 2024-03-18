import { saveSettings, settingsStore } from '$lib/settings';
import { auth, db } from '$lib/ext/firebase/firebase';
import {  Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
import { Event, EventType } from '$lib/pb/event_pb';
import { get } from 'svelte/store';
import {v4 as uuid} from 'uuid';

export const saveAlert = async (alert: Alert) => {
	const user = auth.currentUser;
	if (!user) throw new Error('User not logged in');

	let settings = get(settingsStore);
	if (!settings) {
		throw new Error("Settings not found");
	}

	if(alert.ID === "") {
		alert.ID = uuid().replace(/-/g, '');
		settings.Alerts.push(alert);
	}

	await saveSettings(user.uid)
}

export const removeAlert = async (alert: Alert) => {
	const user = auth.currentUser;
	if (!user) throw new Error('User not logged in');

	let settings = get(settingsStore);
	if (!settings) {
		throw new Error("Settings not found");
	}

	settings.Alerts = settings.Alerts.filter(a => a.ID !== alert.ID);
	await saveSettings(user.uid)
}

export const getDefaultAlert = (eventType: EventType) => {
	return Alert.fromJson({
		ID: "",
		EventType: eventType,
		Min: 1,
		Max: 10,
		TextColor: "#FFFFFF",
		AccentColor: "#2381f8",
		Animation: AnimationType.PULSE,
		AnimationSpeed: Speed.MEDIUM,
		Alignment: Alignment.JUSTIFY,
		TextPosition: Position.BOTTOM,
	});
}

export const eventToAlert = (event: Event, alerts: Alert[]): Alert => {
	let filteredAlerts = alerts.filter(a => a.EventType === event.Type)

	const valueStr = event.Data.Value;
	if (valueStr !== undefined) {
		const value = parseFloat(valueStr);
		filteredAlerts = filteredAlerts.filter(a => value >= a.Min && (a.Max !== undefined ? value <= a.Max  : true));
	}

	if (filteredAlerts.length > 0) {
		const i = Math.floor(Math.random() * filteredAlerts.length);
		return filteredAlerts[i];
	}

	return getDefaultAlert(event.Type);
}