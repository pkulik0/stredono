import { terraformOutput } from '$lib/constants';
import { auth, db } from '$lib/ext/firebase/firebase';
import axios from 'axios';
import { collection, limit, onSnapshot, orderBy, query, where } from 'firebase/firestore';
import { type Writable, writable } from 'svelte/store';
import {Event} from "$lib/pb/event_pb";

export const getEventsDashboardListener = (uid: string, onChange: (events: Event[]) => void) => {
	const q = query(collection(db, "events"),
		where("Uid", "==", uid),
		orderBy("Timestamp", "desc"),
		limit(50)
	);

	return onSnapshot(q, (snapshot) => {
		onChange(snapshot.docs.map(doc => Event.fromJson(doc.data())))
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

export const changeEventState = async (uid: string, eventId: string, action: Action, minutes: number) => {
	const user = auth.currentUser;
	if (!user) {
		throw new Error("User not authenticated");
	}
	const token = await user.getIdToken();

	const url = terraformOutput.FunctionsUrl + `/EventChangeState?uid=${uid}&action=${action}&event=${eventId}&minutes=${minutes}`
	const res = await axios.get(url, {
		headers: {
			'Authorization': 'Bearer ' + token
		},
	})
	if (res.status !== 200) {
		throw new Error("Failed to change event state");
	}
}