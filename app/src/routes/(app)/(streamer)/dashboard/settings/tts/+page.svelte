<script lang="ts">
	import { saveSettings, settingsStore } from '$lib/settings';
	import { Tier } from '$lib/pb/enums_pb';
	import type { Voice } from '$lib/pb/tts_pb';
	import { getVoices } from '$lib/tts';
	import { userStore } from '$lib/user';
	import {
		Alert,
		Button,
		Checkbox,
		Heading,
		Label,
		Select,
	} from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import VoiceSelect from './VoiceSelect.svelte';
	import { locale, t } from 'svelte-i18n';

	$: langName = new Intl.DisplayNames([$locale || 'en'], {type: 'language'})

	let voicesPlus: Voice[] = [];
	$: languagesPlus = voicesPlus.length > 0 ? voicesPlus[0].Languages.map(l => langName.of(l) || l).sort() : Array(30).fill("???");

	let voicesBasic: Map<string, Voice[]> = new Map();
	$: languagesBasic = Array.from(voicesBasic.keys()).map(l => ({
		value: l,
		name: langName.of(l) || l
	})).sort((a, b) => a.name.localeCompare(b.name))
	let selectedLanguage: string = "pl";

	const fetchVoices = async () => {
		const res = await getVoices()
		voicesPlus = res.voicesPlus
		voicesBasic = res.voicesBasic
	}

	onMount(() => {
		fetchVoices()
	})

	const save = async () => {
		const user = get(userStore)
		if(!user) return

		await saveSettings(user.Uid)
	}
</script>

{#if $settingsStore !== undefined}
	<Heading tag="h2">{$t("tts")}</Heading>
	<div class="space-y-10 p-4">
		<div class="space-y-6 flex flex-col">
			<Heading tag="h4">{$t("basic")}</Heading>

			<Label>
				{$t("language")}
				<Select items={languagesBasic} bind:value={selectedLanguage}/>
			</Label>

			<VoiceSelect voices={voicesBasic.get(selectedLanguage) || []} bind:value={$settingsStore.TTS.VoiceIdBasic}/>
		</div>

		<div class="space-y-6 flex flex-col">
			<Heading tag="h4">{$t("plus")}</Heading>

			<Alert class="!items-start" color="gray">
			<span slot="icon">
				<InfoCircleSolid class="w-5 h-5"/>
				<span class="sr-only">{$t("info")}</span>
			</span>
				<p class="font-medium">
					{$t("tts_plus_info")}
					<br/>
					{$t("supported_languages")}:
				</p>
				<ul class="mt-1.5 ms-4 list-disc list-inside grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-1">
					{#each languagesPlus as l}
						<li>{l}</li>
					{/each}
				</ul>
			</Alert>

			<VoiceSelect voices={voicesPlus} bind:value={$settingsStore.TTS.VoiceIdPlus}/>

			<Checkbox bind:checked={$settingsStore.TTS.UsePlus}>{$t("tts_plus_enable")}</Checkbox>
		</div>


		<Button on:click={save}>
			{$t("save")}
		</Button>
	</div>
{/if}