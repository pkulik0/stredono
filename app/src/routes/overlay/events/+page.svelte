<script lang="ts">
	import AlertViewer from '$lib/comp/AlertViewer.svelte';
	import { settingsStore } from '$lib/events_settings';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { keyStore, uidStore } from '../stores';
	import { confirmEvent, eventsStore, getEventsListener } from './listener';
	import {Event} from "$lib/pb/event_pb";

	let eventsUnsub: (() => void) | undefined;
	onMount(() => {
		eventsStore.subscribe(e => {
			console.log("events", e)
		})
		settingsStore.subscribe(s => {
			console.log("settings", s)
		})

		return uidStore.subscribe(uid => {
			if(uid) {
				eventsUnsub = getEventsListener(uid);
			} else if(eventsUnsub) {
				eventsUnsub();
			}
		})
	})

	const onShown = async (e: Event) => {
		const key = get(keyStore);
		if(!key) {
			throw new Error("No key")
		}
		await confirmEvent(key, e.ID)
		console
		event = undefined;
		setTimeout(() => {
			event = $eventsStore[0];
		}, 5000)
	}

	let event: Event|undefined = $eventsStore[0];
	$: if (!event) {
		event = $eventsStore[0];
	}
</script>

<AlertViewer alerts={$settingsStore?.Alerts ?? []} {event} {onShown}/>