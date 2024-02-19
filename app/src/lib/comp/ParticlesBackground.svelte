<script lang="ts">
	import { onMount } from 'svelte'
	import { particlesConfig } from '$lib/particles';

	export let intensity: number = 100

	let tsParticlesInstance: any

	onMount(async () => {
		tsParticlesInstance = await tsParticles.load({
			id: "tsParticles",
			options: particlesConfig(intensity),
		})
	})
</script>

<div id="tsParticles"></div>

<style>
    #tsParticles {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        z-index: -1;
    }
</style>