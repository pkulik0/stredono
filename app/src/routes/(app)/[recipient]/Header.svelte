<script lang="ts">
    import {Avatar, Blockquote, Dropdown, DropdownItem, Heading, TextPlaceholder} from "flowbite-svelte";
    import {DotsHorizontalOutline, ExclamationCircleSolid} from "flowbite-svelte-icons";

    export let recipient: string;
    export let description: string;
    export let avatarUrl: string;
    export let url: string;
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
    <Avatar rounded size="xl" src={avatarUrl} href={url} target="_blank" />
</div>

<div class="text-center mt-4 space-y-2">
    <Heading tag="h3">
        You're donating to
        <span class="capitalize font-black text-primary-700">{recipient}</span>
    </Heading>
    <Blockquote class="text-sm text-gray-500 text-justify px-10">
        {#if description}
            {description}
        {:else}
            <TextPlaceholder />
        {/if}
    </Blockquote>
</div>