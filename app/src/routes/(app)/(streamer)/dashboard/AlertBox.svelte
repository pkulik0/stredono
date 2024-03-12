<script lang="ts">
	import { Alert } from 'flowbite-svelte';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
</script>

<div class="max-w-7xl m-auto w-full">
	<Alert color="red" class="mt-4">
		<InfoCircleSolid slot="icon" class="w-4 h-4" />
		<span class="font-extrabold">Default alert!</span>
		System messages are displayed here.
	</Alert>
</div>