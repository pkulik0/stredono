<script lang="ts">
    import {Checkbox, Helper, Hr, Input, Label} from "flowbite-svelte";
    import {slide} from "svelte/transition";

    export let checked: boolean = false;
    export let label: string = "Checkbox";

    export let size: "sm" | "md" | "lg" = "md";
    const textSize = size === "sm" ? "text-sm" : size === "md" ? "text-md" : "text-lg";
    const padding = size === "sm" ? "p-2" : size === "md" ? "p-4" : "p-6";
</script>

<div class="space-y-2">
    <Checkbox class="{textSize}" bind:checked>{label}</Checkbox>
    {#if checked}
        <div class="{padding}" transition:slide>
            <slot/>
        </div>
    {/if}
</div>