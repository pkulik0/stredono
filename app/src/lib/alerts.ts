import { auth, db } from '$lib/ext/firebase/firebase';
import { type Alert, UsersAlerts, Event } from '$lib/pb/stredono_pb';
import { terraformOutput } from '$lib/constants';
import axios from 'axios';
import { doc, onSnapshot } from 'firebase/firestore';
import { writable, type Writable } from 'svelte/store';

export const alertsStore: Writable<UsersAlerts|undefined|null> = writable(undefined)

export const getAlertsListener = async (uid: string) => {
	return onSnapshot(doc(db, "alerts", uid), (snapshot) => {
		if(!snapshot.exists()) {
			alertsStore.set(null)
			return
		}
		const usersAlerts  = UsersAlerts.fromJson(snapshot.data())
		alertsStore.set(usersAlerts)
	})
}

export const saveAlert = async (alert: Alert) => {
	const user = auth.currentUser;
	if (!user) throw new Error('User not logged in');

	try {
		const res = await axios.post(terraformOutput.FunctionUrls.AlertAdd, alert.toBinary(), {
			headers: {
				'Content-Type': 'application/octet-stream',
				'Authorization': 'Bearer ' + await user.getIdToken()
			}
		});
		console.log(res.data);
	} catch (e) {
		console.error(e);
	}
}

export const eventToAlert = (event: Event, alerts: Alert[]): Alert|undefined => {
	for(const alert of alerts) {
		if(alert.EventType !== event.Type) continue;

		const value = Number.parseFloat(event.Data.Value);
		if(value < alert.Min) continue;
		if(value > alert.Max) continue;

		return alert;
	}
	return undefined;
}