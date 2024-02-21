<script lang="ts">
    import axios from 'axios';
    import TerraformOutput from "$lib/terraform_output.json";
    import {DonateStatus, SendDonateRequest, SendDonateResponse} from "$lib/pb/functions_pb";
    import {
        Card,
    } from "flowbite-svelte";
    import {page} from "$app/stores";
    import {onMount} from "svelte";
    import {FetchError, getUserByUsername} from "$lib/user";
    import {goto} from "$app/navigation";
    import {emailStore, senderStore} from "$lib/stores";
    import Form from "./Form.svelte";
    import UserHeader from "$lib/comp/UserHeader.svelte";
    import type {User} from "$lib/pb/user_pb";

    const sendDonate = async () => {
        if(!user) throw new Error("User is undefined");
        if ($senderStore === "" || $emailStore === "") return;

        const amountFloat = parseFloat(amount);
        if (amountFloat < user.MinimumAmount) return; // TODO: show error

        const sdReq = new SendDonateRequest({
            Amount: amountFloat,
            Currency: currency,
            Email: $emailStore,
            Sender: $senderStore,
            RecipientId: user.Uid,
            Message: message,
            Status: DonateStatus.INITIATED
        });

        const res = await axios.post(TerraformOutput.FunctionUrls.SendTip, sdReq.toBinary(), { responseType: 'arraybuffer' })
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }

        const sdRes = SendDonateResponse.fromBinary(new Uint8Array(res.data));
        window.location.href = sdRes.RedirectUrl;
    }

    let recipient = $page.params.recipient;
    let user: User|undefined = undefined;

    let amount:string = "5";
    let message:string = "";
    let currency:string = "PLN";

    $: isBlocked = $emailStore === ""
    let hasListeners: boolean = true;

    const getListeners = async () => {
        if(!user) return;

        const res = await axios.get(TerraformOutput.FunctionUrls.GetListeners + "?uid=" + user.Uid)
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }
        hasListeners = res.data !== 0;
    }

    onMount(async () => {
        try {
            user = await getUserByUsername(recipient);
            amount = user.MinimumAmount.toString();

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

{#if user}
    <div class="flex justify-center items-center h-screen">
        <Card padding="xl" size="lg">

            <div class="px-[5%]">
                <UserHeader {user}/>

                <Form {recipient} {sendDonate} {isBlocked} {hasListeners} bind:amount bind:message bind:currency/>
            </div>
        </Card>
    </div>
{/if}