<script lang="ts">
    import {Input, Label, Modal} from "flowbite-svelte";
    import FileDropzone from "$lib/comp/FileDropzone.svelte";

    export let open: boolean = false;
    export let upload: boolean = true;

    export let file: File|undefined = undefined
    $: if (file) open = false;

    export let searchTerm = "";
    $: termTrimmed = searchTerm.trim();

    let backdropClass = "fixed inset-0 z-50 bg-gray-900 bg-opacity-50 dark:bg-opacity-80";
</script>


<Modal bind:open title="Gifs" autoclose outsideclose class="z-100" {backdropClass}>
    <svelte:fragment slot="header">
        <Label class="w-full me-5">
            Search
            <Input bind:value={searchTerm} placeholder="Type in anything!" />
        </Label>
    </svelte:fragment>

    <svelte:fragment slot="footer">
        {#if upload}
            <FileDropzone bind:file description="Audio (max 10MB)" />
        {/if}
    </svelte:fragment>

    Euaouaouts
</Modal>