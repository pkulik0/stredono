<script lang="ts">
    import { Event } from '$lib/pb/event_pb';
    import {Listgroup, ListgroupItem} from "flowbite-svelte";
    import {t} from 'svelte-i18n';
    import EventItem from './EventItem.svelte';
    import {slide} from 'svelte/transition';

    export let events: Event[]
</script>

<Listgroup class="mt-4 p-2">
    {#each events as event}
        <div transition:slide>
            <EventItem {event} />
        </div>
    {/each}
    {#if events.length === 0}
        <ListgroupItem class="flex items-center justify-center h-16">
            <p class="text-gray-500 dark:text-gray-400">{$t("no_events_recently")}</p>
        </ListgroupItem>
    {/if}
</Listgroup>