<script lang="ts">
    import {Button, Card, Dropdown, DropdownItem, Listgroup, ListgroupItem, TextPlaceholder} from "flowbite-svelte";
    import {DotsVerticalSolid, RedoOutline, RefreshOutline, ShieldSolid} from "flowbite-svelte-icons";
    import type {SendDonateRequest} from "$lib/pb/functions_pb";
    import {onMount} from "svelte";
    import {streamModeStore, donationsStore} from "$lib/stores";
    import {slide, fade} from "svelte/transition";

    let donations: SendDonateRequest[]|undefined = undefined;

    export let size = 5;

    onMount(() => {
        return donationsStore.subscribe(value => {
            let startIndex = value.length - size;
            if (startIndex < 0) startIndex = 0;
            donations = value.slice(startIndex).reverse();
        });
    })

    let menuButtonClass = "flex";
</script>

<Card padding="xl" size="xl" class="mt-4 flex-1">
    <div class="flex justify-between items-center mb-4">
        <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">Latest Donations</h5>
        <a href="/panel/donations" class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"> View all </a>
    </div>

    {#if donations !== undefined}
        <Listgroup class="border-0">
            {#each donations as item}
                <ListgroupItem class="flex items-center space-x-4 rtl:space-x-reverse">
                    <div class="flex-[0.3]">
                        <p class="text-sm font-bold text-gray-900 truncate dark:text-white">
                            {item.sender}
                        </p>
                        {#if $streamModeStore === false}
                            <p class="text-sm text-gray-500 truncate dark:text-gray-400" transition:slide>
                                <span transition:fade>{item.email}</span>
                            </p>
                        {/if}
                    </div>
                    <div class="flex-[0.70]">
                        {item.message}
                    </div>
                    <div class="inline-flex text-base font-semibold text-gray-900 dark:text-white">
                        {item.amount} {item.currency}
                    </div>
                    <Button size="sm" color="primary" outline class="text-xs">
                        <DotsVerticalSolid class="dots-menu" />
                    </Button>
                    <Dropdown>
                        <DropdownItem class="{menuButtonClass}">
                            <RefreshOutline/>
                            Rerun
                        </DropdownItem>
                        <DropdownItem class="{menuButtonClass}">
                            <ShieldSolid/>
                            Block
                        </DropdownItem>
                        <DropdownItem slot="footer" class="{menuButtonClass} text-red-500">
                            <RedoOutline/>
                            Refund
                        </DropdownItem>
                    </Dropdown>
                </ListgroupItem>
            {/each}
        </Listgroup>
    {:else}
        <TextPlaceholder/>
    {/if}
</Card>