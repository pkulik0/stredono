<script>
    import LoginWithTwitch from '$lib/comp/LoginWithTwitch.svelte';
    import { Button, Checkbox, Helper, Input, Label, P } from 'flowbite-svelte';
    import {isSignInWithEmailLink, sendSignInLinkToEmail, signInWithEmailLink, signInWithPopup} from "firebase/auth";
    import {auth} from "$lib/ext/firebase/firebase";
    import {onMount} from "svelte";
    import {userStore} from "$lib/user";
    import {goto} from "$app/navigation";
    import {slide} from 'svelte/transition';
    import {OAuthProvider} from "firebase/auth";

    let email = "";
    let tosAccepted = false;
    const emailKey = "sign-in-email";

    let success = false;

    const onTosDetailsClicked = () => {
        console.log("tos details clicked");
    }

    const login = async () => {
        tosAccepted = true;
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
            alert("Open the link on the same device you tried to sign in on.");
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
        <div transition:slide>
            <Label class="mb-4">
                Email
                <Input type="text" class="mt-1" bind:value={email}/>
            </Label>

            <div class="mb-10">
                <Checkbox bind:checked={tosAccepted}>
                <span class="inline">
                    I agree to the
                    <button on:click={onTosDetailsClicked} class="text-primary-700 dark:text-primary-500 hover:underline">terms and conditions</button>
                </span>
                </Checkbox>
            </div>

            <Button class="mb-4 w-full" type="submit" on:click={login}>Enter</Button>
            <LoginWithTwitch/>
            <Helper class="mt-4">
                <span class="text-gray-500">By continuing you agree to our <button class="hover:underline" on:click={onTosDetailsClicked}>terms and conditions.</button></span>
            </Helper>
        </div>
    {:else}
        <div transition:slide>
            <P class="text-justify mb-6">An email has been sent to {email}. Please click the link in the email to sign in.</P>

            <Button class="w-full mb-2" outline on:click={checkEmail}>Check email</Button>
            <Helper>
                <span class="text-gray-500">This will send you to your email provider.</span>
            </Helper>
        </div>
    {/if}
</div>