<script>
    import {Button, ButtonGroup, FloatingLabelInput, Input, InputAddon, Label} from "flowbite-svelte";
    import {EnvelopeSolid, EyeOutline, EyeSlashOutline} from "flowbite-svelte-icons";

    let showPassword = false;

    let email = "";
    let password = "";

    const login = () => {
        alert(`Email: ${email}, Password: ${password}`);
    }

    const loginWithTwitch = () => {
        alert("Login with Twitch");
    }
</script>

<div class="flex flex-col space-y-4">
    <FloatingLabelInput style="filled" id="inputEmail" type="text" class="mt-4" bind:value={email}>
        Email
    </FloatingLabelInput>

    <FloatingLabelInput style="filled" id="inputPassword" type="password" class="mt-4" bind:value={password}>
        Password
    </FloatingLabelInput>

    <Button type="submit" color="primary" size="lg" on:click={login}>Login</Button>
    <Button type="submit" color="purple" size="lg" on:click={loginWithTwitch}>Login with Twitch</Button>
</div>

