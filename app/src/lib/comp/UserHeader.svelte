<script lang="ts">
    import type { User } from '$lib/pb/stredono_pb';
    import {Avatar, Blockquote, Dropdown, DropdownItem, Heading} from "flowbite-svelte";
    import {DotsHorizontalOutline, ExclamationCircleSolid} from "flowbite-svelte-icons";
    import { t } from 'svelte-i18n';

    export let user: User | undefined
    export let interactive: boolean = true;

    $: pictureUrl = user?.PictureUrl ?? "";
    $: url = user?.Url ?? "";
    $: displayName = user ? (user.DisplayName.length > 0 ? user.DisplayName : user.Username) : "???";
    $: description = user?.Description ?? "Description";

    $: avatarClass = url ? "hover:opacity-75" : "";

    let interactClass = interactive ? "pointer-events-auto" : "pointer-events-none";
</script>

<div class="flex justify-end">
    <DotsHorizontalOutline />
    <Dropdown class="w-36 {interactClass}">
        <DropdownItem class="flex text-red-500">
            <ExclamationCircleSolid class="mr-2" />
            {$t("report")}
        </DropdownItem>
    </Dropdown>
</div>

<div class="flex justify-center">
    <Avatar rounded size="xl" src={pictureUrl} href={url} target="_blank" class="{avatarClass}"/>
</div>

<div class="text-center mt-6 mb-2 space-y-4">
    <Heading tag="h3">
        {$t("profile_header_msg")}
        <span class="capitalize font-black text-primary-700">{displayName}</span>
    </Heading>
    <Blockquote class="text-lg text-gray-500 text-center px-10">
        {description}
    </Blockquote>
</div>