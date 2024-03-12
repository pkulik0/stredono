<script lang="ts">
	import { eventToAlert } from '$lib/alerts';
	import { settingsStore } from '$lib/events_settings';
	import { type Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
	import { Event, EventType } from '$lib/pb/event_pb';
	import 'animate.css';
	import { onMount } from 'svelte';
	import {fade} from 'svelte/transition';
	import { confirmEvent } from '../../routes/overlay/events/listener';
	import { keyStore } from '../../routes/overlay/stores';

	export let isTest: boolean = false;
	export let minimumDuration: number = 5000;

	export let alerts: Alert[];

	export let events: Event[];

	let event: Event|undefined;

	$: if(events[0] && !event) {
		setTimeout(() => {
			event = events[0];
		}, 1000)
	}


	$: eventAlert = event ? eventToAlert(event, alerts) : undefined;

	$: eventTypeName = event ? JSON.parse(JSON.stringify(EventType))[event.Type] : '';

	$: header = $settingsStore?.Events?.Event[eventTypeName]?.MessageTemplate ?? '';
	$: if(event && eventAlert) {
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

	$: soundAudioElement.src = eventAlert?.SoundUrl ?? ""
	$: ttsAudioElement.src = event?.TTSUrl ?? ""

	$: if(eventAlert?.SoundUrl === "") {
		ttsAudioElement.play().catch((e) => {
			finish()
		})
	}

	let totalDuration = 0;
	let timeout: any;

	const finish = () => {
		if(timeout || isTest) return;

		timeout = setTimeout(async () => {
			await confirmEvent($keyStore, event.ID)
			event = undefined
			timeout = undefined
		}, Math.max(minimumDuration - totalDuration, 0))
	}

	onMount(() => {
		soundAudioElement.autoplay = true

		soundAudioElement.onended = () => {
			totalDuration += soundAudioElement.duration * 1000
			ttsAudioElement.play().catch((e) => {
				finish()
			})
		}
		ttsAudioElement.onended = () => {
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