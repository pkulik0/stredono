<script lang="ts">
    import {tipsStore, type TipsMap, fetchOldTips} from "../../../../../lib/tips";
    import { auth } from '../../../../../lib/ext/firebase/firebase';
    import {Card, Checkbox, Hr, Input, Label, P, Pagination} from "flowbite-svelte";
    import {onMount} from "svelte";
    import TipList from "../../../../../lib/comp/TipList.svelte";
    import {ChevronLeftOutline, ChevronRightOutline} from "flowbite-svelte-icons";
    import {page} from "$app/stores";
    import {goto} from "$app/navigation";
    import {t} from "svelte-i18n";

    onMount(() => {
        return tipsStore.subscribe((value) => {
            tipsMap = value;
        });
    });

    let tipsMap: TipsMap = {};

    $: activePage = Number.parseInt($page.url.searchParams.get('page') || '1')
    $: label = tipsMap[donationsIndex] ? `${tipsMap[donationsIndex].startDate.toLocaleDateString()} - ${tipsMap[donationsIndex].endDate.toLocaleDateString()}` : "??/??/???? - ??/??/????";
    $: pages =  [{
        active: false,
        url: `/tips?page=${activePage}`,
        label: label,
        name: label
    }];

    $: donationsIndex = activePage - 1;
    $: items = tipsMap[donationsIndex] ? tipsMap[donationsIndex].tips : [];

    let searchTerm = "";
    let searchInSender = true;
    let searchInEmail = true;
    let searchInMessage = true;
    $: filteredItems = items.filter((item) => {
        const term = searchTerm.toLowerCase().trim();
        const sender = item.DisplayName.toLowerCase();
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

        const startDate = tipsMap[donationsIndex].startDate
        if(user.metadata.creationTime) {
            const creationTime = new Date(user.metadata.creationTime);
            if (startDate < creationTime) {
                return;
            }
        }

        if(!tipsMap[activePage]) {
            await fetchOldTips(startDate, activePage);
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
            {$t("search")}
            <Input type="text" bind:value={searchTerm}></Input>
        </Label>
        <div class="flex flex-row space-x-4 ms-3">
            <P color="gray">{$t("filter")}:</P>
            <Checkbox bind:checked={searchInSender} class="mt-1">
                {$t("display_names")}
            </Checkbox>
            <Checkbox bind:checked={searchInEmail} class="mt-1">
                {$t("emails")}
            </Checkbox>
            <Checkbox bind:checked={searchInMessage} class="mt-1">
                {$t("messages")}
            </Checkbox>
        </div>

        <Hr/>

        {#if filteredItems.length === 0}
            <P class="text-center">{$t("no_tips_in_period")}</P>
        {:else}
            <TipList items={filteredItems} />
        {/if}
    </Card>
</div>