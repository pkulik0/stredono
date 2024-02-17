<script lang="ts">
    import {Button, Card, Checkbox, Helper, Input, Label} from "flowbite-svelte";
    import {isSignInWithEmailLink, signInWithEmailLink, sendSignInLinkToEmail} from "firebase/auth";
    import {auth} from "$lib/firebase/firebase";
    import {onMount} from "svelte";
    import {userStore} from "$lib/user";
    import {goto} from "$app/navigation";
    import {ArrowRightSolid} from "flowbite-svelte-icons";

    let email = "";
    let tosAccepted = false;
    const lsKey = "emailForSignIn";

    const login = async () => {
        tosAccepted = true;
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
        if($userStore) await goto("/panel");

        if (!isSignInWithEmailLink(auth, window.location.href)) return;
        const email = window.localStorage.getItem(lsKey);
        if (!email) {
            alert("Open the link on the same device you tried to sign in on.");
            return;
        }
        try {
            await signInWithEmailLink(auth, email, window.location.href);
            window.localStorage.removeItem(lsKey);
            await goto("/panel");
        } catch (e) {
            console.error(e);
        }
    });
</script>

<div class="flex flex-row h-full">
    <div class="flex-1 hidden md:flex">

    </div>

    <div class="flex-1 flex bg-primary-500 dark:bg-primary-900 justify-center items-center md:p-0 p-4">
        <Card class="space-y-4" size="md" padding="xl">
            <Label>
                Email
                <Input type="text" class="mt-1" bind:value={email}/>
            </Label>

            <Checkbox bind:checked={tosAccepted}>
                <span class="inline">
                    I agree to the
                    <a href="/app/static" class="text-primary-700 dark:text-primary-500 hover:underline">terms and conditions</a>
                </span>
            </Checkbox>

            <Button type="submit" on:click={login}>
                Enter
                <ArrowRightSolid class="ml-1.5 w-6 h-6"/>
            </Button>
            <Helper>
                <span class="text-gray-500">By clicking "Enter" you agree to our terms and conditions.</span>
            </Helper>
        </Card>
    </div>
</div>


