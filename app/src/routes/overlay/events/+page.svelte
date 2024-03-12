<script lang="ts">
	import EventViewer from '$lib/comp/EventViewer.svelte';
	import { settingsStore } from '$lib/events_settings';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { keyStore, uidStore } from '../stores';
	import { confirmEvent, eventsStore, getEventsListener } from './listener';
	import {Event} from "$lib/pb/event_pb";

	let events: Event[] = [];

	let eventsUnsub: (() => void) | undefined;
	onMount(() => {
		eventsStore.subscribe(e => {
			events = e;
		})

		return uidStore.subscribe(uid => {
			if(uid) {
				eventsUnsub = getEventsListener(uid);
			} else if(eventsUnsub) {
				eventsUnsub();
			}
		})
	})
</script>

<EventViewer alerts={$settingsStore?.Alerts ?? []} {events}/>