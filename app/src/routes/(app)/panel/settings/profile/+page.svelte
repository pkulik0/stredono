<script lang="ts">
    import {Button, Card, Fileupload, Helper, Hr, Input, Label, Textarea,} from "flowbite-svelte";
    import Header from "$lib/comp/ProfileHeader.svelte";
    import type {Profile} from "$lib/pb/profile_pb";
    import {onMount} from "svelte";
    import {getProfileByUsername, saveProfileToDb} from "$lib/page";
    import {getTwitchAuthUrl} from "$lib/twitch";
    import {uploadToStorage} from "$lib/firebase";
    import {usernameStore} from "$lib/stores";

    let profile: Profile|undefined = undefined;

    onMount(() => {
        return usernameStore.subscribe(async (username) => {
            if (!username) {
                profile = undefined;
                return
            }

            profile = await getProfileByUsername(username);
            name = profile.name;
            description = profile.description;
            url = profile.url;
        });
    });

    let files: FileList|undefined = undefined;
    let name: string = "";
    let url: string = "";
    let description: string = "";
    let avatarUrl: string = ""

    $: if(files) {
        const file = files.item(0);
        if (file) avatarUrl = URL.createObjectURL(file);
    }

    $: if (profile) {
        profile.name = name;
        profile.description = description;
        profile.url = url;
        if (avatarUrl) profile.avatarUrl = avatarUrl;
    }

    const clickSave = async () => {
        if(!profile) return;
        const username = $usernameStore;
        if (!username) return;


        const avatarFile = files?.item(0);
        if (avatarFile) {
            profile.avatarUrl = await uploadToStorage("page", "avatar", avatarFile, true);
        }
        await saveProfileToDb(username, profile);
    }
</script>

<div class="flex flex-col justify-center items-center">
    <Card padding="xl" size="lg">
        <Header {profile}/>
    </Card>

    <div class="space-y-6 mt-10 w-full max-w-xl">

        <Label class="flex-1">
            Display Name
            <Input bind:value={name} type="text"/>
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