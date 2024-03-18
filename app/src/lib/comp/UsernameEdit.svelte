<script lang="ts">
	import { db } from '$lib/ext/firebase/firebase';
	import { getUidByUsernameOther, userStore } from '$lib/user';
	import { collection, getDocs, limit, query, where } from 'firebase/firestore';
	import { Alert, Input, Label, Li, List, P } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { t } from 'svelte-i18n';

	export let username: string;

	let isUnique: boolean = username.length > 0;
	$: isLengthValid = username.length > 3 && username.length <= 20;

	$: isCharactersValid = /^[a-zA-Z0-9_]+$/.test(username);

	export let isValid: boolean;
	$: isValid = isUnique && isLengthValid && isCharactersValid;

	$: colorSuccess = username.trim().length > 0 ? "text-green-300" : "text-gray-400";
	let colorError: string = "text-red-400";

	$: uniqueClass = isUnique ? colorSuccess : colorError;
	$: lengthClass = isLengthValid ? colorSuccess : colorError;
	$: charactersClass = isCharactersValid ? colorSuccess : colorError;

	let timeout: any;
	const onUsernameChange = (e: Event) => {
		if(timeout) clearTimeout(timeout);

		timeout = setTimeout(async () => {
			isUnique = (await getUidByUsernameOther(username)) === "";
		}, 500);
	}
</script>

<div class="space-y-4 flex flex-col">
	<Label>
		{$t("username")}
		<Input type="text" bind:value={username} on:input={onUsernameChange} />
	</Label>

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