<script lang="ts">
    import { eventsStore } from '$lib/events';
    import { settingsStore } from '$lib/settings';
    import { Heading } from 'flowbite-svelte';
    import { t } from 'svelte-i18n';
    import EventButtons from './EventButtons.svelte';
    import EventList from './EventList.svelte';

</script>

<div class="flex flex-col space-y-2">
    {#if $settingsStore}
        <EventButtons bind:muted={$settingsStore.Events.IsMuted} bind:paused={$settingsStore.Events.IsPaused}/>
    {/if}

    <div class="flex flex-col py-6 justify-center space-x-2">
        <Heading tag="h5">{$t("latest_events")}</Heading>

        <EventList events={$eventsStore}/>
    </div>
</div>