<script lang="ts">
    import {donationStore, type DonationsMap, fetchOldDonations} from "$lib/donations";
    import {Breadcrumb, BreadcrumbItem, Card, Checkbox, Input, Label, P, Pagination} from "flowbite-svelte";
    import {onMount} from "svelte";
    import DonationList from "$lib/comp/DonationList.svelte";
    import {ChevronLeftOutline, ChevronRightOutline} from "flowbite-svelte-icons";
    import {page} from "$app/stores";
    import {goto} from "$app/navigation";
    import {userStore} from "$lib/stores";

    onMount(() => {
        return donationStore.subscribe((value) => {
            console.log('donationStore', value);
            donations = value;
        });
    });

    let donations: DonationsMap = {};

    $: activePage = Number.parseInt($page.url.searchParams.get('page') || '1')
    $: pages = Object.keys(donations).map((pageNum) => {
        const num = Number.parseInt(pageNum) + 1;
        const entry = donations[num - 1];

        const startDate = entry.startDate.toDateString();
        const endDate = entry.endDate.toDateString();
        const label = `${startDate} - ${endDate}`;

        return {
            active: num === activePage,
            url: `/donations?page=${num}`,
            label: label,
            name: label
        };
    }).reverse();

    $: donationsIndex = activePage - 1;
    $: items = donations[donationsIndex] ? donations[donationsIndex].donate : [];

    let searchTerm = "";
    let searchInSender = true;
    let searchInEmail = true;
    let searchInMessage = true;
    $: filteredItems = items.filter((item) => {
        const term = searchTerm.toLowerCase().trim();
        const sender = item.sender.toLowerCase();
        const email = item.email.toLowerCase();
        const message = item.message.toLowerCase();
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
        const user = $userStore;
        if (!user) throw new Error('Not logged in');

        const endDate = donations[donationsIndex].endDate
        if(user.metadata.creationTime) {
            const creationTime = new Date(user.metadata.creationTime);
            if (endDate < creationTime) {
                console.log('no more previous weeks');
                return;
            }
        }

        console.log('previousWeek');
        if(!donations[activePage]) {
            await fetchOldDonations(endDate, activePage);
        }
        activePage++;
    };

</script>

<Breadcrumb aria-label="Default breadcrumb example">
    <BreadcrumbItem href="/panel" home>Panel</BreadcrumbItem>
    <BreadcrumbItem href="/donations" active>Donations</BreadcrumbItem>
</Breadcrumb>

<div class="flex justify-center flex-col mt-4 space-y-4">
    <div class="flex justify-center">
        <Pagination {pages} on:previous={previousWeek} on:next={nextWeek} icon>
            <svelte:fragment slot="prev">
                <span class="sr-only">Previous</span>
                <ChevronLeftOutline class="w-2.5 h-2.5" />
            </svelte:fragment>
            <svelte:fragment slot="next">
                <span class="sr-only">Next</span>
                <ChevronRightOutline class="w-2.5 h-2.5" />
            </svelte:fragment>
        </Pagination>
    </div>

    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Donations</h5>

    <Card padding="xl" size="xl" class="flex-1">

        <Label class="p-2">
            Search
            <Input type="text" bind:value={searchTerm}></Input>
        </Label>
        <div class="flex flex-row space-x-4 ms-4 mb-2 mt-1">
            <P class="mb-2">Look in:</P>
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

        <DonationList items={filteredItems} />
    </Card>
</div>