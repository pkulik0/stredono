<script lang="ts">
    import { Card, TextPlaceholder} from "flowbite-svelte";
    import type {SendDonateRequest} from "$lib/pb/functions_pb";
    import {onMount} from "svelte";
    import {donationStore} from "$lib/donations";
    import DonationList from "$lib/comp/DonationList.svelte";

    let items: SendDonateRequest[]|undefined = undefined;

    export let size = 4;

    onMount(() => {
        return donationStore.subscribe(donations => {
            if(!donations.hasOwnProperty(0)) return;
            const latestDonations = donations[0].donate;

            let startIndex = latestDonations.length - size;
            if (startIndex < 0) startIndex = 0;
            items = latestDonations.slice(startIndex).reverse();
        });
    })

</script>

<Card padding="xl" size="xl" class="mt-4 flex-1">
    <div class="flex justify-between items-center mb-4">
        <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">Latest Donations</h5>
        <a href="/panel/donations" class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"> View all </a>
    </div>

    {#if items !== undefined}
        <DonationList {items}/>
    {:else}
        <TextPlaceholder/>
    {/if}
</Card>