<script lang="ts">
    import ConfirmationModal from '$lib/comp/ConfirmationModal.svelte';
    import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper} from 'flowbite-svelte';
    import {
        UserSolid,
        BellSolid,
        VolumeUpSolid,
        ArrowRightToBracketSolid, NewspapperOutline, CameraFotoOutline, VideoCameraSolid, UsersGroupSolid
    } from 'flowbite-svelte-icons';
    import {page} from "$app/stores";
    import {auth} from "$lib/ext/firebase/firebase";
    import {t} from 'svelte-i18n';

    let iconClass = "w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-200 group-hover:text-gray-900 dark:group-hover:text-white"
    $: activeUrl = $page.url.pathname;
    let activeClass = 'flex items-center p-2 text-base font-normal text-primary-900 bg-primary-200 dark:bg-gray-700 rounded-lg dark:text-white hover:bg-primary-100 dark:hover:bg-gray-600';

    let baseUrl = "/dashboard/settings";

    let showConfirmation = false;
    const onSignOutClick = async (confirm: boolean) => {
        if (!confirm) {
            showConfirmation = true;
            return;
        }

        await auth.signOut();
    }
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

            <SidebarItem label={$t("overlay")} href="{baseUrl}/overlay">
                <svelte:fragment slot="icon">
                    <VideoCameraSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label={$t("moderators")} href="{baseUrl}/moderators">
                <svelte:fragment slot="icon">
                    <UsersGroupSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>

        <SidebarGroup border>
            <SidebarItem on:click={() => onSignOutClick(false)} label={$t("sign_out")}>
                <svelte:fragment slot="icon">
                    <ArrowRightToBracketSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>

<ConfirmationModal bind:open={showConfirmation} onConfirm={() => onSignOutClick(true)} />