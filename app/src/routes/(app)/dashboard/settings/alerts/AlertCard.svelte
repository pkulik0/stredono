<script lang="ts">
	import { type Alert } from '$lib/pb/alert_pb';
	import { Button, Card, Heading, Img } from 'flowbite-svelte';
	import { PenSolid, TrashBinSolid } from 'flowbite-svelte-icons';

	export let alert: Alert;
	export let open: boolean = false;
</script>

<Card size="sm" padding="xl" class="space-y-2 justify-center w-full max-w-full">
	<div class="flex space-x-4">
		<div class="w-60">
			<Img src={alert.GifUrl} alt={alert.Id + "'s gif"} class="rounded-lg w-full" />
		</div>

		<div class="flex flex-1 flex-col space-y-2 text-center justify-center">
			<Heading tag="h6">{alert.Min} zł - {alert.Max} zł</Heading>
			<!--TODO: currency -->
		</div>

		<div class="flex flex-row space-x-2">
			<Button outline class="m-auto" on:click={() => open = true}>
				<PenSolid />
			</Button>
			<Button outline color="red" class="m-auto">
				<TrashBinSolid />
			</Button>
		</div>
	</div>
</Card>