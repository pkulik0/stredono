<script>
import SideBar from "./SideBar.svelte";
</script>

<div class="flex">
    <div class="min-w-xl p-4">
        <SideBar/>
    </div>

    <div class="flex-1 p-4">
        <slot/>
    </div>
</div>