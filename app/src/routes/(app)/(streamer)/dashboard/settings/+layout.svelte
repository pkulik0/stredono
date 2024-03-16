<script lang="ts">
    import { goto } from '$app/navigation';
    import { page } from '$app/stores';
    import { getSettingsListener, settingsStore } from '$lib/settings';
    import { auth } from '$lib/ext/firebase/firebase';
    import { userStore } from '$lib/user';
    import { onMount } from 'svelte';
    import SideBar from "./SideBar.svelte";

    $: isOnMain = $page.url.pathname.endsWith("/settings")
    $: sideBarClass = isOnMain ? "flex" : "hidden"

    const handleResize = () => {
        const isSmallDevice = window.innerWidth < 768;
        const isMainPage = $page.url.pathname.endsWith("/settings") // due to the way Svelte updates values the one is inOnMage is outdated
        if(!isSmallDevice && isMainPage) {
            goto('/dashboard/settings/profile');
        }
    }

    onMount(() => {
        window.addEventListener('resize', handleResize);
        const unsub = page.subscribe((value) => {
            handleResize();
        });

        return () => {
            window.removeEventListener('resize', handleResize);
            unsub();
        }
    });
</script>

<div class="flex px-4">
    <div class="{sideBarClass} w-full md:pe-4 md:flex md:max-w-72">
        <SideBar/>
    </div>

    {#if !isOnMain}
        <div class="flex flex-col mx-auto max-w-3xl w-full items-center">
            <slot/>
        </div>
    {/if}
</div>