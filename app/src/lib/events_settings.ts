import { rtdb } from '$lib/ext/firebase/firebase';
import { UserSettings } from '$lib/pb/user_data_pb';
import { ref, onValue, child } from "firebase/database";
import { type Writable, writable } from 'svelte/store';

export const settingsStore: Writable<UserSettings|undefined> = writable(undefined);

export const getSettingsListener = (uid: string) => {
	return onValue(ref(rtdb, `Data/${uid}/Settings`), (snapshot) => {
		const data = snapshot.val();
		if (!data) {
			throw new Error("Settings not found");
		}
		settingsStore.set(UserSettings.fromJson(data));
	});
}