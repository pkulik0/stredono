<script lang="ts">
    import {Button, Hr, Img, TableBodyCell, TableBodyRow} from "flowbite-svelte";
    import {EditSolid, EyeSolid, PauseSolid, TrashBinSolid} from "flowbite-svelte-icons";
    import type {Alert} from "$lib/pb/user_pb";

    export let alert: Alert;

    let cellClass = "w-1/6";
    let iconClass = "w-4 h-4 mr-2";
    let btnClass = "w-1/2";

</script>

<TableBodyRow>
    <TableBodyCell class="{cellClass}">
        <Img src={alert.style?.gifUrl ?? ""} alt="{alert.id} GIF" class="rounded w-100 min-w-[100px] md:min-w-[150px]" />
    </TableBodyCell>
    <TableBodyCell class="{cellClass}">
        <audio controls>
            <source src={alert.style?.soundUrl ?? ""} type="audio/mpeg" />
            Your browser does not support the audio element.
        </audio>
    </TableBodyCell>
    <TableBodyCell class="{cellClass}">{alert.template}</TableBodyCell>
    <TableBodyCell class="{cellClass}">{alert.amountTrigger?.min || '?'}</TableBodyCell>
    <TableBodyCell class="{cellClass}">{alert.amountTrigger?.max || '?'}</TableBodyCell>
    <TableBodyCell class="{cellClass}">
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