<script lang="ts">
	import { removeAlert } from '$lib/alerts';
	import ConfirmationModal from '$lib/comp/ConfirmationModal.svelte';
	import { type Alert } from '$lib/pb/alert_pb';
	import { Button, Card, Heading, Img, Modal } from 'flowbite-svelte';
	import { PenSolid, TrashBinSolid } from 'flowbite-svelte-icons';
	import { alertStore, drawerHiddenStore } from './stores';

	export let alert: Alert;
	export let unitLabel: string = "";

	const onRemove = (confirm: boolean) => {
		if(confirm) {
			removeAlert(alert);
		} else {
			showConfirmation = true;
		}
	};

	const onEdit = () => {
		alertStore.set(alert);
		drawerHiddenStore.set(false);
	}

	let showConfirmation = false;
</script>

<Card size="sm" padding="xl" class="space-y-2 justify-center w-full max-w-full">
	<div class="flex space-x-4">
		{#if alert.GifUrl}
			<div class="w-60">
				<Img src={alert.GifUrl} alt="" class="rounded-lg m-auto max-h-40 w-auto" />
			</div>
		{/if}

		<div class="flex flex-1 flex-col space-y-2 text-center justify-center">
			<Heading tag="h6">{alert.Min} {unitLabel} - {alert.Max ?? '∞' } {unitLabel}</Heading>
			<!--TODO: currency -->
		</div>

		<div class="flex flex-row space-x-2">
			<Button outline class="m-auto" on:click={onEdit}>
				<PenSolid />
			</Button>
			<Button outline color="red" class="m-auto" on:click={() => onRemove(false)}>
				<TrashBinSolid />
			</Button>
		</div>
	</div>
</Card>

<ConfirmationModal bind:open={showConfirmation} onConfirm={() => onRemove(true)} />