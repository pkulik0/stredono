<script lang="ts">
    import NavBar from "./NavBar.svelte";
    import {page} from "$app/stores";
    import {Breadcrumb, BreadcrumbItem, Toast} from "flowbite-svelte";
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/firebase/firebase";
    import {getUserListener, userStore} from "$lib/user";
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";
    import {notificationsStore} from "$lib/notifications";
    import Notification from "./Notification.svelte";
    import {DownloadOutline} from "flowbite-svelte-icons";
    import {fly} from "svelte/transition";


    let userListUnsub: (() => void) | undefined;

    onMount(() => {
        if (localStorage.getItem('color-theme') === 'dark' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            document.documentElement.classList.add('dark');
        } else {
            document.documentElement.classList.remove('dark')
        }

        onAuthStateChanged(auth, async (u) => {
            if(u) {
                userListUnsub = await getUserListener(u.uid);

                if (window.location.pathname === "/") {
                    await goto("/panel")
                }
                return
            }

            if(userListUnsub) {
                userListUnsub();
                userListUnsub = undefined;
            }
            userStore.set(undefined)

            if (window.location.pathname !== "/") {
                await goto("/")
            }
        });
    })

    $: currentPage = $page.url;
    $: pagesOnPath = currentPage.pathname.split("/").filter(Boolean);
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

{#each $notificationsStore as notification}
    <div transition:fly>
        <Notification {notification} />
    </div>
{/each}