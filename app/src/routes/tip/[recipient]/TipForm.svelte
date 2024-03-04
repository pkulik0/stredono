<script lang="ts">
    import LoginWithTwitch from '$lib/comp/LoginWithTwitch.svelte';
    import { Currency } from '$lib/pb/stredono_pb';
    import { Alert, Button, ButtonGroup, Hr, Input, InputAddon, Label, P, Textarea } from 'flowbite-svelte';
    import {ExclamationCircleSolid} from "flowbite-svelte-icons";
    import {emailStore, senderStore} from "$lib/stores";
    import {slide} from "svelte/transition";
    import {t} from "svelte-i18n";

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
    export let isEnabled: boolean = true;
</script>

<div class="space-y-6">
    <div class="space-y-4">
        <Label class="mt-6">
            {$t("email")}
            <Input bind:value={$emailStore} />
        </Label>

        <!--    <P class="text-center font-medium">-->
        <!--        {$t("or")}-->
        <!--    </P>-->

        <!--    <LoginWithTwitch/>-->
    </div>

    <Label>
        {$t("display_name")}
        <Input bind:value={$senderStore} />
    </Label>

    <div>
        <Label>
            {$t("amount")}
        </Label>
        <ButtonGroup class="w-full">
            <Input type="number" bind:value={amount} min="1" />
            <InputAddon>{currencyName}</InputAddon>
        </ButtonGroup>
    </div>


    <Label>
        {$t("message")}
        <Textarea bind:value={message} />
    </Label>

    <Button on:click={sendDonate} class="w-full" color="{donateButtonColor}">{$t("continue")}</Button>

    {#if hasListeners === true}
        <div transition:slide>
            <Alert border>
                <ExclamationCircleSolid slot="icon"/>
                {$t("alerts_disabled_label", { values: { user: recipient }})}
            </Alert>
        </div>
    {/if}
</div>