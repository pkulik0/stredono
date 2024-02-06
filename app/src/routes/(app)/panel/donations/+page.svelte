<script lang="ts">
    import {donationsStore} from "$lib/stores";
    import {Breadcrumb, BreadcrumbItem} from "flowbite-svelte";

    donationsStore.subscribe((donations) => {
        console.log(donations);
    });
</script>

<Breadcrumb aria-label="Default breadcrumb example">
    <BreadcrumbItem href="/panel" home>Panel</BreadcrumbItem>
    <BreadcrumbItem href="/donations" active>Donations</BreadcrumbItem>
</Breadcrumb>
