<script lang="ts">
    import { Tip } from '$lib/pb/stredono_pb';
    import {onDestroy, onMount} from "svelte";
    import type {WebTip} from "$lib/tips";
    import Alert from "./Alert.svelte";
    import {page} from "$app/stores";

    let ws: WebSocket;

    const newDonation = (pbDonate: Tip) => {
        console.log(pbDonate);
        const donate: WebTip = {
            amount: pbDonate.Amount,
            currency: pbDonate.Currency,
            message: pbDonate.Message,
            user: pbDonate.Sender
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
            const sdReq = Tip.fromBinary(new Uint8Array(event.data))
            newDonation(sdReq);
        };

        ws.onerror = (err) => {
            console.error(err);
        };
    });

    onDestroy(() => {
        ws.close();
    });

    let donations: WebTip[] = [];
</script>

{#if donations.length > 0}
    <Alert donate={donations[0]} />
{/if}

