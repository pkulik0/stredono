<script lang="ts">
    import {donationStore, type DonationsMap, fetchOldDonations} from "$lib/donations";
    import { auth } from '$lib/firebase/firebase';
    import {Breadcrumb, BreadcrumbItem, Card, Checkbox, Helper, Hr, Input, Label, P, Pagination} from "flowbite-svelte";
    import {onMount} from "svelte";
    import DonationList from "$lib/comp/DonationList.svelte";
    import {ChevronLeftOutline, ChevronRightOutline} from "flowbite-svelte-icons";
    import {page} from "$app/stores";
    import {goto} from "$app/navigation";
    import {userStore} from "$lib/user";

    onMount(() => {
        return donationStore.subscribe((value) => {
            donations = value;
        });
    });

    let donations: DonationsMap = {};

    $: activePage = Number.parseInt($page.url.searchParams.get('page') || '1')
    $: label = donations[donationsIndex] ? `${donations[donationsIndex].startDate.toLocaleDateString()} - ${donations[donationsIndex].endDate.toLocaleDateString()}` : "??/??/???? - ??/??/????";
    $: pages =  [{
        active: false,
        url: `/donations?page=${activePage}`,
        label: label,
        name: label
    }];

    $: donationsIndex = activePage - 1;
    $: items = donations[donationsIndex] ? donations[donationsIndex].donate : [];

    let searchTerm = "";
    let searchInSender = true;
    let searchInEmail = true;
    let searchInMessage = true;
    $: filteredItems = items.filter((item) => {
        const term = searchTerm.toLowerCase().trim();
        const sender = item.Sender.toLowerCase();
        const email = item.Email.toLowerCase();
        const message = item.Message.toLowerCase();

        if (!searchInEmail && !searchInMessage && !searchInSender) return true;

        return (
            (searchInSender && sender.includes(term)) ||
            (searchInEmail && email.includes(term)) ||
            (searchInMessage && message.includes(term))
        );
    });

    $: {
        $page.url.searchParams.set('page', activePage.toString());
        goto($page.url.toString(), {replaceState: true})
    }

    const nextWeek = () => {
        if (activePage === 1) return;
        activePage--;
    };

    const previousWeek = async () => {
        const user = auth.currentUser;
        if (!user) throw new Error('Not logged in');

        const startDate = donations[donationsIndex].startDate
        if(user.metadata.creationTime) {
            const creationTime = new Date(user.metadata.creationTime);
            if (startDate < creationTime) {
                return;
            }
        }

        if(!donations[activePage]) {
            await fetchOldDonations(startDate, activePage);
        }
        activePage++;
    };

</script>

<div class="flex justify-center mt-4 space-y-4">

    <Card padding="xl" size="xl" class="flex-1">

        <div class="flex justify-center">
            <Pagination large {pages} on:previous={previousWeek} on:next={nextWeek} icon>
                <svelte:fragment slot="prev">
                    <span class="sr-only">Previous</span>
                    <ChevronLeftOutline class="w-6 h-6" />
                </svelte:fragment>
                <svelte:fragment slot="next">
                    <span class="sr-only">Next</span>
                    <ChevronRightOutline class="w-6 h-6" />
                </svelte:fragment>
            </Pagination>
        </div>

        <Label class="p-2">
            Search
            <Input type="text" bind:value={searchTerm}></Input>
        </Label>
        <div class="flex flex-row space-x-4 ms-4 mt-1">
            <Checkbox bind:checked={searchInSender} class="mb-4">
                Senders
            </Checkbox>
            <Checkbox bind:checked={searchInEmail} class="mb-4">
                Emails
            </Checkbox>
            <Checkbox bind:checked={searchInMessage} class="mb-4">
                Messages
            </Checkbox>
        </div>

        <Hr/>

        {#if filteredItems.length === 0}
            <P class="text-center">No donations in this period</P>
        {:else}
            <DonationList items={filteredItems} />
        {/if}
    </Card>
</div>