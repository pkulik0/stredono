<script>
import SideBar from "./SideBar.svelte";
</script>

<div class="flex">
    <SideBar/>

    <div class="flex-1 p-4">
        <slot/>
    </div>
</div>