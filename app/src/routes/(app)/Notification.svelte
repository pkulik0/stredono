<script lang="ts">
    import {Toast} from "flowbite-svelte";
    import {BadgeCheckOutline, BellOutline, ExclamationCircleOutline, CloseOutline} from "flowbite-svelte-icons";
    import {fly} from "svelte/transition";
    import {onMount} from "svelte";
    import {type Notification} from "$lib/notifications";

    export let notification: Notification;

    let color: "green" | "red" | "yellow" | "blue" = "green";
    const iconClass = "w-6 h-6";

    let open = true;
    let counter: number;

    onMount(() => {
        if (notification.type === "success") {
            color = "green";
        } else if (notification.type === "error") {
            color = "red";
        } else if (notification.type === "warning") {
            color = "yellow";
        } else if (notification.type === "info") {
            color = "blue";
        }

        if(notification.duration >= 1) {
            counter = notification.duration;

            const interval = setInterval(() => {
                counter--;
                if (counter === 0) {
                    clearInterval(interval);
                    open = false;
                    notification.dismiss();
                }
            }, 1000);
        }
    });
</script>

<Toast position={notification.position} bind:open dismissable={notification.dismissible} transition={fly} class="mb-4" color="green">
    <svelte:fragment slot="icon">
        {#if notification.type === "success"}
            <BadgeCheckOutline class="{iconClass}"/>
        {:else if notification.type === "error"}
            <CloseOutline class="{iconClass}"/>
        {:else if notification.type === "warning"}
            <ExclamationCircleOutline class="{iconClass}"/>
        {:else if notification.type === "info"}
            <BellOutline class="{iconClass}"/>
        {/if}
        <span class="sr-only">{notification.type}</span>
    </svelte:fragment>

    {notification.message}
    <span class="text-gray-500">({counter}s)</span>
</Toast>