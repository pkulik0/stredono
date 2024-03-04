<script lang="ts">
	import type { Voice } from '$lib/pb/stredono_pb';
	import { Button, Label, Select } from 'flowbite-svelte';
	import { PauseSolid, PlaySolid } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';

	export let voices: Voice[];
	export let value: Voice|undefined;
	export let placeholder = "Choose a voice...";

	$: items = voices.map((voice) => ({ value: voice, name: voice.Name })).sort((a, b) => a.name.localeCompare(b.name));

	let audio: HTMLAudioElement = new Audio(value?.SampleUrl)
	$: if (value) audio.src = value.SampleUrl;
	let isPlaying = false;

	const clickPlay = () => {
		if (isPlaying) {
			audio.pause();
			audio.currentTime = 0;
			return;
		}
		if (!value) return;
		audio.play();
	};

	const onChange = () => {
		audio.pause()
		isPlaying = false;
	};

	onMount(() => {
		audio.addEventListener('ended', () => {
			isPlaying = false;
		});
		audio.addEventListener('play', () => {
			isPlaying = true;
		});
		audio.addEventListener('pause', () => {
			isPlaying = false;
		});
	})
</script>

<Label>
	Voice
	<div class="flex flex-row space-x-2">
		<Select on:change={onChange} items={items} bind:value {placeholder}/>
		<Button outline on:click={clickPlay} class="max-w-36 w-full">
			{#if isPlaying}
				<PauseSolid class="me-1.5"/>
				Stop
			{:else}
				<PlaySolid class="me-1.5"/>
				Listen
			{/if}
		</Button>
	</div>
</Label>