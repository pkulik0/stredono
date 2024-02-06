<script lang="ts">
    import {
        ImagePlaceholder,
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell, TextPlaceholder
    } from "flowbite-svelte";
    import {onMount} from "svelte";
    import {getAlerts} from "$lib/alerts";
    import type {Alert} from "$lib/pb/alerts_pb";
    import {userStore} from "$lib/stores";
    import AlertRow from "./AlertRow.svelte";

    let alerts: Alert[]|undefined = undefined;

    onMount(() => {
       return userStore.subscribe(async (user) => {
           alerts = user ? await getAlerts() : undefined;
       });
    });

</script>


<Table noborder>
    <caption class="p-5 rounded-lg mt-3 text-lg font-semibold text-left text-gray-900 bg-white dark:text-white dark:bg-gray-800 mb-4">
        Your alerts
        <p class="mt-1 text-sm font-normal text-gray-500 dark:text-gray-400">
            Alerts are displayed on stream when you receive a donation.
            Each alerts can have a different sound, GIF and message template.
            Below are the alerts you have created. You can edit, delete, disable and test them.
        </p>
    </caption>

    <TableHead>
        <TableHeadCell class="rounded-tl-lg bg-transparent">GIF</TableHeadCell>
        <TableHeadCell>Sound</TableHeadCell>
        <TableHeadCell>Template</TableHeadCell>
        <TableHeadCell>From</TableHeadCell>
        <TableHeadCell>To</TableHeadCell>
        <TableHeadCell class="rounded-tr-lg">Actions</TableHeadCell>
    </TableHead>
    <TableBody>
        {#if alerts !== undefined}
            {#each alerts as alert}
                <AlertRow {alert} />
            {/each}
            {#if alerts.length === 0}
                <TableBodyRow>
                    <TableBodyCell colspan="6" class="text-center">
                        You have no alerts yet.
                    </TableBodyCell>
                </TableBodyRow>
            {/if}
        {:else}
            <TableBodyRow>
                <TableBodyCell>
                    <ImagePlaceholder/>
                </TableBodyCell>
                <TableBodyCell colspan="5">
                    <TextPlaceholder/>
                </TableBodyCell>
            </TableBodyRow>
        {/if}
    </TableBody>
</Table>