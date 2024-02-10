<script>
    import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper, SidebarDropdownItem, SidebarDropdownWrapper } from 'flowbite-svelte';
    import {
        UserSolid, MessageDotsSolid, BellSolid, VolumeUpSolid, UsersGroupSolid, BarsSolid, DollarSolid, WalletSolid
    } from 'flowbite-svelte-icons';
    import {page} from "$app/stores";

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

            <SidebarDropdownWrapper isOpen label="Chat Bot">
                <svelte:fragment slot="icon">
                    <MessageDotsSolid class={iconClass} />
                </svelte:fragment>
                <svelte:fragment slot="arrowup"><span></span></svelte:fragment>
                <svelte:fragment slot="arrowdown"><span></span></svelte:fragment>

                <SidebarDropdownItem label="Rules" href="{baseUrl}/chatbot/rules" />
                <SidebarDropdownItem label="Filters" href="{baseUrl}/chatbot/filters" />
                <SidebarDropdownItem label="Commands" href="{baseUrl}/chatbot/commands" />
            </SidebarDropdownWrapper>

            <SidebarItem label="Moderators" href="{baseUrl}/mods">
                <svelte:fragment slot="icon">
                    <UsersGroupSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>

            <SidebarItem label="Other" href="{baseUrl}/other">
                <svelte:fragment slot="icon">
                    <BarsSolid class={iconClass} />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>