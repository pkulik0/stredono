<script>
    import { Button, Card, Checkbox, Helper, Input, Label, P } from 'flowbite-svelte';
    import {isSignInWithEmailLink, sendSignInLinkToEmail, signInWithEmailLink, signInWithRedirect, getRedirectResult, signInWithCredential, signInWithPopup} from "firebase/auth";
    import {auth} from "$lib/firebase/firebase";
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

    const loginTwitch = async () => {
        const provider = new OAuthProvider("oidc.twitch")
        provider.addScope("user:read:email");
        provider.addScope("moderator:read:followers");
        provider.addScope("channel:read:subscriptions");
        provider.addScope("channel:read:redemptions");
        provider.addScope("bits:read");
        provider.addScope("channel:manage:ads");
        provider.addScope("channel:read:ads");
        provider.addScope("channel:manage:broadcast");
        provider.addScope("channel:edit:commercial");
        provider.addScope("channel:read:hype_train");
        provider.addScope("channel:read:goals");
        provider.addScope("channel:read:vips");
        provider.addScope("user:read:broadcast");
        provider.addScope("user:read:chat");
        await signInWithPopup(auth, provider);
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
            <Button color="purple" outline class="mb-2 w-full" on:click={loginTwitch}>Login with Twitch</Button>
            <Helper>
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