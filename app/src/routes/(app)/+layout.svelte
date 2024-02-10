<script lang="ts">
    import NavBar from "./NavBar.svelte";
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/firebase.js";
    import {userStore} from "$lib/stores";
    import {goto} from "$app/navigation";
    import {fade, fly} from "svelte/transition";
    import {page} from "$app/stores";
    import {Breadcrumb, BreadcrumbItem} from "flowbite-svelte";

    $: currentPage = $page.url;

    $: pagesOnPath = currentPage.pathname.split("/").filter(Boolean);
    $: console.log(pagesOnPath.join("/"))

    if (localStorage.getItem('color-theme') === 'dark' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark');
    } else {
        document.documentElement.classList.remove('dark')
    }

    onAuthStateChanged(auth, async (u) => {
        if(u) {
            userStore.set(u)
            if (window.location.pathname === "/") {
                await goto("/panel")
            }
        } else {
            userStore.set(null)
            if (window.location.pathname !== "/") {
                await goto("/")
            }
        }
    });
</script>

<div class="h-screen w-full">
    <NavBar />

    <div class="flex justify-center p-4">
        <div class="w-full md:w-3/4 space-y-4" >
            <Breadcrumb>
                {#each pagesOnPath as page, i}
                    <BreadcrumbItem active={i === pagesOnPath.length - 1} home={i === 0} href={'/' + pagesOnPath.slice(0, i+1).join('/')}>
                        {page[0].toUpperCase() + page.slice(1)}
                    </BreadcrumbItem>
                {/each}
            </Breadcrumb>

            <slot />
        </div>
    </div>
</div>