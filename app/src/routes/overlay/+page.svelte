<script lang="ts">
    import { SendDonateRequest } from "$lib/pb/functions_pb";
    import {onDestroy, onMount} from "svelte";
    import type {Donate} from "$lib/donations";
    import Alert from "./Alert.svelte";
    import {page} from "$app/stores";

    let ws: WebSocket;

    const newDonation = (pbDonate: SendDonateRequest) => {
        console.log(pbDonate);
        const donate: Donate = {
            amount: pbDonate.amount,
            currency: pbDonate.currency,
            message: pbDonate.message,
            user: pbDonate.sender
        };
        donations = [donate, ...donations];

        setTimeout(() => {
            donations = donations.slice(0, donations.length - 1);
        }, 10000);
    };

    onMount(() => {
        ws = new WebSocket("ws://localhost:8081/ws?uid=" + $page.url.searchParams.get("uid"));
        ws.binaryType = "arraybuffer";

        ws.onmessage = (event) => {
            const sdReq = SendDonateRequest.fromBinary(new Uint8Array(event.data))
            newDonation(sdReq);
        };

        ws.onerror = (err) => {
            console.error(err);
        };
    });

    onDestroy(() => {
        ws.close();
    });

    let donations: Donate[] = [];
</script>

{#if donations.length > 0}
    <Alert donate={donations[0]} />
{/if}

