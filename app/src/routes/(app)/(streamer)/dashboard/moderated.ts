import { auth, rtdb } from '$lib/ext/firebase/firebase';
import type { User } from '$lib/pb/user_pb';
import { getUserByUid } from '$lib/user';
import {ref, onValue} from 'firebase/database';
import { writable, type Writable } from 'svelte/store';

export const moderatedUsersStore: Writable<User[]> = writable([]);

export const getModeratedUsersListener = (uid: string) => {
	const keyRef = ref(rtdb, `Moderators/To/${uid}`);
	return onValue(keyRef,  async(snapshot) => {
		const data = snapshot.val();
		if (!data) {
			moderatedUsersStore.set([]);
			return
		}

		const UIDs: string[] = Object.values(data);
		const promises = UIDs.map((uid: string) => {
			return getUserByUid(uid);
		})
		const moderators = await Promise.all(promises);

		moderatedUsersStore.set(moderators);
	});
}