<script lang="ts">
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/ext/firebase/firebase";
    import { getUserByUid, userStore } from '$lib/user';
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";
    import {notificationsStore} from "$lib/notifications";
    import Notification from "./Notification.svelte";
    import {fly} from "svelte/transition";

    onMount(() => {
        return onAuthStateChanged(auth, async (u) => {
            if(u) {
                const user = await getUserByUid(u.uid)
                userStore.set(user)

                if(user.Username === "") {
                    await goto("/intro")
                }

                if (window.location.pathname === "/") {
                    await goto("/dashboard")
                }
            } else {
                userStore.set(undefined)
                if (window.location.pathname !== "/") {
                    await goto("/")
                }
            }
        });
    })
</script>

<div class="h-screen w-full">
    <slot />
</div>

{#each $notificationsStore as notification}
    <div transition:fly>
        <Notification {notification} />
    </div>
{/each}