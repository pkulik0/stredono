<script>
    import { page } from '$app/stores';
    import {Navbar, NavBrand, NavHamburger, NavLi, NavUl} from "flowbite-svelte";
    import {userStore} from "$lib/user";
    import {t} from 'svelte-i18n';
    import LangPicker from './LangPicker.svelte';

    let activeClass = 'text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent';
    let nonActiveClass = 'text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent';

    $: activeUrl = $page.url.pathname;
    const baseUrl = "/dashboard"
</script>

<Navbar color="primary">
    <NavBrand href={baseUrl}>
        <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">{$t("app_name")}</span>
    </NavBrand>
    <NavHamburger />
    <NavUl {activeUrl} {activeClass} {nonActiveClass}>
        {#if $userStore}
            <NavLi href={baseUrl}>{$t('dashboard')}</NavLi>
            <NavLi href={baseUrl+"/tips"}>{$t("tips")}</NavLi>
            <NavLi href={baseUrl+"/payments"}>{$t("payments")}</NavLi>
            <NavLi href={baseUrl+"/settings"}>{$t("settings")}</NavLi>
        {/if}
        <div class="sm:flex sm:justify-end">
            <LangPicker/>
        </div>
    </NavUl>
</Navbar>