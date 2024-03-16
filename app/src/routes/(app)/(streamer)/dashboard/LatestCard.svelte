<script lang="ts">
    import { eventsStore, getEventsDashboardListener } from '$lib/events';
    import type { Tip } from '$lib/pb/tip_pb';
    import { userStore } from '$lib/user';
    import { Card, TextPlaceholder} from "flowbite-svelte";
    import {onMount} from "svelte";
    import {tipsStore} from "$lib/tips";
    import TipList from "$lib/comp/TipList.svelte";
    import { t } from 'svelte-i18n';
    import EventList from './EventList.svelte';

    let eventsUnsub: (() => void) | undefined;
    onMount(() => {
        return userStore.subscribe(u => {
            if(!u) {
                if(eventsUnsub) eventsUnsub()
                return
            }
            eventsUnsub = getEventsDashboardListener(u.Uid);
        })
    })

</script>

<Card padding="xl" size="xl" class="mt-4 flex-1">
    <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">{$t("latest_events")}</h5>

    <EventList events={$eventsStore}/>
</Card>