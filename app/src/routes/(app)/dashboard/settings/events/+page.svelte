<script lang="ts">
	import { AuthLevel, Frequency } from '$lib/pb/stredono_pb';
	import { Accordion, AccordionItem, Button, Checkbox, Heading, Input, Label, Popover, Select } from 'flowbite-svelte';
	import { InfoCircleOutline, InfoCircleSolid } from 'flowbite-svelte-icons';
	import {slide} from 'svelte/transition';
	import InputMin from './InputMin.svelte';
	import AccordItem from './AccordItem.svelte';
	import { t } from 'svelte-i18n';

	$: authLevels = [
		{ value: AuthLevel.NONE, name: $t("disabled") },
		{ value: AuthLevel.EMAIL, name: $t("email_verification") },
		{ value: AuthLevel.OIDC, name: $t("twitch_login") }
	]
	let selectedAuthLevel = AuthLevel.NONE;

	$: followFrequency = [
		{ value: Frequency.NEVER, name: $t("never") },
		{ value: Frequency.ALWAYS, name: $t("always") },
		{ value: Frequency.EVERY_MINUTE, name: $t("every_minute") },
		{ value: Frequency.EVERY_FIVE_MINUTES, name: $t("every_5_minutes") },
		{ value: Frequency.EVERY_QUARTER_HOUR, name: $t("every_10_minutes") },
		{ value: Frequency.EVERY_HALF_HOUR, name: $t("every_half_hour") },
		{ value: Frequency.EVERY_HOUR, name: $t("every_hour") },
	]
	let selectedFreq = Frequency.NEVER;

</script>

<Heading tag="h2">{$t("events")}</Heading>
<div class="space-y-4 w-full p-4">
	<Heading tag="h4">{$t("tips")}</Heading>
	<div class="space-y-4 px-4">
		<Label>
			<div class="flex">
				{$t("require_auth")}
				<InfoCircleOutline class="ms-1.5"/>
				<Popover class="w-96">
					<div class="p-4">
						<p>{$t("require_auth_help")}</p>
					</div>
				</Popover>
			</div>
			<Select items={authLevels} bind:value={selectedAuthLevel}/>
		</Label>

		<InputMin label={$t("min_tip")} type="number" placeholder="5.00"/>

		<Checkbox checked={true}>
			{$t("tip_availability")}
			<InfoCircleOutline class="ms-1.5 mt-0.5"/>
			<Popover class="w-96">
				<div class="p-4">
					<p>{$t("tip_availability_help")}</p>
				</div>
			</Popover>
		</Checkbox>
	</div>

	<Heading tag="h4">{$t("twitch")}</Heading>

	<div class="space-y-4 px-4">
		<Accordion class="w-full">
			<AccordItem header={$t("subs")}>
				<InputMin label={$t("min_amount")} type="number" placeholder="1" />
			</AccordItem>
			<AccordItem header={$t("sub_gifts")}>
				<InputMin label={$t("min_amount")} type="number" placeholder="1" />
			</AccordItem>
			<AccordItem header={$t("follows")}>
				<Label>
					{$t("frequency")}
					<Select items={followFrequency} bind:value={selectedFreq}/>
				</Label>
			</AccordItem>
			<AccordItem header={$t("cheers")}>
				<InputMin label={$t("min_bits")} type="number" placeholder="1" />
			</AccordItem>
			<AccordItem header={$t("raids")}>
				<InputMin label={$t("min_viewers")} type="number" placeholder="10" />
			</AccordItem>
		</Accordion>
	</div>

	<div class="pt-4">
		<Button>{$t("save")}</Button>
	</div>
</div>