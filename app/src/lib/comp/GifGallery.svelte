<script lang="ts">
    import type {Gif} from "$lib/ext/tenor";
    import {Gallery, Img} from "flowbite-svelte";
    import {onMount} from "svelte";

    export let gifs: Gif[] = [];
    export let columns = 4;
    export let loadMore: () => Promise<void> = async () => {};

    let columnsArray: Gif[][] = [];

    $: {
        const array: Gif[][] = Array.from({length: columns}, () => []);
        gifs.forEach((gif, index) => {
            array[index % columns].push(gif);
        });
        columnsArray = array;
    }

    let observer: IntersectionObserver;

    onMount(() => {
        observer = new IntersectionObserver((entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) {
                    loadMore();
                }
            });
        });
        const sentinel = document.getElementById("sentinel");
        if (sentinel) observer.observe(sentinel);
        return () => {
            if(observer) observer.disconnect();
        };
    })

    export let onGifPress: (gif: Gif) => void = () => {};
</script>


<div class="grid grid-cols-{columns} gap-4">
    {#each columnsArray as column}
        <div class="flex flex-col gap-4">
            {#each column as gif}
                <button on:click={() => onGifPress(gif)}>
                    <Img src={gif.preview} alt={gif.title} class="rounded-xl hover:opacity-85 hover:border-primary-600 border-4 border-transparent w-full"/>
                </button>
            {/each}
        </div>
    {/each}
</div>

<div id="sentinel" class="h-0.5"/>
