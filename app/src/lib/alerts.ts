import { auth, db } from '$lib/ext/firebase/firebase';
import { type Alert, UsersAlerts } from '$lib/pb/alert_pb';
import { Event } from '$lib/pb/event_pb';
import { terraformOutput } from '$lib/constants';
import axios from 'axios';

export const saveAlert = async (alert: Alert) => {
	const user = auth.currentUser;
	if (!user) throw new Error('User not logged in');

	try {
		const res = await axios.post(terraformOutput.FunctionsUrl + "/AlertAdd", alert.toBinary(), {
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

export const eventToAlert = (event: Event, alerts: Alert[], isTest: boolean): Alert|undefined => {
	if(isTest) {
		return alerts[0];
	}

	for(const alert of alerts) {
		if(alert.EventType !== event.Type) continue;

		const value = Number.parseFloat(event.Data.Value);
		if(value < alert.Min) continue;
		if(alert.Max !== undefined) {
			if(value > alert.Max) continue;
		}

		return alert;
	}
	return undefined;
}