<script lang="ts">
	import { goto } from '$app/navigation';
	import { Alert, Button, Heading, Input, Label, Li, List, P } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
	import {t} from 'svelte-i18n';
	import { saveUsername } from './username';
	import UsernameEdit from '$lib/comp/UsernameEdit.svelte';

	let username: string = "";
	let isValid: boolean = false;

	const save = async () => {
		try {
			await saveUsername(username)
			await goto("/dashboard");
		} catch (error) {
			console.error(error);
		}
	}
</script>

<div class="flex flex-col max-w-xl m-auto justify-center h-full space-y-8">
	<div class="space-y-2">
		<Heading tag="h3">{$t("choose_a_username")}</Heading>
		<Heading tag="h5"><span class="text-gray-500">https://stredono.com/user/</span><span class="text-primary-600 dark:text-primary-500">{username}</span></Heading>
	</div>

	<div class="space-y-4">
		<UsernameEdit bind:username bind:isValid />

		<Button color={isValid ? "primary" :  "alternative"} class="w-full" disabled={!isValid} on:click={save}>
			{$t("continue")}
		</Button>
	</div>
</div>