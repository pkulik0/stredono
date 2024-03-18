import { rtdb } from '$lib/ext/firebase/firebase';
import { UserSettings } from '$lib/pb/user_data_pb';
import { ref, onValue, set } from "firebase/database";
import { get, type Writable, writable } from 'svelte/store';

export const settingsStore: Writable<UserSettings|undefined> = writable(undefined);

const getSettingsRef = (uid: string) => {
	return ref(rtdb, `Data/${uid}/Settings`)
}

export const getSettingsListener = (uid: string, onChange: (settings: UserSettings) => void) => {
	return onValue(getSettingsRef(uid), (snapshot) => {
		const data = snapshot.val();
		if (!data) {
			throw new Error("Settings not found");
		}
		onChange(UserSettings.fromJson(data))
	});
}

export const saveSettings = async (uid: string) => {
	const settings = get(settingsStore);
	if (!settings) {
		throw new Error("Settings not found");
	}

	const data = settings.toJson({
		useProtoFieldName: true,
	});
	await set(getSettingsRef(uid), data);
}