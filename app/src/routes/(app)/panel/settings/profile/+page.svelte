<script lang="ts">
    import {Button, Card, Fileupload, Helper, Hr, Input, Label, Textarea,} from "flowbite-svelte";
    import UserHeader from "$lib/comp/UserHeader.svelte";
    import {saveUser, userStore} from "$lib/user";
    import {onMount} from "svelte";
    import {uploadToStorage} from "$lib/firebase/storage";
    import type {User} from "$lib/pb/user_pb";
    import {sendNotification, Notification} from "$lib/notifications";

    let files: FileList|undefined = undefined;
    let username: string = "";
    let url: string = "";
    let description: string = "";
    let avatarUrl: string = ""

    let user: User|undefined = undefined;

    onMount(() => {
        return userStore.subscribe((u) => {
            user = u;

            if(u) {
                username = u.username;
                url = u.url;
                description = u.description;
                avatarUrl = u.avatarUrl;
                return
            }

            username = "";
            url = "";
            description = "";
            avatarUrl = "";
        })
    })

    $: if(files) {
        const file = files.item(0);
        avatarUrl = file ? URL.createObjectURL(file) : "";
    }

    $: if(user) {
        user.username = username;
        user.url = url;
        user.description = description;
        user.avatarUrl = avatarUrl;
    }

    const clickSave = async () => {
        if(!user) return;

        const avatarFile = files?.item(0);
        if (avatarFile) {
            user.avatarUrl = await uploadToStorage("public", "avatar", avatarFile, true);
        }

        await saveUser(user)
        sendNotification(new Notification("Profile saved"))
    }
</script>

<div class="flex flex-col justify-center items-center">
    <Card padding="xl" size="lg">
        <UserHeader interactive={false} {user}/>
    </Card>

    <div class="space-y-6 mt-10 w-full max-w-xl">

        <Label class="flex-1">
            Display Name
            <Input bind:value={username} type="text"/>
        </Label>

        <Label>
            Description
            <Textarea bind:value={description} />
        </Label>

        <Hr/>

        <Label>
            Avatar
            <Fileupload bind:files={files} accept="image/*" />
        </Label>

        <Label>
            Your Link
            <Input bind:value={url} type="text"/>
            <Helper>
                Allowed domains: twitch.tv
            </Helper>
        </Label>

        <Button class="w-full" on:click={clickSave}>Save</Button>
    </div>
</div>