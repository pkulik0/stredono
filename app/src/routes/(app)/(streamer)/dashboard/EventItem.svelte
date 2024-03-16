<script lang="ts">
	import { Action, changeEventState } from '$lib/events';
	import { Button, Dropdown, DropdownItem, ListgroupItem, P } from 'flowbite-svelte';
	import {
		CheckOutline, CloseOutline,
		DotsVerticalSolid, EyeSlashOutline,
		EyeSlashSolid,
		RedoOutline,
		RefreshOutline,
		ShieldSolid, TrashBinOutline
	} from 'flowbite-svelte-icons';
	import { t } from 'svelte-i18n';
	import { Event, EventType } from '$lib/pb/event_pb';

	export let event: Event;
</script>

<ListgroupItem class="flex items-center justify-between space-x-4 rtl:space-x-reverse">
	<div>
		<P class="text-sm text-gray-400 dark:text-gray-500">
			{new Date(Number(event.Timestamp) * 1000).toLocaleString()}
		</P>
		<P class="text-sm text-gray-500 dark:text-gray-400">
			{$t(JSON.parse(JSON.stringify(EventType))[event.Type].toLowerCase())}
			{#if event.Data.Value}
				({event.Data.Value})
			{/if}
		</P>
		<P class="text-lg font-bold text-gray-900 dark:text-white">
			{event.SenderName}
		</P>
	</div>


	<P class="text-lg">
		{event.Data.Message}
	</P>

	<div class="flex">
		{#if !event.IsApproved}
			<Button color="green" outline size="xs" on:click={() => changeEventState(Action.Approve, event.ID, 0)}>
				<CheckOutline class="me-1"/>
				{$t("approve")}
			</Button>
		{:else}
			{#if event.WasShown}
				<Button size="xs" outline on:click={() => changeEventState(Action.Rerun, event.ID, 0)}>
					<RefreshOutline class="me-1"/>
					{$t("rerun")}
				</Button>
			{:else}
				<Button size="xs" color="red" outline on:click={() => changeEventState(Action.Cancel, event.ID, 0)}>
					<CloseOutline class="me-1"/>
					{$t("dont_show")}
				</Button>
			{/if}
		{/if}
	</div>

</ListgroupItem>