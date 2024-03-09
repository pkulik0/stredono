<script lang="ts">
	import { auth } from '$lib/ext/firebase/firebase';
	import { OAuthProvider, signInWithPopup } from 'firebase/auth';
	import { Button } from 'flowbite-svelte';
	import {t} from 'svelte-i18n';

	const loginTwitch = async () => {
		const provider = new OAuthProvider("oidc.twitch")

		provider.addScope("user:read:email");
		provider.addScope("channel:bot");
		provider.addScope("channel:moderate");
		provider.addScope("channel:manage:moderators");
		provider.addScope("moderation:read");
		provider.addScope("moderator:read:followers");
		provider.addScope("channel:read:subscriptions");
		provider.addScope("channel:read:redemptions");
		provider.addScope("channel:manage:redemptions");
		provider.addScope("bits:read");
		provider.addScope("channel:manage:ads");
		provider.addScope("channel:read:ads");
		provider.addScope("channel:manage:broadcast");
		provider.addScope("channel:edit:commercial");
		provider.addScope("channel:read:hype_train");
		provider.addScope("channel:read:goals");
		provider.addScope("channel:read:vips");
		provider.addScope("user:read:broadcast");
		provider.addScope("user:read:chat");

		await signInWithPopup(auth, provider);
	}
</script>

<Button color="purple" outline class="w-full" on:click={loginTwitch}>{$t("continue_with_twitch")}</Button>