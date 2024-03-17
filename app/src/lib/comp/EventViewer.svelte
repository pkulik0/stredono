<script lang="ts">
	import { eventToAlert } from '$lib/alerts';
	import { settingsStore } from '$lib/settings';
	import { type Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
	import { Event, EventType } from '$lib/pb/event_pb';
	import 'animate.css';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { confirmEvent } from '../../routes/overlay/events/listener';
	import { keyStore } from '../../routes/overlay/stores';

	export let isTest: boolean = false;
	export let alerts: Alert[];
	export let events: Event[];

	$: minimumDuration = ($settingsStore?.Events?.MinDisplayTime ?? 5) * 1000;
	let event: Event|undefined;

	let lastChangeTimestamp: number = 0;
	let eventTimeout: any;
	$: if($settingsStore?.Events?.IsPaused !== true) {
		if(events[0] && !event) {
			clearTimeout(eventTimeout);
			eventTimeout = setTimeout(() => {
				event = events[0];
				lastChangeTimestamp = Date.now();
			}, 1000)
		}
	}

	$: eventAlert = event ? eventToAlert(event, alerts) : undefined;

	$: template = (() => {
		switch(event?.Type) {
			case EventType.TIP: return $settingsStore?.Events?.Tip?.Template;
			case EventType.CHEER: return $settingsStore?.Events?.Cheer?.Template;
			case EventType.FOLLOW: return $settingsStore?.Events?.Follow?.Template;
			case EventType.SUB: return $settingsStore?.Events?.Sub?.Template;
			case EventType.SUB_GIFT: return $settingsStore?.Events?.SubGift?.Template;
			case EventType.RAID: return $settingsStore?.Events?.Raid?.Template;
			case EventType.CHAT_TTS: return $settingsStore?.Events?.ChatTTS?.Template;
		}
	})() ?? "";

	let header = "";
	$: if(event && eventAlert && eventAlert.AccentColor) {
		header = template
		Object.entries(event.Data).forEach(([key, value]) => {header = header.replace(`{${key.toLowerCase()}}`, `<span style="color: ${eventAlert.AccentColor};">${value}</span>`);});
		header = header.replace("{user}", `<span style="color: ${eventAlert.AccentColor};">${event.SenderName}</span>`)
	}
	$: message = event ? event.Data.Message : '';

	$: alignment = eventAlert ? eventAlert.Alignment : Alignment.CENTER;
	$: alignmentClass = (() => {
		switch(alignment) {
			case Alignment.START: return 'text-start';
			case Alignment.CENTER: return 'text-center';
			case Alignment.END: return 'text-end';
			case Alignment.JUSTIFY: return 'text-justify';
		}
	})();

	$: textPosition = eventAlert?.TextPosition ?? Position.TOP;
	$: isVertical = textPosition === Position.TOP || textPosition === Position.BOTTOM;
	$: imgClass = isVertical ? 'my-4' : 'mx-4';
	$: containerClass = (() => {
		switch(textPosition) {
			case Position.TOP: return 'flex-col-reverse';
			case Position.BOTTOM: return 'flex-col';
			case Position.LEFT: return 'flex-row-reverse';
			case Position.RIGHT: return 'flex-row';
		}
	})();

	$: animationClass = (() => {
		if(!eventAlert) return;

		let base = "animate__animated animate__infinite animate__slow";
		switch(eventAlert.AnimationSpeed) {
			case Speed.OFF: {
				return "";
			}
			case Speed.SLOW: {
				base += ' animate__slower';
				break;
			}
			case Speed.MEDIUM: {
				base += ' animate__slow';
				break;
			}
			case Speed.FAST: {
				base += ' animate__fast';
				break;
			}
			case Speed.FASTER: {
				base += ' animate__faster';
				break;
			}
		}
		switch(eventAlert.Animation) {
			case AnimationType.PULSE: return base + ' animate__pulse';
			case AnimationType.SHAKE_HORIZONTALLY: return base + ' animate__shakeX';
			case AnimationType.SHAKE_VERTICALLY: return base + ' animate__shakeY';
			case AnimationType.HEART_BEAT: return base + ' animate__heartBeat';
			case AnimationType.JELLO: return base + ' animate__jello';
			case AnimationType.TADA: return base + ' animate__tada';
			case AnimationType.BOUNCE: return base + ' animate__bounce';
		}
	})();

	let soundAudioElement: HTMLAudioElement = new Audio("")
	let ttsAudioElement: HTMLAudioElement = new Audio("")

	$: soundAudioElement.volume = $settingsStore?.Events?.IsMuted ? 0 : 1
	$: ttsAudioElement.volume = $settingsStore?.Events?.IsMuted ? 0 : 1

	let soundTimeout: any
	let lastPlayTimestamp: number = 0;
	$: if(eventAlert?.SoundUrl && lastChangeTimestamp > lastPlayTimestamp) {
		soundAudioElement.src = eventAlert.SoundUrl

		clearTimeout(soundTimeout)
		soundTimeout = setTimeout(() => {
			lastPlayTimestamp = lastChangeTimestamp;
			soundAudioElement.play().catch(() => {
				console.error("Sound failed")
				ttsAudioElement.play().catch(() => {
					console.log("TTS failed (after sound failed)")
					finish()
				})
			})
		}, 250)
	}
	$: ttsAudioElement.src = event?.TTSUrl ?? ""

	$: if(eventAlert?.SoundUrl === "") {
		ttsAudioElement.play().catch(() => {
			console.log("TTS failed (without sound)")
			finish()
		})
	}

	let totalDuration = 0;
	let finishTimeout: any;
	const finish = async () => {
		if(finishTimeout || isTest) return;
		if(!event) return;

		await confirmEvent($keyStore, event.ID)

		const waitTime = Math.max(minimumDuration - totalDuration, 0);
		finishTimeout = setTimeout(() => {
			event = undefined
			finishTimeout = undefined;
			totalDuration = 0;
		}, waitTime)
	}

	onMount(() => {
		soundAudioElement.onended = () => {
			console.log("Sound ended")
			totalDuration += soundAudioElement.duration * 1000
			ttsAudioElement.play().catch((e) => {
				console.log("TTS failed after sound")
				finish()
			})
		}
		ttsAudioElement.onended = () => {
			console.log("TTS ended")
			totalDuration += ttsAudioElement.duration * 1000
			finish()
		}
	})

</script>

{#if eventAlert && event}
	<div transition:fade class="z-75 fixed inset-0 pointer-events-none mt-20 flex flex-col items-center animate__animated animate__jackInTheBox">
		<div class="flex items-center  {alignmentClass} {containerClass} {animationClass}" style="color: {eventAlert.TextColor}">
			{#if eventAlert.GifUrl}
				<img src={eventAlert.GifUrl} alt="" class="rounded-xl shadow-2xl {imgClass} max-w-sm max-h-64" />
			{/if}

			<div class="flex flex-col space-y-2 max-w-lg">
				<h5 class="font-extrabold text-2xl pt-2 {alignment === Alignment.JUSTIFY ? 'text-center' : ''}">{@html header}</h5>
				<p class="text-lg">{message}</p>
			</div>
		</div>
	</div>
{/if}