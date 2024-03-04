<script lang="ts">
	import { AccordionItem, Checkbox } from 'flowbite-svelte';

	export let open = false;
	export let header = 'Accordion header';
</script>

<AccordionItem bind:open>
	<span slot="header">{header}</span>
	<div class="space-y-4">
		<Checkbox>Enable TTS</Checkbox>
		<slot/>
	</div>
</AccordionItem>