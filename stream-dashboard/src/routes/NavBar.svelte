<script>
    import { page } from '$app/stores';

    import {DarkMode, Navbar, NavBrand, NavHamburger, NavLi, NavUl} from "flowbite-svelte";

    let activeClass = 'text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent';
    let nonActiveClass = 'text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent';
    let btnClass = "m-0"

    $: activeUrl = $page.url.pathname;

</script>

<Navbar>
    <NavBrand href="/">
        <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Stredono</span>
    </NavBrand>
    <NavHamburger />
    <NavUl {activeUrl} {activeClass} {nonActiveClass}>
        <NavLi href="/">Dashboard</NavLi>
        <NavLi href="/docs/components/accordion">Donations</NavLi>
        <NavLi href="/docs/components/navbar">Settings</NavLi>
        <DarkMode {btnClass} />
    </NavUl>
</Navbar>