<script lang="ts">
	import { AccordionItem, Checkbox } from 'flowbite-svelte';
	import {t} from 'svelte-i18n';

	export let open = false;
	export let header = 'Accordion header';
</script>

<AccordionItem bind:open>
	<span slot="header">{header}</span>
	<div class="space-y-4">
		<Checkbox>{$t("read_out_loud")}</Checkbox>
		<slot/>
	</div>
</AccordionItem>