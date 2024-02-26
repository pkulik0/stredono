<script lang="ts">
    import { alertsStore, getAlertsListener } from '$lib/alerts';
    import { EventType, UsersAlerts } from '$lib/pb/stredono_pb';
    import {PlusSolid} from "flowbite-svelte-icons";
    import AlertCard from './AlertCard.svelte';
    import AlertsDrawer from "./AlertsDrawer.svelte";
    import {
        Button,
        ImagePlaceholder, Input, Label, P, Select
    } from 'flowbite-svelte';
    import {onMount} from "svelte";
    import {userStore} from "$lib/user";
    import AlertViewer from './AlertViewer.svelte';

    let userAlerts: UsersAlerts|undefined|null;

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

    let drawerHidden = true;

    let selectOptions = [
        {value: EventType.TIP, name: 'Tip'},
        {value: EventType.CHEER, name: 'Cheer'},
        {value: EventType.SUBSCRIBE, name: 'Sub'},
        {value: EventType.SUBGIFT, name: 'Sub Gift'},
        {value: EventType.FOLLOW, name: 'Follow'},
        {value: EventType.RAID, name: 'Raid'},
    ];
    let selectedType: EventType = EventType.TIP;

</script>

<div class="space-y-4">
    <Label>
        Type
        <Select items={selectOptions} bind:value={selectedType} placeholder="Filter by type"/>
    </Label>

    <div class="flex justify-end">
        <Button on:click={() => drawerHidden = false}>
            <PlusSolid class="w-4 h-4" />
            <span class="ms-2">New</span>
        </Button>
    </div>

    {#if userAlerts !== undefined}
        {#if userAlerts === null}
            <P class="text-center py-8">
                You have no alerts of this type.
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

<AlertsDrawer eventType={selectedType} bind:hidden={drawerHidden}/>