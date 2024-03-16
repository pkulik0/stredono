import { terraformOutput } from '$lib/constants';
import { auth, db } from '$lib/ext/firebase/firebase';
import axios from 'axios';
import { collection, limit, onSnapshot, orderBy, query, where } from 'firebase/firestore';
import { type Writable, writable } from 'svelte/store';
import {Event} from "$lib/pb/event_pb";

export const eventsStore: Writable<Event[]> = writable([]);

export const getEventsDashboardListener = (uid: string) => {
	const q = query(collection(db, "events"),
		where("Uid", "==", uid),
		orderBy("Timestamp", "desc"),
		limit(50)
	);

	return onSnapshot(q, (snapshot) => {
		eventsStore.set(snapshot.docs.map(doc => Event.fromJson(doc.data())))
	});
}

export enum Action {
	Cancel = "cancel",
	Rerun = "rerun",
	Approve = "approve",
	Resume = "resume",
	Pause = "pause",
	Mute = "mute",
	Unmute = "unmute",
}

export const changeEventState = async (action: Action, eventId: string, minutes: number) => {
	const user = auth.currentUser;
	if (!user) {
		throw new Error("User not authenticated");
	}
	const token = await user.getIdToken();

	const url = terraformOutput.FunctionsUrl + `/EventChangeState?uid=${user.uid}&action=${action}&event=${eventId}&minutes=${minutes}`
	const res = await axios.get(url, {
		headers: {
			'Authorization': 'Bearer ' + token
		},
	})
	if (res.status !== 200) {
		throw new Error("Failed to change event state");
	}
}