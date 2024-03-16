<script lang="ts">
    import { getEventsDashboardListener } from '$lib/events';
    import { getSettingsListener } from '$lib/settings';
    import { userStore } from '$lib/user';
    import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
    import {page} from "$app/stores";
    import { onMount } from 'svelte';
    import NavBar from './NavBar.svelte';
    import {t} from 'svelte-i18n';

    $: currentPage = $page.url;
    $: pagesOnPath = currentPage.pathname.split("/").filter(Boolean);

    onMount(() => {
        let settingsUnsub: (() => void) | undefined;
        let eventsUnsub: (() => void) | undefined;
        const userUnsub = userStore.subscribe(u => {
            if(!u) {
                if(settingsUnsub) {
                    settingsUnsub();
                    settingsUnsub = undefined;
                }
                if(eventsUnsub) {
                    eventsUnsub();
                    eventsUnsub = undefined;
                }
                return;
            }
            settingsUnsub = getSettingsListener(u.Uid)
            eventsUnsub = getEventsDashboardListener(u.Uid);
        })

        return () => {
            userUnsub();
            if(settingsUnsub) {
                settingsUnsub();
            }
            if(eventsUnsub) {
                eventsUnsub();
            }
        }
    })
</script>

<NavBar />

<div class="flex justify-center p-4">
    <div class="w-full max-w-7xl space-y-4" >
        <Breadcrumb>
            {#each pagesOnPath as page, i}
                <BreadcrumbItem active={i === pagesOnPath.length - 1} home={i === 0} href={'/' + pagesOnPath.slice(0, i+1).join('/')}>
                    {$t(page)}
                </BreadcrumbItem>
            {/each}
        </Breadcrumb>

<!--        <AlertBox/>-->

        <slot />
    </div>
</div>