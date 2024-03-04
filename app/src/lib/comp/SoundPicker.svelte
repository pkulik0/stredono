<script lang="ts">
    import FileDropzone from '$lib/comp/FileDropzone.svelte';
    import { storage, uploadToStorage } from '$lib/ext/firebase/storage';
    import { Button, Heading, Input, Label, Modal, P } from 'flowbite-svelte';
    import {ref, listAll, getDownloadURL} from 'firebase/storage';
    import { onMount } from 'svelte';
    import { v4 as uuidv4 } from 'uuid';
    import {t} from 'svelte-i18n';

    export let url: string;
    export let open: boolean = false;
    export let upload: boolean = true;

    let file: File|undefined;

    $: if (file) {
        uploadFile(file);
    }

    export const uploadFile = async (file: File) => {
        url = await uploadToStorage("sounds", uuidv4().replace(/-/g, ""), file, false);
        open = false;
    }

    export let searchTerm = "";
    $: termTrimmed = searchTerm.trim();

    interface StorageItem {
        path: string;
        name: string;
        url: string;
    }
    let items: StorageItem[] = [];

    onMount(async () => {
        const res = await listAll(ref(storage, 'public/audio'))
        const m = res.items.map(async (item) => {
            const name = item.name.substring(0, item.name.lastIndexOf('.')).split('-').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');
            return {
                path: item.fullPath,
                name: name,
                url: await getDownloadURL(item)
            }
        })
        items = await Promise.all(m);
    })

    $: if (open) url = "";

    const pickSound = (soundUrl: string) => {
        url = soundUrl
        open = false
    }

    let backdropClass = "fixed inset-0 z-50 bg-gray-900 bg-opacity-50 dark:bg-opacity-80";
</script>


<Modal bind:open title={$t("sounds")} autoclose outsideclose class="z-100" {backdropClass}>
    <svelte:fragment slot="footer">
        {#if upload}
            <FileDropzone bind:file description={$t("size_limit_label", {values: {format: "Audio", "size": "10", "unit": "MB"}})} />
        {/if}
    </svelte:fragment>

    <P>{$t("sounds_commercial_use")}</P>

    {#each items as item}
        <div class="flex items-center justify-between w-full px-4 space-y-22">
            <P class="font-bold flex-[0.6]">{item.name}</P>

            <audio src={item.url} controls/>

            <Button on:click={() => pickSound(item.url)}>{$t("pick")}</Button>
        </div>
    {/each}
</Modal>