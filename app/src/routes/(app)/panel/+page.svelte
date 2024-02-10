<script lang="ts">
    import LatestCard from "./LatestCard.svelte";
    import {Alert, Card} from "flowbite-svelte";
    import {InfoCircleSolid} from "flowbite-svelte-icons";
</script>

<div class="flex flex-col space-y-2">
    <Alert class="mt-4">
        <InfoCircleSolid slot="icon" class="w-4 h-4" />
        <span class="font-medium">Default alert!</span>
        System messages are displayed here.
    </Alert>
    <div class="flex justify-center space-x-2">
        <LatestCard/>
    </div>
</div>