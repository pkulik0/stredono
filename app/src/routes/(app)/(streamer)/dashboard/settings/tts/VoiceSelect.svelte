<script lang="ts">
	import type { Voice } from '../../../../../../lib/pb/tts_pb';
	import { Button, Label, Select } from 'flowbite-svelte';
	import { PauseSolid, PlaySolid } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { t } from 'svelte-i18n';

	export let voices: Voice[];
	export let value: string|undefined;
	$: voice = voices.find((voice) => voice.Id === value);
	export let placeholder = $t("voice_placeholder");

	$: items = voices.map((voice) => ({ value: voice.Id, name: voice.Name })).sort((a, b) => a.name.localeCompare(b.name));

	let audio: HTMLAudioElement = new Audio(voice?.SampleUrl)
	$: audio.src = voice?.SampleUrl ?? "";
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
	{$t("voice")}
	<div class="flex flex-row space-x-2">
		<Select on:change={onChange} items={items} bind:value {placeholder}/>
		<Button outline on:click={clickPlay} class="max-w-36 w-full">
			{#if isPlaying}
				<PauseSolid class="me-1.5"/>
				{$t("stop")}
			{:else}
				<PlaySolid class="me-1.5"/>
				{$t("listen")}
			{/if}
		</Button>
	</div>
</Label>