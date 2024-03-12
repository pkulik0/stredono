<script lang="ts">
	import ParticlesBackground from '$lib/comp/ParticlesBackground.svelte';
	import { onMount } from 'svelte';

	onMount(() => {
		// Has to be here to not include /overlay/*
		document.body.classList.add('bg-gray-200')
		document.body.classList.add('dark:bg-gray-900')
	})
</script>

<slot/>

<ParticlesBackground/>