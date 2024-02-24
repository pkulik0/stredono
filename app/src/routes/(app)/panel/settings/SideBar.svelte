<script>
    import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper} from 'flowbite-svelte';
    import {
        UserSolid,
        MessageDotsSolid,
        BellSolid,
        VolumeUpSolid,
        DollarSolid,
        WalletSolid,
        ArrowRightToBracketSolid
    } from 'flowbite-svelte-icons';
    import {page} from "$app/stores";
    import {auth} from "$lib/ext/firebase/firebase";

    let iconClass = "w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-200 group-hover:text-gray-900 dark:group-hover:text-white"
    $: activeUrl = $page.url.pathname;
    let activeClass = 'flex items-center p-2 text-base font-normal text-primary-900 bg-primary-200 dark:bg-gray-700 rounded-lg dark:text-white hover:bg-primary-100 dark:hover:bg-gray-600';

    let baseUrl = "/panel/settings";
</script>

<Sidebar {activeUrl} {activeClass} class="rounded-xl">
    <SidebarWrapper>
        <SidebarGroup>
            <SidebarItem label="Profile" href="{baseUrl}/profile">
                <svelte:fragment slot="icon">
                    <UserSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label="Donations" href="{baseUrl}/donations">
                <svelte:fragment slot="icon">
                    <DollarSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label="Payments" href="{baseUrl}/payments">
                <svelte:fragment slot="icon">
                    <WalletSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
        <SidebarGroup border>
            <SidebarItem label="Twitch" href="{baseUrl}/twitch">
                <svelte:fragment slot="icon">
                    <MessageDotsSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label="Alerts" href="{baseUrl}/alerts">
                <svelte:fragment slot="icon">
                    <BellSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label="Text-to-Speech" href="{baseUrl}/tts">
                <svelte:fragment slot="icon">
                    <VolumeUpSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>

        <SidebarGroup border>
            <SidebarItem on:click={async () => { await auth.signOut(); }} label="Sign Out">
                <svelte:fragment slot="icon">
                    <ArrowRightToBracketSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>