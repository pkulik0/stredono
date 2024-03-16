<script lang="ts">
	import { settingsStore } from '$lib/settings.js';
	import { Checkbox, Input, Label, Popover } from 'flowbite-svelte';
	import { InfoCircleOutline } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { t } from 'svelte-i18n';

	export let requireApproval: boolean;
	export let minDisplayTime: number;

	let minDurationString: string = "";
	$: if(minDurationString) {
		minDisplayTime = parseInt(minDurationString);
	}

	onMount(() => {
		minDurationString = minDisplayTime.toString();
	})
</script>

<div class="py-4 space-y-4">
	<Label>
		{$t("min_display_time")}
		<Input type="number" bind:value={minDurationString} min="1" max="30"/>
	</Label>

	<Checkbox bind:checked={requireApproval}>
		{$t("require_approval")}
		<InfoCircleOutline class="w-4 h-4 ms-1" />
		<Popover>
			<div class="p-4 w-80">
				{$t("require_approval_help")}
			</div>
		</Popover>
	</Checkbox>
</div>