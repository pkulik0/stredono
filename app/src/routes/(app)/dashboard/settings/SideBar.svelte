<script>
    import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper} from 'flowbite-svelte';
    import {
        UserSolid,
        BellSolid,
        VolumeUpSolid,
        ArrowRightToBracketSolid, NewspapperOutline
    } from 'flowbite-svelte-icons';
    import {page} from "$app/stores";
    import {auth} from "$lib/ext/firebase/firebase";
    import {t} from 'svelte-i18n';

    let iconClass = "w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-200 group-hover:text-gray-900 dark:group-hover:text-white"
    $: activeUrl = $page.url.pathname;
    let activeClass = 'flex items-center p-2 text-base font-normal text-primary-900 bg-primary-200 dark:bg-gray-700 rounded-lg dark:text-white hover:bg-primary-100 dark:hover:bg-gray-600';

    let baseUrl = "/dashboard/settings";
</script>

<Sidebar {activeUrl} {activeClass} class="rounded-xl w-full max-w-3xl">
    <SidebarWrapper>
        <SidebarGroup>
            <SidebarItem label={$t("profile")} href="{baseUrl}/profile">
                <svelte:fragment slot="icon">
                    <UserSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label={$t("events")} href="{baseUrl}/events">
                <svelte:fragment slot="icon">
                    <NewspapperOutline class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label={$t("alerts")} href="{baseUrl}/alerts">
                <svelte:fragment slot="icon">
                    <BellSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label={$t("tts")} href="{baseUrl}/tts">
                <svelte:fragment slot="icon">
                    <VolumeUpSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>

        <SidebarGroup border>
            <SidebarItem on:click={async () => { await auth.signOut(); }} label={$t("sign_out")}>
                <svelte:fragment slot="icon">
                    <ArrowRightToBracketSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>