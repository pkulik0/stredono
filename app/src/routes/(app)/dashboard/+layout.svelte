<script lang="ts">
    import ParticlesBackground from '$lib/comp/ParticlesBackground.svelte';
    import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
    import {page} from "$app/stores";
    import AlertBox from './AlertBox.svelte';
    import NavBar from './NavBar.svelte';

    $: currentPage = $page.url;
    $: pagesOnPath = currentPage.pathname.split("/").filter(Boolean);
</script>

<NavBar />

<div class="flex justify-center p-4">
    <div class="w-full max-w-7xl space-y-4" >
        <Breadcrumb>
            {#each pagesOnPath as page, i}
                <BreadcrumbItem active={i === pagesOnPath.length - 1} home={i === 0} href={'/' + pagesOnPath.slice(0, i+1).join('/')}>
                    {page[0].toUpperCase() + page.slice(1)}
                </BreadcrumbItem>
            {/each}
        </Breadcrumb>

        <AlertBox/>

        <slot />
    </div>
</div>

<ParticlesBackground intensity={50}/>