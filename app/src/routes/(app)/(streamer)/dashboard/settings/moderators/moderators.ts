import { auth, rtdb } from '$lib/ext/firebase/firebase';
import type { User } from '$lib/pb/user_pb';
import { getUserByUid } from '$lib/user';
import {ref, onValue} from 'firebase/database';
import { writable, type Writable } from 'svelte/store';

export const moderatorsStore: Writable<User[]> = writable([]);

export const getModeratorsListener = (uid: string) => {
	const keyRef = ref(rtdb, `Moderators/From/${uid}`);
	return onValue(keyRef,  async(snapshot) => {
		const data = snapshot.val();
		if (!data) {
			moderatorsStore.set([]);
			return
		}

		const UIDs: string[] = Object.values(data);
		const promises = UIDs.map((uid: string) => {
			return getUserByUid(uid);
		})
		const moderators = await Promise.all(promises);

		moderatorsStore.set(moderators);
	});
}