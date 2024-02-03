<script lang="ts">
    import '../app.pcss';
    import NavBar from "./NavBar.svelte";
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/firebase.js";
    import {userStore} from "$lib/userStore";
    import type {OptionalUser} from "$lib/userStore";
    import {goto} from "$app/navigation";

    if (localStorage.getItem('color-theme') === 'dark' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark');
    } else {
        document.documentElement.classList.remove('dark')
    }

    onAuthStateChanged(auth, async (u: OptionalUser) => {
        console.log("User state changed", u);
        if(u) {
            userStore.set(u)
            if (window.location.pathname === "/") {
                await goto("/dashboard")
            }
        } else {
            userStore.set(null)
            if (window.location.pathname !== "/") {
                await goto("/")
            }
        }
    });
</script>


<div class="flex flex-col min-h-screen">
    <NavBar />

    <div class="flex flex-1 justify-center p-4">
        <div class="max-w-4xl w-full">
            <slot />
        </div>
    </div>
</div>