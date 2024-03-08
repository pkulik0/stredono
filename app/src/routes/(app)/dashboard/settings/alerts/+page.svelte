<script lang="ts">
    import { alertsStore, getAlertsListener } from '$lib/alerts';
    import { EventType } from '$lib/pb/event_pb';
    import { UsersAlerts } from '$lib/pb/alert_pb';
    import { pbEnumToItems } from '$lib/util';
    import {PlusSolid} from "flowbite-svelte-icons";
    import AlertCard from './AlertCard.svelte';
    import AlertsDrawer from "./AlertsDrawer.svelte";
    import {
        Button, Heading,
        ImagePlaceholder, Label, P, Select
    } from 'flowbite-svelte';
    import {onMount} from "svelte";
    import {userStore} from "$lib/user";
    import {t} from "svelte-i18n";

    let userAlerts: UsersAlerts|undefined|null;

    $: eventTypesItems = pbEnumToItems(EventType)
    let selectedType: EventType = EventType.TIP;

    let drawerHidden = true;

    onMount(() => {
        let listenerUnsub: (() => void)|undefined;
        const userUnsub = userStore.subscribe(async (user) => {
            if(!user) {
                userAlerts = undefined;
                return;
            }
            listenerUnsub = await getAlertsListener(user.Uid)
        });
        const alertsUnsub = alertsStore.subscribe((alertsData) => {
            userAlerts = alertsData;
        })
        return () => {
            if(listenerUnsub) listenerUnsub();
            userUnsub();
            alertsUnsub();
        }
    });
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

    <div class="p-4">
        {#if userAlerts !== undefined}
            {#if userAlerts === null}
                <P class="text-center py-8">
                    {$t("no_alerts")}
                </P>
            {:else}
                {#each userAlerts.Alerts as alert}
                    <AlertCard {alert}/>
                {/each}
            {/if}
        {:else}
            <ImagePlaceholder/>
        {/if}
    </div>
</div>

<AlertsDrawer eventType={selectedType} bind:hidden={drawerHidden}/>