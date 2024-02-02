<script lang="ts">
    import { SendDonateRequest } from  "../../../pb/functions_pb";
    import {onDestroy, onMount} from "svelte";

    let ws: WebSocket;

    onMount(() => {
        ws = new WebSocket("ws://localhost:8081/ws");
        ws.binaryType = "arraybuffer";

        ws.onmessage = (event) => {
            console.log(event)
            const sdReq = SendDonateRequest.fromBinary(new Uint8Array(event.data))
            console.log(sdReq);
        };

        ws.onclose = () => {
            console.log("disconnected");
        };

        ws.onerror = (err) => {
            console.error(err);
        };
    });

    onDestroy(() => {
        ws.close();
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
