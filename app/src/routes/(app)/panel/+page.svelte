<script lang="ts">
    import {Breadcrumb, BreadcrumbItem, Card, Toggle} from "flowbite-svelte";
    import LatestCard from "./LatestCard.svelte";
</script>

<Breadcrumb aria-label="Default breadcrumb example">
    <BreadcrumbItem href="/panel" home>Panel</BreadcrumbItem>
</Breadcrumb>

<div class="flex justify-center">
    <LatestCard/>
</div>