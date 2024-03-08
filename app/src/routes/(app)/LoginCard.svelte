<script>
    import LoginWithTwitch from '$lib/comp/LoginWithTwitch.svelte';
    import { Button, Helper, Input, Label, P } from 'flowbite-svelte';
    import {isSignInWithEmailLink, sendSignInLinkToEmail, signInWithEmailLink} from "firebase/auth";
    import {auth} from "$lib/ext/firebase/firebase";
    import {onMount} from "svelte";
    import {userStore} from "$lib/user";
    import {goto} from "$app/navigation";
    import { t } from 'svelte-i18n';
    import {slide} from 'svelte/transition';

    let email = "";
    const emailKey = "sign-in-email";

    let success = false;

    const onTosDetailsClicked = () => {
        console.log("tos details clicked");
        // TODO: impl
    }

    const login = async () => {
        try {
            const actionCodeSettings = {
                url: window.location.href,
                handleCodeInApp: true,
            };
            await sendSignInLinkToEmail(auth, email, actionCodeSettings);
            window.localStorage.setItem(emailKey, email);

            success = true;
        } catch (e) {
            console.error(e);
        }
    }

    const checkEmail = () => {
        const domain = email.split("@")[1];
        window.open(`https://${domain}`, "_blank");
    }

    onMount(async () => {
        if($userStore) await goto("/panel");

        if (!isSignInWithEmailLink(auth, window.location.href)) return;
        const email = window.localStorage.getItem(emailKey);
        if (!email) {
            alert($t("new_alert"));
            // TODO: handle better
            return;
        }
        try {
            await signInWithEmailLink(auth, email, window.location.href);
            window.localStorage.removeItem(emailKey);
            await goto("/panel");
        } catch (e) {
            console.error(e);
        }
    });
</script>

<div class="w-full max-w-lg">
    {#if !success}
        <div transition:slide class="space-y-4">
            <Label class="mb-4">
                {$t("email")}
                <Input type="text" class="mt-1" bind:value={email}/>
            </Label>

            <Button class="w-full" type="submit" on:click={login}>{$t("continue")}</Button>
            <LoginWithTwitch/>
            <Helper>
                <span class="text-gray-500">{$t("tos_agree")} <button class="hover:underline" on:click={onTosDetailsClicked}>{$t("tos")}.</button></span>
            </Helper>
        </div>
    {:else}
        <div transition:slide>
            <P class="text-justify mb-6">{$t("email_confirmation", { values: {email: email}})}</P>

            <Button class="w-full mb-2" outline on:click={checkEmail}>{$t("email_check")}</Button>
            <Helper>
                <span class="text-gray-500">{$t("email_button_help")}</span>
            </Helper>
        </div>
    {/if}
</div>