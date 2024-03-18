<script lang="ts">
	import type { UserSettings } from '$lib/pb/user_data_pb';
	import { getSettingsListener, settingsStore } from '$lib/settings';
	import { userStore } from '$lib/user';
	import { Heading } from 'flowbite-svelte';
	import { onDestroy, onMount } from 'svelte';
	import { t } from 'svelte-i18n';
	import EventButtons from './EventButtons.svelte';
	import EventList from './EventList.svelte';
	import {Event} from "$lib/pb/event_pb"
	import { getEventsDashboardListener } from './events';
	import { dashboardUidStore } from './store';

	let events: Event[] = []
	let settings: UserSettings | undefined

	let eventUnsub: any;
	let settingsUnsub: any;

	onMount(() => {
		const uidUnsub = dashboardUidStore.subscribe(uid => {
			if(eventUnsub) eventUnsub()
			if(settingsUnsub) settingsUnsub()

			eventUnsub = getEventsDashboardListener(uid, e => {
				events = e
			});

			settingsUnsub = getSettingsListener(uid, s => {
				settings = s
			})
		})

		return () => {
			uidUnsub()
			eventUnsub()
			settingsUnsub()
		}
	})

</script>

<div class="flex flex-col space-y-10">
	{#if settings}
		<EventButtons bind:muted={settings.Events.IsMuted} bind:paused={settings.Events.IsPaused}/>
	{/if}

	<div class="flex flex-col justify-center space-x-2">
		<Heading tag="h5">{$t("latest_events")}</Heading>

		<EventList {events}/>
	</div>
</div>