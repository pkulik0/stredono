<script lang="ts">
    import '../app.pcss';
		import { onMount } from 'svelte';
		import { init, register } from 'svelte-i18n';

		const setupLocale = () => {
			console.log('Setting up locale');
			register('en', () => import('$lib/i18n/en.json'));

			init({
				fallbackLocale: 'en',
				initialLocale: getLocaleFromNavigator()
			});

			function getLocaleFromNavigator() {
				return 'en';
			}
		}

		onMount(setupLocale);
</script>

<slot/>