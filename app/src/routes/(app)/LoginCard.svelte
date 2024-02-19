<script>
    import { Button, Card, Checkbox, Helper, Input, Label, P } from 'flowbite-svelte';
    import {isSignInWithEmailLink, sendSignInLinkToEmail, signInWithEmailLink} from "firebase/auth";
    import {auth} from "$lib/firebase/firebase";
    import {onMount} from "svelte";
    import {userStore} from "$lib/user";
    import {goto} from "$app/navigation";
    import {slide} from 'svelte/transition';

    let email = "";
    let tosAccepted = false;
    const emailKey = "sign-in-email";

    let success = false;

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
            <Label class="mb-2">
                Email
                <Input type="text" class="mt-1" bind:value={email}/>
            </Label>

            <div class="mb-6">
                <Checkbox bind:checked={tosAccepted}>
                <span class="inline">
                    I agree to the
                    <a href="/app/static" class="text-primary-700 dark:text-primary-500 hover:underline">terms and conditions</a>
                </span>
                </Checkbox>
            </div>

            <Button class="mb-2 w-full" type="submit" on:click={login}>Enter</Button>
            <Helper>
                <span class="text-gray-500">By clicking "Enter" you agree to our terms and conditions.</span>
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