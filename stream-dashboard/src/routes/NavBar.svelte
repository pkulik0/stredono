<script>
    import { page } from '$app/stores';

    import {DarkMode, Navbar, NavBrand, NavHamburger, NavLi, NavUl} from "flowbite-svelte";
    import {auth} from "$lib/firebase";
    import {goto} from "$app/navigation";
    import {userStore} from "$lib/userStore";

    let activeClass = 'text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent';
    let nonActiveClass = 'text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent';
    let btnClass = "m-0"

    $: activeUrl = $page.url.pathname;

    const signOut = async () => {
        await auth.signOut()
    }

</script>

<Navbar>
    <NavBrand href="/">
        <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Stredono</span>
    </NavBrand>
    <NavHamburger />
    <NavUl {activeUrl} {activeClass} {nonActiveClass}>
        {#if $userStore}
            <NavLi href="/dashboard">Dashboard</NavLi>
            <NavLi href="/donations">Donations</NavLi>
            <NavLi href="/alerts">Alerts</NavLi>
            <NavLi href="/payments">Payments</NavLi>
            <NavLi href="/settings">Settings</NavLi>
            <NavLi on:click={signOut}>Sign out</NavLi>
        {/if}
        <DarkMode {btnClass} />
    </NavUl>
</Navbar>