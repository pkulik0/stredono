<script lang="ts">
    import {Marquee} from "flowbite-svelte";
    import ColoredImage from "./ColoredImage.svelte";

    export let rowSpeed = 0.25;
    export let images: string[]
    export let perRow = 5;

    $: rows = ((): string[][] => {
        const rows = [];
        for (let i = 0; i < images.length; i += perRow) {
            rows.push(images.slice(i, i + perRow));
        }
        return rows;
    })();

    const rowMultipliers = Array.from({length: 20}, (_, i) => Math.random() / 2.5 + 0.8);

</script>

{#each rowMultipliers as multiplier, index}
    <Marquee shadow speed={rowSpeed * multiplier} hoverSpeed={rowSpeed * multiplier} class="opacity-40 dark:opacity-35">
        {#each rows[index % rows.length] as image}
            <div class="p-3">
                <ColoredImage color="gray" src={image}/>
            </div>
        {/each}
    </Marquee>
{/each}