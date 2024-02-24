<script lang="ts">
    import { Currency, Tip, TipStatus, User } from '$lib/pb/stredono_pb';
    import { terraformOutput } from '$lib/terraform_output';
    import axios from 'axios';
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

    const sendDonate = async () => {
        if(!user) throw new Error("User is undefined");
        if ($senderStore === "" || $emailStore === "") return;

        const amountFloat = parseFloat(amount);
        if (amountFloat < user.MinimumAmount) return; // TODO: show error

        const sdReq = new Tip({
            Amount: amountFloat,
            Currency: currency,
            Email: $emailStore,
            Sender: $senderStore,
            RecipientId: user.Uid,
            Message: message,
            Status: TipStatus.INITIATED
        });

        const res = await axios.post(terraformOutput.FunctionUrls.TipSend, sdReq.toBinary(), { responseType: 'arraybuffer' })
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }
        window.location.href = res.data;
    }

    console.log(terraformOutput.FunctionUrls.TipSend)

    let recipient = $page.params.recipient;
    let user: User|undefined = undefined;

    let amount:string = "5";
    let message:string = "";
    let currency:Currency = Currency.UNKNOWN;

    $: isBlocked = $emailStore === ""
    let hasListeners: boolean = true;

    // const getListeners = async () => {
    //     if(!user) return;
    //
    //     const res = await axios.get(terraformOutput.FunctionUrls.GetListeners + "?uid=" + user.Uid)
    //     if(res.status !== 200) {
    //         console.error(res.data); // TODO: handle error
    //         return;
    //     }
    //     hasListeners = res.data !== 0;
    // }

    onMount(async () => {
        try {
            user = await getUserByUsername(recipient);
            if(!user) return;

            amount = user.MinimumAmount.toString();
            currency = user.Currency;
        } catch (e) {
            console.error(e);
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

                <Form {recipient} {sendDonate} {isBlocked} {hasListeners} bind:amount bind:message {currency}/>
            </div>
        </Card>
    </div>
{/if}