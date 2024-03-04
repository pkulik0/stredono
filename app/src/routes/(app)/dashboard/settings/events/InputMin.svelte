<script lang="ts">
	import { Checkbox, Input, type InputType, Label } from 'flowbite-svelte';
	import { t } from 'svelte-i18n';
	import {slide} from 'svelte/transition';

	export let checked = false;
	export let type: InputType = 'text';
	export let placeholder = '5.00';

	export let label = '';
	export let value = '';
	export let ttsValue = '';
</script>

<Label>
	{label}
	<Input {type} {placeholder} bind:value/>
</Label>

<div class="space-y-4">
	<Checkbox bind:checked>
		{$t("diff_min_tts")}
	</Checkbox>

	{#if checked}
		<div transition:slide>
			<Input {type} {placeholder} bind:value={ttsValue}/>
		</div>
	{/if}
</div>