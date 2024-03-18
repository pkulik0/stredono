import { auth, rtdb } from '$lib/ext/firebase/firebase';
import {ref, onValue} from 'firebase/database';
import { writable, type Writable } from 'svelte/store';

export const overlayKeyStore: Writable<string> = writable('');

export const getKeyListener =  (uid: string) => {
	const keyRef = ref(rtdb, `Data/${uid}/OverlayKey`);
	return onValue(keyRef, (snapshot) => {
		const data = snapshot.val();
		overlayKeyStore.set(data || '');
	});
}