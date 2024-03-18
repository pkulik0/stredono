<script lang="ts">
    import FileDropzone from '$lib/comp/FileDropzone.svelte';
    import UsernameEdit from '$lib/comp/UsernameEdit.svelte';
    import type { User } from '$lib/pb/user_pb';
    import { Button, Card, Heading, Helper, Input, Label, Textarea } from 'flowbite-svelte';
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
                pictureUrl = u.PictureUrl;
                return
            }

            pictureUrl = "";
        })
    })

    $: if(pictureFile) {
        pictureUrl = URL.createObjectURL(pictureFile);
    }

    $: if(user) {
        user.PictureUrl = pictureUrl;
    }

    const clickSave = async () => {
        if(!user) return;

        if (pictureFile) {
            pictureUrl = await uploadToStorage("public", "picture", pictureFile, true);
            user.PictureUrl = pictureUrl;
        }

        await saveUser(user)
        sendNotification(new Notification($t("profile_saved")))
    }

    let isValid = true;
</script>

{#if $userStore}
    <Heading tag="h2">{$t("profile")}</Heading>
    <div class="flex flex-col w-full items-center justify-center p-4">
        <Heading tag="h4" class="mb-4">{$t("preview")}</Heading>
        <Card padding="xl" size="xl">
            <UserHeader interactive={false} {user}/>
        </Card>

        <Heading tag="h4" class="my-4">{$t("edit")}</Heading>
        <div class="space-y-6 w-full max-w-2xl">
<!--            <UsernameEdit bind:username={$userStore.Username} bind:isValid/>-->

            <Label>
                {$t("description")}
                <Textarea placeholder={$t("want_visitors_to_see")} bind:value={$userStore.Description} />
            </Label>

            <Label>
                {$t("your_url")}
                <Input placeholder="https://stredono.com" bind:value={$userStore.Url} type="text"/>
                <Helper class="mt-1">{$t("your_url_help")}</Helper>
            </Label>

            <Label>
                {$t("picture")}
                <FileDropzone description=".png .jpg .jpeg .webp" bind:file={pictureFile} />
            </Label>

            <Button on:click={clickSave}>{$t("save")}</Button>
        </div>
    </div>
{/if}