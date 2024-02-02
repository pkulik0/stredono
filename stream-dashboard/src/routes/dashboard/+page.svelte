<script lang="ts">
    import {onMount} from "svelte";
    import {get} from "svelte/store";
    import {userStore} from "$lib/userStore";
    import {goto} from "$app/navigation";

    onMount(async () => {
        if (!get(userStore)) {
            await goto("/login");
        }
    })
</script>

DASHBOARD