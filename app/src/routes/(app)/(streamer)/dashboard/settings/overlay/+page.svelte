<script lang="ts">
	import ConfirmationModal from '$lib/comp/ConfirmationModal.svelte';
	import { terraformOutput } from '$lib/constants';
	import { auth } from '$lib/ext/firebase/firebase';
	import { sendNotification, Notification } from '$lib/notifications';
	import {settingsStore} from '$lib/settings';
	import { userStore } from '$lib/user';
	import axios from 'axios';
	import { Alert, Button, ButtonGroup, Heading, Input, InputAddon, Label, P } from 'flowbite-svelte';
	import {
		ClipboardListOutline, ClipboardOutline,
		ExclamationCircleOutline,
		InfoCircleSolid,
		RefreshOutline
	} from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import {t} from 'svelte-i18n';
	import { getKeyListener, overlayKeyStore } from './key';

	const overlayUrlBase = "https://stredono.com/overlay?key=";
	$: overlayUrl = overlayUrlBase + $overlayKeyStore

	const onCopyClick = () => {
		navigator.clipboard.writeText(overlayUrl);
		sendNotification(new Notification($t("url_copied")));
	};

	let showConfirmation = false;
	const onRegenerateClick = async (confirm: boolean) => {
		if(!confirm) {
			showConfirmation = true;
			return;
		}

		const user = auth.currentUser
		if(!user) return
		const token = await user.getIdToken();

		const res = await axios.get(terraformOutput.FunctionsUrl + "/UserRegenerateKey", {
			"headers": {
				"Authorization": "Bearer " + token
			}
		})
		if(res.status === 200) {
			sendNotification(new Notification($t("successfully_regenerated_key")));
			return
		}
	}

	let keyUnsub: any;
	onMount(() => {
		const unsub = userStore.subscribe(u => {
			if(!u) {
				if(keyUnsub) keyUnsub();
				keyUnsub = undefined;
				return;
			}

			keyUnsub = getKeyListener(u.Uid);
		})

		return () => {
			unsub();
			if(keyUnsub) keyUnsub();
		}
	})
</script>

{#if $settingsStore}
	<Heading tag="h2">{$t("overlay")}</Heading>

	<div class="w-full space-y-4 mt-6 px-4">
		<Label>
			{$t("url")}
			<Input value={overlayUrl} type="password">
				<div slot="right" class="space-x-1">
					<Button size="sm" color="red" outline on:click={() => onRegenerateClick(false)}>
						<RefreshOutline class="w-4 h-4 me-1" />
						{$t("regenerate")}
					</Button>

					<Button size="sm" on:click={onCopyClick}>
						<ClipboardOutline class="w-4 h-4 me-1" />
						{$t("copy_verb")}
					</Button>
				</div>
			</Input>
		</Label>
		<Alert color="red">
			<span slot="icon">
				<ExclamationCircleOutline slot="icon" class="w-6 h-6" />
				<span class="sr-only">{$t("warning")}</span>
			</span>
			<P class="text-sm font-semibold">
				<span class="text-red-300">{$t("overlay_url_warning")}</span>
			</P>
		</Alert>
	</div>
{/if}

<ConfirmationModal bind:open={showConfirmation} onConfirm={() => onRegenerateClick(true)} />