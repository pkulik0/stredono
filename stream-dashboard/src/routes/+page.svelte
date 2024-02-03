<script lang="ts">
    import {Button, FloatingLabelInput} from "flowbite-svelte";
    import {isSignInWithEmailLink, signInWithEmailLink, sendSignInLinkToEmail} from "firebase/auth";
    import {auth} from "$lib/firebase";
    import {onMount} from "svelte";
    import {userStore} from "$lib/userStore";
    import {goto} from "$app/navigation";

    let email = "";
    const lsKey = "emailForSignIn";

    const login = async () => {
        try {
            const actionCodeSettings = {
                url: window.location.href,
                handleCodeInApp: true,
            };
            await sendSignInLinkToEmail(auth, email, actionCodeSettings);
            window.localStorage.setItem(lsKey, email);
            alert("An email has been sent to you. Please use the link in the email to sign in.");
        } catch (e) {
            console.error(e);
        }
    }

    onMount(async () => {
        if (!isSignInWithEmailLink(auth, window.location.href)) return;
        const email = window.localStorage.getItem(lsKey);
        if (!email) {
            alert("Open the link on the same device you tried to sign in on.");
            return;
        }
        try {
            const cred = await signInWithEmailLink(auth, email, window.location.href);
            window.localStorage.removeItem(lsKey);
            userStore.set(cred.user);
            await goto("/dashboard");
        } catch (e) {
            console.error(e);
        }
    });
</script>

<div class="flex flex-col space-y-4">
    <FloatingLabelInput style="filled" id="inputEmail" type="text" class="mt-4" bind:value={email}>
        Email
    </FloatingLabelInput>

    <Button type="submit" color="primary" size="lg" on:click={login}>Login</Button>
</div>

