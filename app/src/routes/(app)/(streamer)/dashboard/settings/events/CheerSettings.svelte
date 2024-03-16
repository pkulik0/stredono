<script lang="ts">
	import { AccordionItem, Badge, ButtonGroup, Helper, Input, InputAddon, Label } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { t } from 'svelte-i18n';

	export let template: string;
	export let minimum: number;

	let minimumStr: string;
	$: if(minimumStr) minimum = parseInt(minimumStr);

	onMount(() => {
		minimumStr = minimum.toString()
	});
</script>

<AccordionItem>
	<svelte:fragment slot="header">{$t("cheers")}</svelte:fragment>
	<div class="space-y-4">
		<Label>
			{$t("template")}
			<Input type="text" bind:value={template}/>
		</Label>
		<div class="space-y-1">
			<Helper>{$t("available_variables")}:</Helper>
			<div class="flex flex-row space-x-2">
				<Badge>&lbrace;user&rbrace;</Badge>
				<Badge color="red">&lbrace;value&rbrace;</Badge>
			</div>
		</div>

		<Label>
			{$t("min_amount")}
			<ButtonGroup class="w-full" size="sm">
				<InputAddon>{$t("bits_amount")}</InputAddon>
				<Input min="0" type="number" bind:value={minimumStr}/>
			</ButtonGroup>
		</Label>
	</div>
</AccordionItem>