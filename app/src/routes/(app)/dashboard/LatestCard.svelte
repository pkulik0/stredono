<script lang="ts">
    import type { Tip } from '$lib/pb/stredono_pb';
    import { Card, TextPlaceholder} from "flowbite-svelte";
    import {onMount} from "svelte";
    import {tipsStore} from "$lib/tips";
    import DonationList from "$lib/comp/DonationList.svelte";

    let items: Tip[]|undefined = undefined;
    $: itemsSlice = items?.slice(Math.min(items?.length - size, 0)).reverse();

    export let size = 4;

    onMount(() => {
        return tipsStore.subscribe(donations => {
            if(!donations.hasOwnProperty(0)) return;
            items = donations[0].tips;
        });
    })

</script>

<Card padding="xl" size="xl" class="mt-4 flex-1">
    <div class="flex justify-between mb-4">
        <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">Latest Donations</h5>
        <a href="/panel/donations" class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"> View all </a>
    </div>

    {#if itemsSlice}
        <DonationList items={itemsSlice}/>
    {:else}
        <TextPlaceholder/>
    {/if}
</Card>