<script lang="ts">
    import {Avatar, Blockquote, Dropdown, DropdownItem, Heading} from "flowbite-svelte";
    import {DotsHorizontalOutline, ExclamationCircleSolid} from "flowbite-svelte-icons";
    import type {User} from "$lib/pb/user_pb";

    export let user: User | undefined
    export let interactive: boolean = true;

    $: avatarUrl = user?.avatarUrl ?? "";
    $: url = user?.url ?? "";
    $: profileName = user?.username ?? "Stredono";
    $: description = user?.description ?? "Description";

    $: avatarClass = url ? "hover:opacity-75" : "";

    let interactClass = interactive ? "pointer-events-auto" : "pointer-events-none";
</script>

<div class="flex justify-end">
    <DotsHorizontalOutline />
    <Dropdown class="w-36 {interactClass}">
        <DropdownItem class="flex text-red-500">
            <ExclamationCircleSolid class="mr-2" />
            Report
        </DropdownItem>
    </Dropdown>
</div>

<div class="flex justify-center">
    <Avatar rounded size="xl" src={avatarUrl} href={url} target="_blank" class="{avatarClass}"/>
</div>

<div class="text-center mt-6 mb-2 space-y-4">
    <Heading tag="h3">
        You're donating to
        <span class="capitalize font-black text-primary-700">{profileName}</span>
    </Heading>
    <Blockquote class="text-lg text-gray-500 text-center px-10">
        {description}
    </Blockquote>
</div>