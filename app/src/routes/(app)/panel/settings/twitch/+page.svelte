<script lang="ts">
    import {
        Accordion,
        AccordionItem,
        Button,
        Checkbox,
        Heading,
        Helper, Hr,
        Input,
        Label,
        Range,
        Select
    } from "flowbite-svelte";
    import {getTwitchAuthUrl} from "$lib/ext/twitch";
    import CheckboxDropdown from "$lib/comp/CheckboxDropdown.svelte";
    import axios from "axios";
    import {auth} from "$lib/firebase/firebase";

    const connectTwitch = async () => {
        window.location.href = await getTwitchAuthUrl();
    }

    const getTwitchData = async () => {
        const user = auth.currentUser
        if (!user) return;
        const token = await user.getIdToken(true);

        //TODO: get twitch data
        // const res = await axios.get(PUBLIC_FUNC_LINK + "getTwitchData", {
        //     headers: {
        //         "Authorization": "Bearer " + token
        //     }
        // });
        // if (res.status === 200) {
        //     console.log(res.data);
        // }
    }

    let follows: boolean = false;
    let subscriptions: boolean = true;
    let bits: boolean = true;
    let redemptions: boolean = false;

    let minBits: number = 100;
    let minMonths: number = 1;

    let checkboxClass = "text-md";

    let selectItems = [
        { value: 0, name: 'Always' },
        { value: 1, name: 'One per minute' },
        { value: 2, name: 'Every 5 minutes' },
        { value: 3, name: 'Every 15 minutes' },
        { value: 4, name: 'Every 30 minutes' },
    ];
    let selectValue = 0;
</script>

<Button size="xl" color="purple" class="max-w-xl" on:click={connectTwitch}>Connect Twitch</Button>

<Button size="xl" color="red" class="max-w-xl" on:click={getTwitchData}>Get Data</Button>


<Accordion class="mt-4">
    <AccordionItem open>
        <span slot="header">Events</span>
        <div class="space-y-4">
            <CheckboxDropdown label="Follows" bind:checked={follows}>
                <Label>
                    Frequency
                    <Select items="{selectItems}" bind:value={selectValue} class="max-w-xl" />
                    <Helper>How often you want to be notified about this event</Helper>
                </Label>
            </CheckboxDropdown>

            <CheckboxDropdown label="Subscriptions" bind:checked={subscriptions}>
                <Label>
                    Minimum months
                    <Input bind:value={minMonths} type="number" min="1" class="max-w-xl" />
                    <Helper>You will only be notified if a user has been subscribed for more than this amount of months</Helper>
                </Label>
            </CheckboxDropdown>

            <CheckboxDropdown label="Bits" bind:checked={bits}>
                <Label>
                    Minimum bits
                    <Input bind:value={minBits} type="number" min="1" class="max-w-xl" />
                    <Helper>You will only be notified if a user sends more than this amount of bits</Helper>
                </Label>
            </CheckboxDropdown>

            <Checkbox class="{checkboxClass}" bind:checked={redemptions}>Redemptions</Checkbox>
        </div>
    </AccordionItem>
</Accordion>