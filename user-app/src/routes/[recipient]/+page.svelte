<script lang="ts">
    import axios from 'axios';
    import {PUBLIC_GET_LISTENERS_LINK, PUBLIC_SENT_DONATE_LINK} from "$env/static/public";
    import {DonateStatus, SendDonateRequest, SendDonateResponse} from "../../../../pb/functions_pb";
    import {
        Alert,
        Avatar,
        Button,
        ButtonGroup,
        Card,
        Dropdown,
        DropdownItem,
        Hr,
        Img,
        Input,
        InputAddon,
        Label,
        Textarea,
        TextPlaceholder
    } from "flowbite-svelte";
    import {slide} from "svelte/transition";
    import {page} from "$app/stores";
    import {DotsHorizontalOutline, ExclamationCircleSolid} from "flowbite-svelte-icons";
    import {onMount} from "svelte";
    import {FetchError, getAvatarUrl, getPageDocByUsername} from "$lib/data";
    import {goto} from "$app/navigation";

    const sendDonate = async () => {
        if (isBlocked) return;
        if ((username === "" || email === "") && !isSignedIn) return;

        const amountFloat = parseFloat(amount);
        if (amountFloat <= 1) return;

        const sdReq = new SendDonateRequest({
            amount: amountFloat,
            currency: currency,
            email: email,
            sender: username,
            recipientId: uid,
            message: message,
            status: DonateStatus.INITIATED
        });

        const res = await axios.post(PUBLIC_SENT_DONATE_LINK, sdReq.toBinary(), { responseType: 'arraybuffer' })
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }

        const sdRes = SendDonateResponse.fromBinary(new Uint8Array(res.data));
        window.location.href = sdRes.redirectUrl;
    }

    let recipient = $page.params.recipient;
    let uid:string = "";
    let avatarUrl:string = "";
    let description:string|undefined = undefined;
    let url:string = "";

    let username:string = "";
    let email:string = "";
    let amount:string = "5";
    let message:string = "";
    let currency:string = "PLN";

    let isSignedIn = false;
    $: isBlocked = !isSignedIn && email === "" // TODO: better validation
    $: blockClass = isBlocked ? "pointer-events-none blur-sm" : "";
    $: donateButtonColor = isBlocked ? "alternative" : "primary";

    let hasListeners: boolean|undefined = undefined;

    const getListeners = async () => {
        // TODO: error handling
        const res = await axios.get(PUBLIC_GET_LISTENERS_LINK + "?uid=" + uid)
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }
        hasListeners = res.data !== 0;
    }

    onMount(async () => {
        try {
            const doc = await getPageDocByUsername(recipient);
            const data = doc.data()
            if (!data) {
                await goto("/");
                return;
            }

            uid = data.uid;
            getListeners().then(r => {})
            setInterval(getListeners, 10000)

            console.log("uid", uid)

            avatarUrl = await getAvatarUrl(uid);
            description = data.description;
            url = data.url;
        } catch (e) {
            if(e instanceof FetchError) {
                await goto("/");
                return;
            }
            throw e;
        }
    });

    $: cardClass = uid === "" ? "blur" : "";
</script>


<Card padding="xl" class="{cardClass}">

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

    <div class="text-center mt-4">
        <h2 class="text-lg font-bold">
            You're donating to
            <span class="capitalize font-black text-primary-700">{recipient}</span>
        </h2>
        <p class="text-sm text-gray-500 text-justify">
            {#if description}
                {description}
            {:else}
                <TextPlaceholder />
            {/if}
        </p>
    </div>

    <Img/>

    <div class="space-y-4 mt-6">
        <Label>
            Email
            <Input bind:value={email} />
        </Label>

        {#if isBlocked}
            <div transition:slide>
                Or

                <Button color="purple" size="xs">Login with Twitch</Button>
            </div>
        {/if}
    </div>

    <Hr/>

    <div class="space-y-6 {blockClass}">
        <Label>
            Name
            <Input bind:value={username} />
        </Label>

        <div>
            <Label>
                Amount
            </Label>
            <ButtonGroup class="w-full">
                <Input type="number" bind:value={amount} min="1" />
                <InputAddon>{currency}</InputAddon>
            </ButtonGroup>
        </div>


        <Label>
            Message
            <Textarea bind:value={message} />
        </Label>

        <Button on:click={sendDonate} class="w-full" color="{donateButtonColor}">Donate</Button>

        {#if hasListeners === false}
            <div transition:slide>
                <Alert border>
                    <ExclamationCircleSolid slot="icon"/>
                    <span class="capitalize">{recipient}</span>
                    is currently not displaying the donations on stream.
                </Alert>
            </div>
        {/if}
    </div>
</Card>