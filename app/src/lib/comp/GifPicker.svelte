<script lang="ts">
    import {getTopGifs, type Gif, searchGifs} from "$lib/ext/tenor";
    import {onMount} from "svelte";
    import {Dropzone, Input, Label, Modal} from "flowbite-svelte";
    import GifGallery from "$lib/comp/GifGallery.svelte";
    import {writable, type Writable} from "svelte/store";
    import {UploadOutline} from "flowbite-svelte-icons";
    import axios from "axios";
    import FileDropzone from "$lib/comp/FileDropzone.svelte";

    const limit = 30;

    export let open: boolean = false;
    export let upload: boolean = true;

    export let file: File|undefined = undefined
    $: if (file) open = false;

    export let searchTerm = "";

    $: termTrimmed = searchTerm.trim();
    let lastTerm: string|undefined;

    let trendingGifs: Gif[] = [];
    let foundGifsMap: Writable<Map<string, Gif[]>> = writable(new Map());
    $: foundGifs = $foundGifsMap.get(termTrimmed);
    $: lastGifs = lastTerm ? $foundGifsMap.get(lastTerm) : undefined;

    let timeout: any;
    let lastChangeTime: number = 0;
    const inputDelay = 500;

    const onSearchInputChange = (e: Event) => {
        if (timeout) {
            clearTimeout(timeout);
        }

        lastChangeTime = Date.now();

        timeout = setTimeout(async () => {
            if (!termTrimmed) return;

            const diff = Date.now() - lastChangeTime;
            if (diff > inputDelay) {
                await getGifsByTerm(termTrimmed);
            }

            timeout = null;
        }, inputDelay);
    }

    let trendingNext: string|undefined;
    const getTrendingGifs = async () => {
        const gifsResponse = await getTopGifs(limit, trendingNext);
        trendingNext = gifsResponse.next;
        trendingGifs = trendingGifs.concat(gifsResponse.gifs);
    }

    let searchNext: string|undefined;
    const getGifsByTerm = async (term: string) => {
        let gifs = $foundGifsMap.get(term);
        if (!gifs) gifs = [];

        const gifsResposne = await searchGifs(term, limit, searchNext);
        searchNext = gifsResposne.next;
        gifs = gifs.concat(gifsResposne.gifs);
        foundGifsMap.update(map => {
            map.set(term, gifs || []);
            return map;
        })
        lastTerm = term;
    }
    const onGifPress = async (gif: Gif) => {
        const res = await axios.get(gif.url, { responseType: "blob" });
        const blob = new Blob([res.data], { type: "image/gif" });
        file = new File([blob], gif.title, {type: "image/gif"});
    }

    onMount(async () => {
        if(searchTerm) await getGifsByTerm(searchTerm);
        else await getTrendingGifs();
    });

    $: gifs = function () {
        if (foundGifs) return foundGifs;
        if (lastGifs && termTrimmed) return lastGifs;
        return trendingGifs;
    }()

    $: loadMore = function () {
        if (foundGifs) return () => getGifsByTerm(termTrimmed);
        if (lastGifs && termTrimmed) return () => getGifsByTerm(lastTerm || "ALWAYS NOT NULL, DONT MIND");
        return getTrendingGifs;
    }()

    let backdropClass = "fixed inset-0 z-50 bg-gray-900 bg-opacity-50 dark:bg-opacity-80";
</script>


<Modal bind:open title="Gifs" autoclose outsideclose class="z-100" {backdropClass}>

    <svelte:fragment slot="header">
        <Label class="w-full me-5">
            Search
            <Input bind:value={searchTerm} on:input={onSearchInputChange} placeholder="Type in anything!" />
        </Label>
    </svelte:fragment>

    <svelte:fragment slot="footer">
        {#if upload}
            <FileDropzone bind:file description=".gif (max. 5MB)"/>
        {/if}
    </svelte:fragment>

    <GifGallery onGifPress={onGifPress} gifs={gifs} loadMore={loadMore} columns={3} />
</Modal>