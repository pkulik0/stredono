<script lang="ts">
	import { type Alert, AnimationType, Alignment, Position, Speed } from '$lib/pb/stredono_pb';
	import { Button, Card, Heading, Hr, Img, P } from 'flowbite-svelte';
	import { ChevronDownSolid, ChevronUpSolid } from 'flowbite-svelte-icons';
	import {slide} from 'svelte/transition';
	import {t} from 'svelte-i18n';

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
			<!--TODO: currency -->
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
			<audio src={alert.SoundUrl} controls class="w-full" />

			<P><span class="font-extrabold">{$t("text_color")}: </span><span style="color: {alert.TextColor};">{alert.TextColor}</span></P>
			<P><span class="font-extrabold">{$t("accent_color")}: </span><span style="color: {alert.AccentColor};">{alert.AccentColor}</span></P>
			<P><span class="font-extrabold">{$t("alignment")}: </span>{JSON.parse(JSON.stringify(Alignment))[alert.Alignment]}</P>
			<P><span class="font-extrabold">{$t("text_position")}: </span>{JSON.parse(JSON.stringify(Position))[alert.TextPosition]}</P>
			<P><span class="font-extrabold">{$t("animation_type")}: </span>{JSON.parse(JSON.stringify(AnimationType))[alert.Animation]}</P>
			<P><span class="font-extrabold">{$t("animation_speed")}: </span>{JSON.parse(JSON.stringify(Speed))[alert.AnimationSpeed]}</P>
		</div>
	{/if}
</Card>