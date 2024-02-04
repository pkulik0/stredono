<script lang="ts">
    import {
        Button,
        Hr, ImagePlaceholder,
        Img,
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell, TextPlaceholder
    } from "flowbite-svelte";
    import {onMount} from "svelte";
    import {getAlerts} from "$lib/alerts";
    import type {Alert} from "../../../../pb/alerts_pb";
    import {userStore} from "$lib/userStore";
    import {EditSolid, EyeSolid, PauseSolid, TrashBinSolid} from "flowbite-svelte-icons";

    let alerts: Alert[]|undefined = undefined;

    onMount(() => {
       return userStore.subscribe(async (user) => {
           alerts = user ? await getAlerts() : undefined;
       });
    });

    let btnClass = "w-1/2";
    let iconClass = "w-4 h-4 mr-2";
</script>


<Table>
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
                <TableBodyRow>
                    <TableBodyCell>
                        <Img src={alert.gifUrl} alt="{alert.id} GIF" class="w-40 rounded" />
                    </TableBodyCell>
                    <TableBodyCell>
                        <audio controls>
                            <source src={alert.soundUrl} type="audio/mpeg" />
                            Your browser does not support the audio element.
                        </audio>
                    </TableBodyCell>
                    <TableBodyCell>{alert.template}</TableBodyCell>
                    <TableBodyCell>{alert.from}</TableBodyCell>
                    <TableBodyCell>{alert.to}</TableBodyCell>
                    <TableBodyCell>
                        <Button color="red" outline class="{btnClass}">
                            <TrashBinSolid class="{iconClass}" />
                            Delete
                        </Button>
                        <Button color="alternative" class="{btnClass}">
                            <PauseSolid class="{iconClass}" />
                            Disable
                        </Button>

                        <Hr/>

                        <Button outline color="green" class="{btnClass}">
                            <EyeSolid class="{iconClass}" />
                            Test
                        </Button>

                        <Button outline color="blue" class="{btnClass}">
                            <EditSolid class="{iconClass}" />
                            Edit
                        </Button>
                    </TableBodyCell>
                </TableBodyRow>
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