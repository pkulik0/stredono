<script lang="ts">
    import NavBar from "./NavBar.svelte";
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/firebase.js";
    import {userStore} from "$lib/stores";
    import {goto} from "$app/navigation";

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
        <div class="w-full md:w-3/4">
            <slot />
        </div>
    </div>
</div>