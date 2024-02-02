<script lang="ts">
    import {onMount} from "svelte";
    import {get} from "svelte/store";
    import {goto} from "$app/navigation";
    import {userStore} from "$lib/userStore";

    onMount(async () => {
        const currentUser = get(userStore);
        if (currentUser) {
            await goto("/dashboard");
        } else {
            await goto("/login");
        }
    })
</script>
