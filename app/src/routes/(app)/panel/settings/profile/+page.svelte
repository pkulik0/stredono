<script lang="ts">
    import FileDropzone from '$lib/comp/FileDropzone.svelte';
    import type { User } from '$lib/pb/stredono_pb';
    import {Button, Card, Fileupload, Helper, Hr, Input, Label, Textarea,} from "flowbite-svelte";
    import UserHeader from "$lib/comp/UserHeader.svelte";
    import {saveUser, userStore} from "$lib/user";
    import {onMount} from "svelte";
    import {uploadToStorage} from "$lib/ext/firebase/storage";
    import {sendNotification, Notification} from "$lib/notifications";

    let displayName: string = "";
    let url: string = "";
    let description: string = "";
    let pictureFile: File|undefined = undefined;
    let pictureUrl: string = ""

    let user: User|undefined = undefined;

    onMount(() => {
        return userStore.subscribe((u) => {
            user = u;

            if(u) {
                displayName = u.DisplayName;
                url = u.Url;
                description = u.Description;
                pictureUrl = u.PictureUrl;
                return
            }

            displayName = "";
            url = "";
            description = "";
            pictureUrl = "";
        })
    })

    $: if(pictureFile) {
        pictureUrl = URL.createObjectURL(pictureFile);
    }

    $: if(user) {
        user.DisplayName = displayName;
        user.Url = url;
        user.Description = description;
        user.PictureUrl = pictureUrl;
    }

    const clickSave = async () => {
        if(!user) return;

        if (pictureFile) {
            pictureUrl = await uploadToStorage("public", "picture", pictureFile, true);
            user.PictureUrl = pictureUrl;
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
            <Input placeholder={user?.Username || "???"} bind:value={displayName} type="text"/>
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
            Picture
            <FileDropzone description=".png .jpg .jpeg .webp" bind:file={pictureFile} />
        </Label>

        <Button class="w-full" on:click={clickSave}>Save</Button>
    </div>
</div>