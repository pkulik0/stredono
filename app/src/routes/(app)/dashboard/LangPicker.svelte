<script lang="ts">
	import { Img, Listgroup, Popover } from 'flowbite-svelte';
	import {locale, locales} from 'svelte-i18n';

	$: items = $locales.map(l => ({
		current: l === $locale, // Although filtered lets make it idiot-proof
		name: l,
	})).filter(l => l.name !== $locale)

	const setLanguage = (code: string) => {
		locale.set(code)
	}

	const imgClass = "h-6 w-auto rounded";
</script>

<Img src={`/flags/${$locale}.png`} class={imgClass}/>
<Popover border={false}>
	<Listgroup active border={false} {items} let:item on:click={(e) => setLanguage(e.detail.name)}>
		<Img src={`/flags/${item.name}.png`} class={imgClass}/>
	</Listgroup>
</Popover>