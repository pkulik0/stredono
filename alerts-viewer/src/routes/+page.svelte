<script lang="ts">
    const ws = new WebSocket("ws://localhost:8081/ws");

    ws.onopen = () => {
        console.log("connected");
        ws.send("Hello");
    };

    ws.onmessage = (event) => {
        console.log(event.data);
    };

    ws.onclose = () => {
        console.log("disconnected");
    };

    ws.onerror = (err) => {
        console.error(err);
    };
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
