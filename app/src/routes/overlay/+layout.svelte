<script lang="ts">
	import { getSettingsListener } from '$lib/events_settings';
	import { rtdb } from '$lib/ext/firebase/firebase';
	import { onMount } from 'svelte';
	import { ref, get } from "firebase/database";
	import { keyStore, uidStore } from './stores';

	onMount(async () => {
		const queryArgs = new URLSearchParams(window.location.search)
		const key = queryArgs.get("key")
		if(!key) {
			// error
			console.log("No key")
			return
		}
		keyStore.set(key)

		const snapshot = await get(ref(rtdb, `Users/Overlay/${key}`))
		if(!snapshot.exists()) {
			// error
			console.log("No snapshot")
			return
		}
		const uid = snapshot.val() as string;
		uidStore.set(uid)
		getSettingsListener(uid)
	})
</script>

<slot/>