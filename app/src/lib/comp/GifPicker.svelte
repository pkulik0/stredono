<script lang="ts">
    import {getTopGifs, type Gif, searchGifs} from "$lib/ext/tenor";
    import {onMount} from "svelte";
    import {Dropzone, Input, Label, Modal} from "flowbite-svelte";
    import GifGallery from "$lib/comp/GifGallery.svelte";
    import {writable, type Writable} from "svelte/store";
    import {UploadOutline} from "flowbite-svelte-icons";
    import axios from "axios";

    const limit = 30;

    export let open: boolean = false;

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
        <Label class="w-full me-10">
            Search
            <Input bind:value={searchTerm} on:input={onSearchInputChange} placeholder="Type in anything!" />
        </Label>
    </svelte:fragment>

    <svelte:fragment slot="footer">
        <Dropzone class="max-h-40" id="dropzone" on:drop={dropHandle} on:dragover={e => e.preventDefault()} on:change={handleChange}>
            <UploadOutline class="mb-3 w-10 h-10 text-gray-400" />
            <p class="mb-2 text-sm text-gray-500 dark:text-gray-400"><span class="font-semibold">Click to upload</span> or drag and drop</p>
            <p class="text-xs text-gray-500 dark:text-gray-400">.gif (max. 5MB)</p>
        </Dropzone>
    </svelte:fragment>

    <GifGallery onGifPress={onGifPress} gifs={gifs} loadMore={loadMore} columns={3} />
</Modal>