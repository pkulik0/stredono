<script lang="ts">
    import { settingsStore } from '../../../../../../lib/events_settings';
    import { EventType } from '$lib/pb/event_pb';
    import { pbEnumToItems } from '$lib/util';
    import {PlusSolid} from "flowbite-svelte-icons";
    import AlertCard from './AlertCard.svelte';
    import AlertsDrawer from "./AlertsDrawer.svelte";
    import {
        Button, Heading,
        ImagePlaceholder, Label, P, Select
    } from 'flowbite-svelte';
    import { locale, t } from 'svelte-i18n';

    $: alerts = $settingsStore?.Alerts;
    $: filteredAlerts = alerts?.filter((alert) => alert.EventType === selectedType);

    let eventTypesItems = pbEnumToItems(EventType);
    $: if($locale) {
        eventTypesItems = pbEnumToItems(EventType);
    }
    let selectedType: EventType = EventType.TIP;

    let drawerHidden = true;
</script>

<Heading tag="h2">{$t("alerts")}</Heading>
<div class="space-y-4 w-full p-4">
    <Label>
        {$t("type")}
        <Select items={eventTypesItems} bind:value={selectedType} placeholder={$t("type_filter")}/>
    </Label>

    <div class="flex justify-end">
        <Button on:click={() => drawerHidden = false}>
            <PlusSolid class="w-4 h-4" />
            <span class="ms-2">{$t("new")}</span>
        </Button>
    </div>

    <div class="p-4 space-y-4">
        {#if filteredAlerts !== undefined}
            {#if filteredAlerts.length === 0}
                <P class="text-center py-8">
                    {$t("no_alerts")}
                </P>
            {:else}
                {#each filteredAlerts as alert}
                    <AlertCard {alert}/>
                {/each}
            {/if}
        {:else}
            <ImagePlaceholder/>
        {/if}
    </div>
</div>

<AlertsDrawer eventType={selectedType} bind:hidden={drawerHidden}/>