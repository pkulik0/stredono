<script lang="ts">
	import { Img, Listgroup, Popover } from 'flowbite-svelte';
	import {locale, locales} from 'svelte-i18n';

	$: src = `/flags/${$locale}.png`

	$: items = $locales.map(l => ({
		current: l === $locale,
		name: l,
	}))

	const setLanguage = (code: string) => {
		locale.set(code)
	}
</script>

<Img {src} class="h-6 w-auto rounded"/>
<Popover border={false}>
	<Listgroup active border={false} {items} let:item on:click={(e) => setLanguage(e.detail.name)}>
		<Img src={`/flags/${item.name}.png`} class="h-6 w-auto rounded"/>
	</Listgroup>
</Popover>