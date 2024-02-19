<script lang="ts">
    import FileDropzone from '$lib/comp/FileDropzone.svelte';
    import {Button, Card, Fileupload, Helper, Hr, Input, Label, Textarea,} from "flowbite-svelte";
    import UserHeader from "$lib/comp/UserHeader.svelte";
    import {saveUser, userStore} from "$lib/user";
    import {onMount} from "svelte";
    import {uploadToStorage} from "$lib/firebase/storage";
    import type {User} from "$lib/pb/user_pb";
    import {sendNotification, Notification} from "$lib/notifications";

    let username: string = "";
    let url: string = "";
    let description: string = "";
    let avatarFile: File|undefined = undefined;
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

    $: if(avatarFile) {
        avatarUrl = URL.createObjectURL(avatarFile);
    }

    $: if(user) {
        user.username = username;
        user.url = url;
        user.description = description;
        user.avatarUrl = avatarUrl;
    }

    const clickSave = async () => {
        if(!user) return;

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

        <Label>
            Your Link
            <Input bind:value={url} type="text"/>
            <Helper class="mt-1">
                Clicking on your profile picture will take visitors to this link
            </Helper>
        </Label>

        <Label>
            Avatar
            <FileDropzone description=".png .jpg .jpeg .webp" bind:file={avatarFile} />
        </Label>

        <Button class="w-full" on:click={clickSave}>Save</Button>
    </div>
</div>