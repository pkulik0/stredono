<script>
    import { page } from '$app/stores';

    import {DarkMode, Navbar, NavBrand, NavHamburger, NavLi, NavUl, Toggle} from "flowbite-svelte";
    import {streamModeStore, userStore} from "$lib/stores";

    let activeClass = 'text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent';
    let nonActiveClass = 'text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent';

    let btnClass = "md:m-0 md:me-6 my-6"

    $: activeUrl = $page.url.pathname;
</script>

<Navbar color="primary">
    <NavBrand href="/panel">
        <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Stredono</span>
    </NavBrand>
    <NavHamburger />
    <NavUl {activeUrl} {activeClass} {nonActiveClass}>
        {#if $userStore}
            <NavLi href="/panel">Panel</NavLi>
            <NavLi href="/panel/donations">Donations</NavLi>
            <NavLi href="/panel/payments">Payments</NavLi>
            <NavLi href="/panel/settings">Settings</NavLi>
        {/if}
        <div class="sm:flex sm:justify-end">
            <DarkMode {btnClass} />
            {#if $userStore}
                <Toggle checked={$streamModeStore} on:click={() => { streamModeStore.update((v) => { return !v }) }}>Stream Mode</Toggle>
            {/if}
        </div>
    </NavUl>
</Navbar>