<script lang="ts">
	import { terraformOutput } from '$lib/constants';
	import { auth } from '$lib/ext/firebase/firebase';
	import { getUidByUsernameOther, userStore } from '$lib/user';
	import axios from 'axios';
	import { Button, Heading, Helper, Input, Label, Listgroup, ListgroupItem, P } from 'flowbite-svelte';
	import { CloseSolid } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { t } from 'svelte-i18n';
	import { getModeratorsListener, moderatorsStore } from './moderators';

	let username = '';

	$: isValid = false;

	let usernameTimeout: any;
	let modUid: string = "";
	const onUsernameChange = async () => {
		clearTimeout(usernameTimeout);
		usernameTimeout = setTimeout(async () => {
			modUid = await getUidByUsernameOther(username.trim());
			isValid = modUid !== ""; // only if user exists
		}, 200);
	};

	enum Action {
		Add = "add",
		Remove = "remove"
	}

	const sendChange = async (uid: string, action: Action) => {
		const user = auth.currentUser
		if(!user) throw new Error("User not logged in");
		const token = await user.getIdToken();

		console.log("sending change", uid, action);

		const res = await axios.get(terraformOutput.FunctionsUrl + `/UserModeratorChange?uid=${uid}&action=${action}`, {
			headers: {
				Authorization: `Bearer ${token}`
			},
		})
		if(res.status !== 200) {
			console.error(res);
		}

	}

	const clickAdd = async () => {
		if(!isValid) return;
		await sendChange(modUid, Action.Add);
		username = "";
		modUid = "";
	}

	let modsUnsub: any;
	onMount(() => {
		const unsub = userStore.subscribe(u => {
			if(!u) {
				if(modsUnsub) modsUnsub();
				modsUnsub = undefined;
				return
			}

			modsUnsub = getModeratorsListener(u.Uid)
		})

		return () => {
			if(modsUnsub) modsUnsub();
			unsub();
		}
	})
</script>

<Heading tag="h2">{$t("moderators")}</Heading>

<div class="space-y-6 w-full p-4">
	<Label class="px-4">
		{$t("username")}
		<Input placeholder={$t("who_you_want_add")} type="text" bind:value={username} on:input={onUsernameChange}>
			<Button size="sm" slot="right" disabled={!isValid} on:click={clickAdd}>
				{$t("add")}
			</Button>
		</Input>
		{#if !isValid && username.trim()}
			<Helper class="mt-0.5">
				{$t("only_existent_users_can_be_moderators")}
			</Helper>
		{/if}
	</Label>

	<div class="space-y-2">
		<Heading tag="h4">{$t("current_moderators")}</Heading>
		<div class="px-4">

			{#if $moderatorsStore.length === 0}
				<P class="text-center py-4">
					<span class="text-gray-400">{$t("no_moderators")}</span>
				</P>
			{:else}
				<Listgroup>
					{#each $moderatorsStore as moderator}
						<ListgroupItem class="flex justify-between">
							<div>
								<span class="font-bold text-lg">
									{moderator.Username}
								</span>
								<span class="text-sm text-gray-500">
									({moderator.Uid})
								</span>
							</div>

							<Button size="xs" color="red" on:click={() => { console.log("aha"); sendChange(moderator.Uid, Action.Remove)}}>
								<CloseSolid class="w-4 h-4" />
							</Button>
						</ListgroupItem>
					{/each}
				</Listgroup>
			{/if}

		</div>
	</div>
</div>