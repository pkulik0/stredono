<script lang="ts">
    import {Img} from "flowbite-svelte";

    export let src: string;
    export let alt: string = "";
    export let color: string = "primary";
</script>

<div class="relative rounded-xl overflow-hidden brightness-100 dark:brightness-50">
    <Img {src} {alt} />
    <div class="absolute inset-0 bg-{color}-100 dark:bg-{color}-700 opacity-65"></div>
</div>