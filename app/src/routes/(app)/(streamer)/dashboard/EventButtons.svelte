<script lang="ts">
	import { Action, changeEventState } from '$lib/events';
	import { Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import {
		CloseSolid,
		PauseSolid,
		PlaySolid, RefreshOutline,
		VolumeDownSolid, VolumeUpSolid
	} from 'flowbite-svelte-icons';
	import {t} from "svelte-i18n";

	export let muted: boolean;
	export let paused: boolean;

	let isDropdownOpen = false;
</script>

<div class="flex space-x-4">
	{#if muted}
		<Button size="xl" class="w-full"  on:click={() => changeEventState(Action.Unmute, "", 0)} >
			<VolumeUpSolid class="w-6 h-6 me-2"/>
			{$t("unmute")}
		</Button>
	{:else}
		<Button size="xl" color="alternative" outline class="w-full"  on:click={() => changeEventState(Action.Mute, "", 0)} >
			<VolumeDownSolid class="me-2 h-6 w-6"/>
			{$t("mute")}
		</Button>
	{/if}

	{#if paused}
		<Button size="xl" color="green" class="w-full" on:click={() => changeEventState(Action.Resume, "", 0)}>
			<PlaySolid class="w-6 h-6 me-2"/>
			{$t("resume")}
		</Button>
	{:else}
		<Button size="xl" color="red" outline class="w-full" on:click={() => changeEventState(Action.Pause, "", 0)} >
			<PauseSolid class="w-6 h-6 me-2"/>
			{$t("pause")}
		</Button>
	{/if}

	<Button size="xl" outline class="w-full" >
		<RefreshOutline class="w-6 h-6 me-2"/>
		{$t("rerun")}
	</Button>
	<Dropdown bind:open={isDropdownOpen}>
		<DropdownItem on:click={() => { changeEventState(Action.Rerun, "", 1); isDropdownOpen = false }} >{$t("last_min")}</DropdownItem>
		<DropdownItem on:click={() => { changeEventState(Action.Rerun, "", 3); isDropdownOpen = false }} >{$t("last_3_min")}</DropdownItem>
		<DropdownItem on:click={() => { changeEventState(Action.Rerun, "", 5); isDropdownOpen = false }} >{$t("last_5_min")}</DropdownItem>
		<DropdownItem on:click={() => { changeEventState(Action.Rerun, "", 15); isDropdownOpen = false }} >{$t("last_15_min")}</DropdownItem>
		<DropdownItem on:click={() => { changeEventState(Action.Rerun, "", 30); isDropdownOpen = false }} >{$t("last_30_min")}</DropdownItem>
		<DropdownItem on:click={() => { changeEventState(Action.Rerun, "", 60); isDropdownOpen = false }} >{$t("last_hour")}</DropdownItem>
	</Dropdown>
</div>