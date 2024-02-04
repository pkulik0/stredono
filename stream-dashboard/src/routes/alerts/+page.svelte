<script lang="ts">
    import AlertsTable from "./AlertsTable.svelte";
    import {Breadcrumb, BreadcrumbItem, Button} from "flowbite-svelte";
    import {PlusSolid} from "flowbite-svelte-icons";
    import AlertsDrawer from "./AlertsDrawer.svelte";

    let drawerHidden = true;

</script>

<Breadcrumb aria-label="Default breadcrumb example">
    <BreadcrumbItem href="/dashboard" home>Dashboard</BreadcrumbItem>
    <BreadcrumbItem href="/alerts">Alerts</BreadcrumbItem>
</Breadcrumb>


<AlertsTable/>

<div class="flex justify-end">
    <Button on:click={() => drawerHidden = false} class="mt-8">
        <PlusSolid class="w-4 h-4" />
        <span class="ms-2">New Alert</span>
    </Button>
</div>

<AlertsDrawer bind:hidden={drawerHidden}/>

