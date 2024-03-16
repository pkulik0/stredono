import { rtdb } from '$lib/ext/firebase/firebase';
import { child, onDisconnect, push, ref, set } from 'firebase/database';

export const setupOnlineStatus = (key: string) => {
	const sessionsRef = ref(rtdb, `Sessions/${key}`)

	const sessionKey = push(sessionsRef, Date.now()).key;
	if (!sessionKey) {
		throw new Error("Failed to create session");
	}
	const sessionRef = child(sessionsRef, sessionKey);

	setInterval(async () => {
		await set(sessionRef, Date.now())
	}, 5000)

	onDisconnect(sessionRef).remove() // fine to ignore promise
}