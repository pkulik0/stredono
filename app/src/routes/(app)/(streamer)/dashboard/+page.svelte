<script lang="ts">
    import { InfoCircleOutline, InfoCircleSolid } from 'flowbite-svelte-icons';
    import Dashboard from './Dashboard.svelte';
    import { userStore } from '$lib/user';
    import { Label, P, Popover, Select } from 'flowbite-svelte';
    import { onMount } from 'svelte';
    import { t } from 'svelte-i18n';
    import { getModeratedUsersListener, moderatedUsersStore } from './moderated';
    import {slide} from 'svelte/transition';
    import { dashboardUidStore } from './store';

    let moderatedUsersUnsub: any;
    $: console.log($moderatedUsersStore)

    onMount(() => {
        const unsub = userStore.subscribe(u => {
            if(!u) {
                if(moderatedUsersUnsub) moderatedUsersUnsub();
                moderatedUsersUnsub = undefined;
                return
            }

            moderatedUsersUnsub = getModeratedUsersListener(u.Uid)
        })

        return () => {
            unsub();
            if(moderatedUsersUnsub) moderatedUsersUnsub();
        }
    })

    const capitalize = (s: string) => s.charAt(0).toUpperCase() + s.slice(1);
    $: items = (() => {
        const user = $userStore;
        if(!user) return [];

        let users = $moderatedUsersStore.map(u => ({ value: u.Uid, name: capitalize(u.Username) })).sort();
        users.unshift({ value: user.Uid, name: capitalize(user.Username) + ` (${$t("me")})` });

        return users;
    })();
    let selectedUser = "";
    $: dashboardUidStore.set(selectedUser);

    $: if(selectedUser === "" && $userStore) selectedUser = $userStore.Uid;
</script>

<div class="flex flex-col space-y-10">
    {#if $moderatedUsersStore.length > 0}
        <div transition:slide>
            <Label>
                <div class="flex items-center">
                    {$t("selected_user")}
                    <InfoCircleOutline class="w-4 h-4 ml-1" />
                    <Popover class="w-80">
                        <P class="text-sm">
                            {$t("selected_user_help")}
                        </P>
                    </Popover>
                </div>
                <Select bind:value={selectedUser} {items}/>
            </Label>
        </div>
    {/if}

    {#if selectedUser}
        <Dashboard uid={selectedUser} />
    {/if}
</div>