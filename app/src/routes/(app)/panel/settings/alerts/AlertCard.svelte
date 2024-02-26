<script lang="ts">
	import { type Alert, AnimationType, Alignment, Position, Speed } from '$lib/pb/stredono_pb';
	import { Button, Card, Checkbox, Heading, Hr, Img, Label, P } from 'flowbite-svelte';
	import {
		ArrowDownOutline,
		ChevronDownSolid, ChevronUpSolid,
		EditSolid,
		EyeSolid,
		PauseSolid,
		TrashBinSolid
	} from 'flowbite-svelte-icons';
	import {slide} from 'svelte/transition';

	export let alert: Alert;
	export let open: boolean = false;
</script>


<Card size="sm" padding="xl" class="space-y-2 justify-center w-full max-w-full">
	<div class="flex space-x-4">
		<div class="w-60">
			<Img src={alert.GifUrl} alt={alert.Id + "'s gif"} class="rounded-lg w-full" />
		</div>

		<div class="flex flex-1 flex-col space-y-2 text-center justify-center">
			<Heading tag="h5">{alert.Message}</Heading>
			<Heading tag="h6">{alert.Min} zł - {alert.Max} zł</Heading>
		</div>



		<Button color="alternative" class="m-auto" on:click={() => open = !open}>
			{#if open}
				<ChevronUpSolid />
			{:else}
				<ChevronDownSolid />
			{/if}
		</Button>
	</div>

	{#if open}
		<Hr/>
		<div class="space-y-2" transition:slide>
			<audio controls class="w-full">
				<source src={alert.SoundUrl} type="audio/mpeg" />
				Your browser does not support the audio element.
			</audio>

			<P><span class="font-extrabold">Text Color: </span><span style="color: {alert.TextColor};">{alert.TextColor}</span></P>
			<P><span class="font-extrabold">Accent Color: </span><span style="color: {alert.AccentColor};">{alert.AccentColor}</span></P>
			<P><span class="font-extrabold">Alignment: </span>{JSON.parse(JSON.stringify(Alignment))[alert.Alignment]}</P>
			<P><span class="font-extrabold">Text Position: </span>{JSON.parse(JSON.stringify(Position))[alert.TextPosition]}</P>
			<P><span class="font-extrabold">Animation Type: </span>{JSON.parse(JSON.stringify(AnimationType))[alert.Animation]}</P>
			<P><span class="font-extrabold">Animation Speed: </span>{JSON.parse(JSON.stringify(Speed))[alert.AnimationSpeed]}</P>
		</div>
	{/if}
</Card>

<!--	<div>-->
<!--		<Button color="red" outline>-->
<!--			<TrashBinSolid/>-->
<!--			Delete-->
<!--		</Button>-->

<!--		<Button outline color="green">-->
<!--			<EyeSolid />-->
<!--			Test-->
<!--		</Button>-->

<!--		<Button outline color="blue">-->
<!--			<EditSolid/>-->
<!--			Edit-->
<!--		</Button>-->
<!--	</div>-->

