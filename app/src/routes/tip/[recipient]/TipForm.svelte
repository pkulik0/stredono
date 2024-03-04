<script lang="ts">
    import LoginWithTwitch from '$lib/comp/LoginWithTwitch.svelte';
    import { Currency } from '$lib/pb/stredono_pb';
    import { Alert, Button, ButtonGroup, Hr, Input, InputAddon, Label, P, Textarea } from 'flowbite-svelte';
    import {ExclamationCircleSolid} from "flowbite-svelte-icons";
    import {emailStore, senderStore} from "$lib/stores";
    import {slide} from "svelte/transition";

    $: donateButtonColor = isEnabled ? "alternative" : "primary";

    export let amount: string;
    export let message: string;
    export let currency: Currency;

    $: currencyName = (() => {
        switch(currency) {
            case Currency.PLN:
                return "zł";
            default:
                return "?";
        }
    })();

    export let recipient: string;
    export let hasListeners: boolean;
    export let sendDonate: () => void;
    export let isEnabled: boolean;
</script>

<div class="space-y-6">
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
            <InputAddon>{currencyName}</InputAddon>
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

<Hr/>

<div class="space-y-4">
    <Label class="mt-6">
        Email
        <Input bind:value={$emailStore} />
    </Label>

    <P class="text-center font-bold">
        Or
    </P>

    <LoginWithTwitch/>
</div>