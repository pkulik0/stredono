<script lang="ts">
    import type { Tip } from '$lib/pb/tip_pb';
    import { Card, TextPlaceholder} from "flowbite-svelte";
    import {onMount} from "svelte";
    import {tipsStore} from "$lib/tips";
    import TipList from "$lib/comp/TipList.svelte";
    import { t } from 'svelte-i18n';

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
        <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">{$t("latest_tips")}</h5>
        <a href="/dashboard/tips" class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500">{$t("view_all")}</a>
    </div>

    {#if itemsSlice}
        <TipList items={itemsSlice}/>
    {:else}
        <TextPlaceholder/>
    {/if}
</Card>