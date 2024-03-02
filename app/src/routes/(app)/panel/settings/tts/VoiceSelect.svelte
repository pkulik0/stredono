<script lang="ts">
	import type { Voice } from '$lib/pb/stredono_pb';
	import { Button, Label, Select } from 'flowbite-svelte';
	import { PauseSolid, PlaySolid } from 'flowbite-svelte-icons';

	export let voices: Voice[];
	export let value: Voice|undefined;
	export let placeholder = "Choose a voice...";

	$: items = voices.map((voice) => ({ value: voice, name: voice.Name })).sort((a, b) => a.name.localeCompare(b.name));

	let audio: HTMLAudioElement|undefined;
	let isPlaying = false;
	const clickPlay = () => {
		if (isPlaying) {
			audio?.pause();
			isPlaying = false;
			return;
		}
		if (!value) return;
		isPlaying = true;

		audio = new Audio(value.SampleUrl);
		audio.onended = () => {
			isPlaying = false;
		};
		audio.play();
	};
</script>

<Label>
	Voice
	<div class="flex flex-row space-x-2">
		<Select items={items} bind:value {placeholder}/>
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