<script lang="ts">
	import { Tier, User, type Voice } from '$lib/pb/stredono_pb';
	import { getVoices } from '$lib/tts';
	import { userStore } from '$lib/user';
	import {
		Alert,
		Button,
		Checkbox,
		Heading,
		Helper,
		Hr,
		ImagePlaceholder,
		Label,
		P,
		Select,
		TextPlaceholder
	} from 'flowbite-svelte';
	import { InfoCircleSolid, PlaySolid } from 'flowbite-svelte-icons';
	import { onDestroy, onMount } from 'svelte';
	import VoiceSelect from './VoiceSelect.svelte';

	const langName = new Intl.DisplayNames(['en'], {type: 'language'})

	let voicesPlus: Voice[] = [];
	$: languagesPlus = voicesPlus.length > 0 ? voicesPlus[0].Languages.map(l => langName.of(l) || l).sort() : Array(30).fill("???");
	let selectedPlus: Voice|undefined;

	let voicesBasic: Map<string, Voice[]> = new Map();
	$: languagesBasic = Array.from(voicesBasic.keys()).map(l => ({
		value: l,
		name: langName.of(l) || l
	})).sort((a, b) => a.name.localeCompare(b.name))
	let selectedLanguage: string = "pl";
	let selectedBasic: Voice|undefined;

	let enablePlus = true;

	const fetchVoices = async () => {
		const res = await getVoices()
		voicesPlus = res.voicesPlus
		voicesBasic = res.voicesBasic
	}

	$: {
		if(!selectedBasic) selectedBasic = voicesBasic.get(selectedLanguage)?.find(v => v.Id === user?.VoiceBasic);
		if(!selectedPlus) selectedPlus = voicesPlus.find(v => v.Id === user?.VoicePlus);
	}

	let user: User|undefined;
	onMount(() => {
		fetchVoices()

		return userStore.subscribe(u => {
			user = u || undefined;
		})
	})
</script>

<Heading tag="h2">Text-to-Speech</Heading>
<div class="space-y-10 p-4">
	<div class="space-y-6 flex flex-col">
		<Heading tag="h4">Plus</Heading>

		<Alert class="!items-start" color="gray">
			<span slot="icon">
				<InfoCircleSolid class="w-5 h-5"/>
				<span class="sr-only">Info</span>
			</span>
			<p class="font-medium">
				Plus voices are capable of speaking multiple languages even in one sentence. They might still have a slight accent or work better in some languages than others. Pick the one that suits your needs.
				<br/>
				Supported languages:
			</p>
			<ul class="mt-1.5 ms-4 list-disc list-inside grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-1">
				{#each languagesPlus as l}
					<li>{l}</li>
				{/each}
			</ul>
		</Alert>

		<VoiceSelect voices={voicesPlus} bind:value={selectedPlus}/>

		<Checkbox bind:checked={enablePlus}>Enable TTS Plus</Checkbox>
	</div>

	<div class="space-y-6 flex flex-col">
		<Heading tag="h4">Basic</Heading>

		<Label>
			Language
			<Select items={languagesBasic} bind:value={selectedLanguage}/>
		</Label>

		<VoiceSelect voices={voicesBasic.get(selectedLanguage) || []} bind:value={selectedBasic}/>
	</div>

	<Button>
		Save
	</Button>
</div>