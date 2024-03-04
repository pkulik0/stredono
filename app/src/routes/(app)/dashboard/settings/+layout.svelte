<script lang="ts">
    import SideBar from "./SideBar.svelte";
</script>

<div class="flex p-4">
    <div class="hidden md:pe-4 md:flex">
        <SideBar/>
    </div>

    <div class="flex flex-col mx-auto max-w-3xl w-full items-center">
        <slot/>
    </div>
</div>