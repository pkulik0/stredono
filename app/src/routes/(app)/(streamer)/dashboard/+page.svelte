<script lang="ts">
    import { settingsStore } from '$lib/settings';
    import { set } from 'firebase/database';
    import EventButtons from './EventButtons.svelte';
    import LatestCard from "./LatestCard.svelte";

</script>

<div class="flex flex-col space-y-2">
    {#if $settingsStore}
        <EventButtons bind:muted={$settingsStore.Events.IsMuted} bind:paused={$settingsStore.Events.IsPaused}/>
    {/if}

    <div class="flex justify-center space-x-2">
        <LatestCard/>
    </div>
</div>