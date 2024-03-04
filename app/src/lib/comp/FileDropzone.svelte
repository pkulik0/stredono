<script lang="ts">
    import {UploadOutline} from "flowbite-svelte-icons";
    import { Dropzone, Heading } from 'flowbite-svelte';
    import { t } from 'svelte-i18n';

    export let file: File | undefined;
    export let description: string = "";

    const dropHandle = (event: DragEvent) => {
        file = undefined;
        event.preventDefault();

        if (!event.dataTransfer) return;
        const dataTransfer = event.dataTransfer as DataTransfer;

        const item = dataTransfer.items[0];
        if (!item) return;

        const gif = item.getAsFile();
        if (!gif) return;

        file = gif;
    };

    const handleChange = (event: Event) => {
        if (!event.target) return;
        const target = event.target as HTMLInputElement;

        const files = target.files;
        if (!files) return;

        const gif = files.item(0)
        if (!gif) return;
        file = gif;
    };
</script>

<Dropzone class="max-h-40" id="dropzone" on:drop={dropHandle} on:dragover={e => e.preventDefault()} on:change={handleChange}>
    <UploadOutline class="mb-3 w-10 h-10 text-gray-400" />
    <p class="mb-2 text-sm text-gray-500 dark:text-gray-400">{$t("dropzone_msg")}</p>
    <p class="text-xs text-gray-500 dark:text-gray-400">{description}</p>
</Dropzone>