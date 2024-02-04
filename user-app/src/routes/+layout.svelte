<script>
	import '../app.pcss';
    import {DarkMode} from "flowbite-svelte";
</script>

<div class="h-screen flex items-center justify-center">
    <slot />
</div>

<div id="darkMode" class="p-4 fixed top-0 right-0">
    <DarkMode />
</div>