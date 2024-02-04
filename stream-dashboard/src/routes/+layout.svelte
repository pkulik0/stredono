<script lang="ts">
    import '../app.pcss';
    import NavBar from "./NavBar.svelte";
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/firebase.js";
    import {userStore} from "$lib/userStore";
    import {goto} from "$app/navigation";

    if (localStorage.getItem('color-theme') === 'dark' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark');
    } else {
        document.documentElement.classList.remove('dark')
    }

    onAuthStateChanged(auth, async (u) => {
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

<NavBar />

<div class="p-4 flex justify-center">
    <div class="w-full md:w-3/4">
        <slot />
    </div>
</div>