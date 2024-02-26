<script lang="ts">
    import { storage, uploadToStorage } from '$lib/ext/firebase/storage';
    import { Button, Heading, Input, Label, Modal, P } from 'flowbite-svelte';
    import FileDropzone from "$lib/comp/FileDropzone.svelte";
    import {ref, listAll, getDownloadURL} from 'firebase/storage';
    import { onMount } from 'svelte';
    import { v4 as uuidv4 } from 'uuid';

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


<Modal bind:open title="Sounds" autoclose outsideclose class="z-100" {backdropClass}>
    <svelte:fragment slot="header">
<!--        <Label class="w-full me-5">-->
<!--            Search-->
<!--            <Input bind:value={searchTerm} placeholder="Type in anything!" />-->
<!--        </Label>-->
    </svelte:fragment>

    <svelte:fragment slot="footer">
        {#if upload}
            <FileDropzone bind:file description="Audio (max 10MB)" />
        {/if}
    </svelte:fragment>

    <P>All these sounds are copyright-free and available for commercial use.</P>

    {#each items as item}
        <div class="flex items-center justify-between w-full px-4 space-y-22">
            <P class="font-bold flex-[0.6]">{item.name}</P>

            <audio controls>
                <source src={item.url} type="audio/mpeg" />
                Your browser does not support the audio element.
            </audio>

            <Button on:click={() => pickSound(item.url)}>Pick</Button>
        </div>
    {/each}
</Modal>