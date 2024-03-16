<script lang="ts">
	import {
		AccordionItem,
		Badge,
		ButtonGroup,
		Checkbox,
		Helper,
		Input,
		InputAddon,
		Label,
		Select
	} from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { t } from 'svelte-i18n';

	export let template: string;
	export let enabled: boolean;
</script>

<AccordionItem>
	<svelte:fragment slot="header">{$t("chat_tts")}</svelte:fragment>
	<div class="space-y-4">
		<Label>
			{$t("template")}
			<Input type="text" bind:value={template}/>
		</Label>
		<div class="space-y-1">
			<Helper>{$t("available_variables")}:</Helper>
			<div class="flex flex-row space-x-2">
				<Badge>&lbrace;user&rbrace;</Badge>
			</div>
		</div>

		<Checkbox bind:checked={enabled}>{$t("enable_chat_tts")}</Checkbox>
	</div>
</AccordionItem>