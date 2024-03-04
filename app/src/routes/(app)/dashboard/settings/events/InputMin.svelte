<script lang="ts">
	import { Checkbox, Input, type InputType, Label } from 'flowbite-svelte';
	import {slide} from 'svelte/transition';

	export let checked = false;
	export let type: InputType = 'text';
	export let placeholder = '5.00';

	export let label = 'Minimum tip';
	export let value = '';
	export let ttsValue = '';
</script>

<Label>
	{label}
	<Input {type} {placeholder} bind:value/>
</Label>

<div class="space-y-4">
	<Checkbox bind:checked>
		Use a different minimum for reading out loud
	</Checkbox>

	{#if checked}
		<div transition:slide>
			<Input {type} {placeholder} bind:value={ttsValue}/>
		</div>
	{/if}
</div>