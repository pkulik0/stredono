import { terraformOutput } from '$lib/constants';
import { db, rtdb } from '$lib/ext/firebase/firebase';
import axios from 'axios';
import { ref, onDisconnect, push, child, set } from 'firebase/database';
import { collection, onSnapshot, query, where, limit, orderBy } from 'firebase/firestore';
import { writable, type Writable } from 'svelte/store';
import { Event } from '$lib/pb/event_pb';

export const eventsStore: Writable<Event[]> = writable([]);

export const getEventsOverlayListener = (uid: string) => {
	const q = query(collection(db, "events"),
		where("Uid", "==" , uid),
		where("IsApproved", "==", true),
		where("WasShown", "==", false),
		orderBy("Timestamp", "asc"),
		limit(3)
	);
	return onSnapshot(q, (snapshot) => {
		eventsStore.set(snapshot.docs.map(doc => Event.fromJson(doc.data())));
	})
}

export const confirmEvent = async (key: string, eventId: string) => {
	const res = await axios.post(terraformOutput.FunctionsUrl+`/EventChangeState?key=${key}&event=${eventId}&action=confirm`, {}, {
		headers: {
			'Content-Type': 'application/json'
		}
	})
	if (res.status !== 200) {
		throw new Error("Failed to confirm event");
	}
}