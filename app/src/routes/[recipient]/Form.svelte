<script lang="ts">
    import {Alert, Button, ButtonGroup, Hr, Input, InputAddon, Label, Textarea} from "flowbite-svelte";
    import {ExclamationCircleSolid} from "flowbite-svelte-icons";
    import {emailStore, senderStore} from "$lib/stores";
    import {slide} from "svelte/transition";

    $: blockClass = isBlocked ? "pointer-events-none blur-sm" : "";
    $: donateButtonColor = isBlocked ? "alternative" : "primary";

    export let amount: string;
    export let message: string;
    export let currency: string;

    export let recipient: string;
    export let hasListeners: boolean;
    export let isBlocked: boolean;
    export let sendDonate: () => void;
</script>

<Label class="mt-6">
    Email
    <Input bind:value={$emailStore} />
</Label>

<Hr/>

<div class="space-y-6 {blockClass}">
    <Label>
        Name
        <Input bind:value={$senderStore} />
    </Label>

    <div>
        <Label>
            Amount
        </Label>
        <ButtonGroup class="w-full">
            <Input type="number" bind:value={amount} min="1" />
            <InputAddon>{currency}</InputAddon>
        </ButtonGroup>
    </div>


    <Label>
        Message
        <Textarea bind:value={message} />
    </Label>

    <Button on:click={sendDonate} class="w-full" color="{donateButtonColor}">Donate</Button>

    {#if hasListeners === false}
        <div transition:slide>
            <Alert border>
                <ExclamationCircleSolid slot="icon"/>
                <span class="capitalize">{recipient}</span>
                is currently not displaying the donations on stream.
            </Alert>
        </div>
    {/if}
</div>