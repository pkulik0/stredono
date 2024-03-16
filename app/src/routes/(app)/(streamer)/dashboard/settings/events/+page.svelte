<script lang="ts">
	import { saveSettings, settingsStore } from '$lib/settings';
	import { userStore } from '$lib/user';
	import { Accordion, Button, Checkbox, Heading, Popover } from 'flowbite-svelte';
	import { InfoCircleOutline } from 'flowbite-svelte-icons';
	import { t } from 'svelte-i18n';
	import ChatTTSSettings from './ChatTTSSettings.svelte';
	import CheerSettings from './CheerSettings.svelte';
	import FollowSettings from './FollowSettings.svelte';
	import RaidSettings from './RaidSettings.svelte';
	import SubGiftSettings from './SubGiftSettings.svelte';
	import SubSettings from './SubSettings.svelte';
	import TipSettings from './TipSettings.svelte';

	const save = async () => {
		const user = $userStore;
		if(!user) return;
		await saveSettings(user.Uid)
	}

</script>

{#if $settingsStore}
	<Heading tag="h2">{$t("events")}</Heading>
	<div class="space-y-4 w-full p-4">
		<div class="space-y-4 px-4">
			<div class="py-4">
				<Checkbox bind:checked={$settingsStore.Events.RequireApproval}>
					{$t("require_approval")}
					<InfoCircleOutline class="w-4 h-4 ms-1" />
					<Popover>
						<div class="p-4 w-80">
							{$t("require_approval_help")}
						</div>
					</Popover>
				</Checkbox>
			</div>

			<Accordion class="w-full">
				<TipSettings bind:template={$settingsStore.Events.Tip.Template} bind:minimum={$settingsStore.Events.Tip.MinAmount} />
				<CheerSettings bind:template={$settingsStore.Events.Cheer.Template} bind:minimum={$settingsStore.Events.Cheer.MinAmount} />
				<SubSettings bind:template={$settingsStore.Events.Sub.Template} bind:minimum={$settingsStore.Events.Sub.MinMonths} />
				<SubGiftSettings bind:template={$settingsStore.Events.SubGift.Template} bind:minimum={$settingsStore.Events.SubGift.MinCount} />
				<RaidSettings bind:template={$settingsStore.Events.Raid.Template} bind:minimum={$settingsStore.Events.Raid.MinViewers} />
				<FollowSettings bind:template={$settingsStore.Events.Follow.Template} bind:enabled={$settingsStore.Events.Follow.IsEnabled} />
				<ChatTTSSettings bind:template={$settingsStore.Events.ChatTTS.Template} bind:enabled={$settingsStore.Events.ChatTTS.IsEnabled} />
			</Accordion>
		</div>

		<div class="pt-4">
			<Button on:click={save}>{$t("save")}</Button>
		</div>
	</div>
{/if}