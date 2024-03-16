<script lang="ts">
	import { goto } from '$app/navigation';
	import { Alert, Button, Heading, Input, Label, Li, List, P } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
	import {t} from 'svelte-i18n';
	import { checkIfUnique, saveUsername } from './username';

	let username: string = "";

	let timeout: any;

	let isUnique: boolean = false;
	$: isLengthValid = username.length > 3 && username.length <= 20;

	$: isCharactersValid = /^[a-zA-Z0-9_]+$/.test(username);
	$: isValid = isUnique && isLengthValid && isCharactersValid;

	const onUsernameChange = (e: Event) => {
		if(timeout) clearTimeout(timeout);

		timeout = setTimeout(async () => {
			isUnique = await checkIfUnique(username);
		}, 500);
	}

	const save = async () => {
		await saveUsername(username)
		await goto("/dashboard");
	}

	$: colorSuccess = username.trim().length > 0 ? "text-green-400" : "text-gray-400";
	let colorError: string = "text-red-400";

	$: uniqueClass = isUnique ? colorSuccess : colorError;
	$: lengthClass = isLengthValid ? colorSuccess : colorError;
	$: charactersClass = isCharactersValid ? colorSuccess : colorError;
</script>

<div class="flex flex-col max-w-xl m-auto justify-center h-full space-y-8">
	<div class="space-y-2">
		<Heading tag="h3">{$t("choose_a_username")}</Heading>
		<Heading tag="h5"><span class="text-gray-500">https://stredono.com/user/</span><span class="text-primary-600 dark:text-primary-500">{username}</span></Heading>
	</div>

	<div class="space-y-4">
		<Label>
			{$t("username")}
			<Input type="text" bind:value={username} on:input={onUsernameChange} />
		</Label>
		<Button color={isValid ? "primary" :  "alternative"} class="w-full" disabled={!isValid} on:click={save}>
			{$t("continue")}
		</Button>
	</div>

	<Alert class="!items-start">
		<span slot="icon">
			<InfoCircleSolid slot="icon" class="w-6 h-6" />
			<span class="sr-only">Info</span>
		</span>
		<P>{$t("requirements")}</P>
		<List class="mt-1.5 ms-4">
			<Li><span class={uniqueClass}>{$t("requirements_unique")}</span></Li>
			<Li><span class={lengthClass}>{$t("requirements_length")}</span></Li>
			<Li><span class={charactersClass}>{$t("requirements_characters")}</span></Li>
		</List>
	</Alert>
</div>