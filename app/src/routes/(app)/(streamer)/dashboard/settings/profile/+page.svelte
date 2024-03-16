<script lang="ts">
    import FileDropzone from '$lib/comp/FileDropzone.svelte';
    import type { User } from '$lib/pb/user_pb';
    import { Button, Card, Fileupload, Heading, Helper, Hr, Input, Label, Textarea } from 'flowbite-svelte';
    import UserHeader from "$lib/comp/UserHeader.svelte";
    import {saveUser, userStore} from "$lib/user";
    import {onMount} from "svelte";
    import {uploadToStorage} from "$lib/ext/firebase/storage";
    import {sendNotification, Notification} from "$lib/notifications";
    import {t} from "svelte-i18n";

    let displayName: string = "";
    let url: string = "";
    let description: string = "";
    let pictureFile: File|undefined = undefined;
    let pictureUrl: string = ""

    let user: User|undefined = undefined;

    onMount(() => {
        return userStore.subscribe((u) => {
            user = u||undefined;

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

<Heading tag="h2">{$t("profile")}</Heading>
<div class="flex flex-col w-full items-center justify-center p-4">
    <Heading tag="h4" class="mb-4">{$t("preview")}</Heading>
    <Card padding="xl" size="xl">
        <UserHeader interactive={false} {user}/>
    </Card>

    <Heading tag="h4" class="my-4">{$t("edit")}</Heading>
    <div class="space-y-6 w-full max-w-2xl">
        <Label class="flex-1">
            {$t("display_name")}
            <Input placeholder={user?.Username || "???"} bind:value={displayName} type="text"/>
        </Label>

        <Label>
            {$t("description")}
            <Textarea bind:value={description} />
        </Label>

        <Label>
            {$t("your_url")}
            <Input placeholder="https://stredono.com" bind:value={url} type="text"/>
            <Helper class="mt-1">{$t("your_url_help")}</Helper>
        </Label>

        <Label>
            {$t("picture")}
            <FileDropzone description=".png .jpg .jpeg .webp" bind:file={pictureFile} />
        </Label>

        <Button on:click={clickSave}>{$t("save")}</Button>
    </div>
</div>