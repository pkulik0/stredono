<script lang="ts">
	import { AuthLevel, Frequency } from '$lib/pb/stredono_pb';
	import { Accordion, AccordionItem, Button, Checkbox, Heading, Input, Label, Popover, Select } from 'flowbite-svelte';
	import { InfoCircleOutline, InfoCircleSolid } from 'flowbite-svelte-icons';
	import {slide} from 'svelte/transition';
	import InputMin from './InputMin.svelte';
	import AccordItem from './AccordItem.svelte';

	const authLevels = [
		{ value: AuthLevel.NONE, name: 'Disabled' },
		{ value: AuthLevel.EMAIL, name: 'Email verification' },
		{ value: AuthLevel.OIDC, name: 'Twitch login' }
	]
	let selectedAuthLevel = AuthLevel.NONE;

	const followFrequency = [
		{ value: Frequency.NEVER, name: "Never" },
		{ value: Frequency.ALWAYS, name: "Always" },
		{ value: Frequency.EVERY_MINUTE, name: "Every minute" },
		{ value: Frequency.EVERY_FIVE_MINUTES, name: "Every 5 minutes" },
		{ value: Frequency.EVERY_QUARTER_HOUR, name: "Every 15 minutes" },
		{ value: Frequency.EVERY_HALF_HOUR, name: "Every half an hour" },
		{ value: Frequency.EVERY_HOUR, name: "Every hour" },
	]
	let selectedFreq = Frequency.NEVER;

</script>

<Heading tag="h2">Events</Heading>
<div class="space-y-4 w-full p-4">
	<Heading tag="h4">Tips</Heading>
	<div class="space-y-4 px-4">
		<Label>
			<div class="flex">
				Require authentication
				<InfoCircleOutline class="ms-1.5"/>
				<Popover class="w-96">
					<div class="p-4">
						<p>If enabled, before tipping, users will have to verify their email or log in with Twitch.</p>
					</div>
				</Popover>
			</div>
			<Select items={authLevels} bind:value={selectedAuthLevel}/>
		</Label>

		<InputMin label="Minimum tip" type="number" placeholder="5.00"/>

		<Checkbox checked={true}>
			Show availability on the tipping page
			<InfoCircleOutline class="ms-1.5 mt-0.5"/>
			<Popover class="w-96">
				<div class="p-4">
					<p>If you go offline or temporarily turn off alerts, sound or TTS, a message will be displayed on the tipping page.</p>
				</div>
			</Popover>
		</Checkbox>
	</div>

	<Heading tag="h4">Twitch</Heading>

	<div class="space-y-4 px-4">
		<Accordion class="w-full">
			<AccordItem header="Subscriptions">
				<InputMin label="Minimum months" type="number" placeholder="1" />
			</AccordItem>
			<AccordItem header="Gifts">
				<InputMin label="Minimum amount" type="number" placeholder="1" />
			</AccordItem>
			<AccordItem header="Follows">
				<Label>
					Frequency
					<Select items={followFrequency} bind:value={selectedFreq}/>
				</Label>
			</AccordItem>
			<AccordItem header="Cheers">
				<InputMin label="Minimum bits" type="number" placeholder="1" />
			</AccordItem>
			<AccordItem header="Raids">
				<InputMin label="Minimum viewers" type="number" placeholder="10" />
			</AccordItem>
		</Accordion>
	</div>

	<div class="pt-4">
		<Button>Save</Button>
	</div>
</div>