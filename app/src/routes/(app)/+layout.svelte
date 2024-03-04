<script lang="ts">
    import {onAuthStateChanged} from "firebase/auth";
    import {auth} from "$lib/ext/firebase/firebase";
    import {getUserListener, userStore} from "$lib/user";
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";
    import {notificationsStore} from "$lib/notifications";
    import Notification from "./Notification.svelte";
    import {fly} from "svelte/transition";

    onMount(() => {
        let userListUnsub: (() => void) | undefined;
        let authUnsub = onAuthStateChanged(auth, async (u) => {
            if(u) {
                userListUnsub = await getUserListener(u.uid);

                if (window.location.pathname === "/") {
                    await goto("/dashboard")
                }
                return
            }

            if(userListUnsub) {
                userListUnsub();
                userListUnsub = undefined;
            }
            userStore.set(null)

            if (window.location.pathname !== "/") {
                await goto("/")
            }
        });

        return () => {
            authUnsub();
            if(userListUnsub) {
                userListUnsub();
            }
        }
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