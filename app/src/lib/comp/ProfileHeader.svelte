<script lang="ts">
    import {Avatar, Blockquote, Dropdown, DropdownItem, Heading} from "flowbite-svelte";
    import {DotsHorizontalOutline, ExclamationCircleSolid} from "flowbite-svelte-icons";
    import type {Profile} from "$lib/pb/profile_pb";

    export let profile: Profile | undefined

    $: avatarUrl = profile?.avatarUrl ?? "";
    $: url = profile?.url ?? "https://stredono.com";
    $: profileName = profile?.name ?? "Stredono";
    $: description = profile?.description ?? "Description";
</script>

<div class="flex justify-end">
    <DotsHorizontalOutline />
    <Dropdown class="w-36">
        <DropdownItem class="flex text-red-500">
            <ExclamationCircleSolid class="mr-2" />
            Report
        </DropdownItem>
    </Dropdown>
</div>

<div class="flex justify-center">
    <Avatar rounded size="xl" src={avatarUrl} href={url} target="_blank" class="hover:opacity-75"/>
</div>

<div class="text-center mt-6 mb-2 space-y-4">
    <Heading tag="h3">
        You're donating to
        <span class="capitalize font-black text-primary-700">{profileName}</span>
    </Heading>
    <Blockquote class="text-sm text-gray-500 text-justify px-10">
        {description}
    </Blockquote>
</div>