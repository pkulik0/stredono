<script lang="ts">
    import { auth } from '$lib/ext/firebase/firebase';
    import { terraformOutput } from '$lib/constants';
    import { Currency } from '$lib/pb/enums_pb';
    import { Tip, TipStatus } from '$lib/pb/tip_pb';
    import type { User } from '$lib/pb/user_pb';
    import axios from 'axios';
    import { onAuthStateChanged } from 'firebase/auth';
    import {
        Card,
    } from "flowbite-svelte";
    import {page} from "$app/stores";
    import {onMount} from "svelte";
    import {getUserByUsername} from "$lib/user";
    import {goto} from "$app/navigation";
    import {emailStore, senderStore} from "./stores";
    import TipForm from "./TipForm.svelte";
    import UserHeader from "$lib/comp/UserHeader.svelte";

    const sendDonate = async () => {
        if(!user) throw new Error("User is undefined");
        if ($senderStore === "" || $emailStore === "") return;

        const amountFloat = parseFloat(amount);
        // if (amountFloat < user.MinAmount) return; // TODO: show error

        const sdReq = new Tip({
            Amount: amountFloat,
            Currency: currency,
            Email: $emailStore,
            DisplayName: $senderStore,
            RecipientId: user.Uid,
            Message: message,
            Status: TipStatus.INITIATED
        });

        const res = await axios.post(terraformOutput.FunctionsUrl + "/TipSend", sdReq.toBinary(), { responseType: 'arraybuffer' })
        if(res.status !== 200) {
            console.error(res.data); // TODO: handle error
            return;
        }
        window.location.href = res.data;
    }

    let recipient = $page.params.recipient;
    let user: User|undefined = undefined;

    let amount:string = "5";
    let message:string = "";
    let currency:Currency = Currency.UNKNOWN;

    let hasListeners: boolean = true;

    const fetchData = async () => {
        user = await getUserByUsername(recipient);
        if(!user) return;

        // amount = user.MinAmount.toString();
        // currency = user.Currency;
    }

    onMount(() => {
        fetchData()

        return onAuthStateChanged(auth, (user) => {
            if(!user) return;

            user.isAnonymous
        })
    });
</script>

{#if user}
    <div class="flex justify-center items-center h-screen">
        <Card padding="xl" size="lg">

            <div class="px-[5%]">
                <UserHeader {user}/>

                <TipForm {recipient} {sendDonate} {hasListeners} bind:amount bind:message {currency}/>
            </div>
        </Card>
    </div>
{/if}