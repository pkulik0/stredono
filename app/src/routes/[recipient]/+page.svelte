<script lang="ts">
    import axios from 'axios';
    import {PUBLIC_FUNC_LINK} from "$env/static/public";
    import {DonateStatus, SendDonateRequest, SendDonateResponse} from "$lib/pb/functions_pb";
    import {
        Card,
    } from "flowbite-svelte";
    import {page} from "$app/stores";
    import {onMount} from "svelte";
    import {FetchError, getProfileByUsername} from "$lib/page";
    import {goto} from "$app/navigation";
    import {emailStore, senderStore} from "$lib/stores";
    import Form from "./Form.svelte";
    import ProfileHeader from "$lib/comp/ProfileHeader.svelte";
    import type {Profile} from "$lib/pb/profile_pb";

    const sendDonate = async () => {
        if(!profile) throw new Error("Profile is undefined");
        if ($senderStore === "" || $emailStore === "") return;

        const amountFloat = parseFloat(amount);
        if (amountFloat < profile.minimumAmount) return; // TODO: show error

        const sdReq = new SendDonateRequest({
            amount: amountFloat,
            currency: currency,
            email: $emailStore,
            sender: $senderStore,
            recipientId: profile.uid,
            message: message,
            status: DonateStatus.INITIATED
        });

        const res = await axios.post(PUBLIC_FUNC_LINK + "sendDonate", sdReq.toBinary(), { responseType: 'arraybuffer' })
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }

        const sdRes = SendDonateResponse.fromBinary(new Uint8Array(res.data));
        window.location.href = sdRes.redirectUrl;
    }

    let recipient = $page.params.recipient;
    let profile: Profile|undefined = undefined;

    let amount:string = "5";
    let message:string = "";
    let currency:string = "PLN";

    $: isBlocked = $emailStore === ""
    let hasListeners: boolean = true;

    const getListeners = async () => {
        const res = await axios.get(PUBLIC_FUNC_LINK + "getListeners?uid=" + profile!.uid)
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }
        hasListeners = res.data !== 0;
    }

    onMount(async () => {
        try {
            profile = await getProfileByUsername(recipient);
            amount = profile.minimumAmount.toString();

            getListeners().then(r => {})
            setInterval(getListeners, 10000)
        } catch (e) {
            if(e instanceof FetchError) {
                await goto("/");
                return;
            }
            throw e;
        }
    });
</script>

{#if profile}
    <div class="flex justify-center items-center h-screen">
        <Card padding="xl" size="lg">

            <div class="px-[5%]">
                <ProfileHeader {profile}/>

                <Form {recipient} {sendDonate} {isBlocked} {hasListeners} bind:amount bind:message bind:currency/>
            </div>
        </Card>
    </div>
{/if}