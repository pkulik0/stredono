<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import {t} from 'svelte-i18n';

	export let open: boolean;
	export let onConfirm: () => void;
</script>

<Modal bind:open size="xs" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{$t("are_you_sure")}</h3>
		<Button on:click={onConfirm} color="red" class="me-2">{$t("yes_im_sure")}</Button>
		<Button color="alternative">{$t("no_cancel")}</Button>
	</div>
</Modal>