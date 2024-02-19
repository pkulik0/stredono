<script lang="ts">
    import ParticlesBackground from '$lib/comp/ParticlesBackground.svelte';
    import LoginCard from './LoginCard.svelte';
    import { Heading, Mark, P, Span } from 'flowbite-svelte';
    import {loopOnce} from 'svelte-typewriter'
</script>

<div class="flex h-full w-full">
    <div class="flex flex-col justify-center items-center h-full w-full">

        <div class="max-w-2xl text-justify p-4 mb-6">
            <Heading tag="h1" class="mb-4">Join Stredono  <Span highlight>right now!</Span></Heading>
            <span use:loopOnce={{ interval: 150, delay: 500 }} class="h-16 inline-block">
                <Heading tag="h2" class="mb-4 text-gray-700"><Span gradient gradientClass="">Profit can't wait.</Span></Heading>
            </span>
        </div>

        <LoginCard/>
    </div>
</div>

<ParticlesBackground/>