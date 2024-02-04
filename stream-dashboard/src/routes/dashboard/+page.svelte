<script lang="ts">
    import {Breadcrumb, BreadcrumbItem, Card} from "flowbite-svelte";
    import LatestCard from "./LatestCard.svelte";
</script>

<Breadcrumb aria-label="Default breadcrumb example">
    <BreadcrumbItem href="/dashboard" home>Dashboard</BreadcrumbItem>
</Breadcrumb>


<div class="flex justify-center">
    <LatestCard/>
</div>