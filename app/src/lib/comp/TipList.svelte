<script lang="ts">
    import type { Tip } from '$lib/pb/tip_pb';
    import {Button, Dropdown, DropdownItem, Listgroup, ListgroupItem} from "flowbite-svelte";
    import {DotsVerticalSolid, RedoOutline, RefreshOutline, ShieldSolid} from "flowbite-svelte-icons";
    import {slide, fade} from "svelte/transition";
    import {t} from 'svelte-i18n';

    export let items: Tip[]

    let menuButtonClass = "flex";
</script>

<Listgroup class="border-0">
    {#each items as item}
        <ListgroupItem class="flex items-center space-x-4 rtl:space-x-reverse">
            <div class="flex-[0.3]">
                <p class="text-sm font-bold text-gray-900 truncate dark:text-white">
                    {item.DisplayName}
                </p>
                <p class="text-sm text-gray-500 truncate dark:text-gray-400" transition:slide>
                    <span transition:fade>{item.Email}</span>
                </p>
            </div>
            <div class="flex-[0.70]">
                {item.Message}
            </div>
            <div class="inline-flex text-base font-semibold text-gray-900 dark:text-white">
                {item.Amount} {item.Currency}
            </div>
            <Button size="sm" color="primary" outline class="text-xs">
                <DotsVerticalSolid class="dots-menu" />
            </Button>
            <Dropdown>
                <DropdownItem class="{menuButtonClass}">
                    <RefreshOutline/>
                    {$t("rerun")}
                </DropdownItem>
                <DropdownItem class="{menuButtonClass}">
                    <ShieldSolid/>
                    {$t("block")}
                </DropdownItem>
                <DropdownItem slot="footer" class="{menuButtonClass} text-red-500">
                    <RedoOutline/>
                    {$t("refund")}
                </DropdownItem>
            </Dropdown>
        </ListgroupItem>
    {/each}
    {#if items.length === 0}
        <ListgroupItem class="flex items-center justify-center h-16">
            <p class="text-gray-500 dark:text-gray-400">{$t("no_tips_recently")}</p>
        </ListgroupItem>
    {/if}
</Listgroup>