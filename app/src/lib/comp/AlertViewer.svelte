<script lang="ts">
	import { eventToAlert } from '$lib/alerts';
	import { settingsStore } from '$lib/events_settings';
	import { type Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
	import { Event, EventType } from '$lib/pb/event_pb';
	import 'animate.css';
	import {fade} from 'svelte/transition';

	export let alerts: Alert[];

	export let visible: boolean = true;
	export let event: Event|undefined;
	export let isTest: boolean = false;
	export let minimumDuration: number = 5;

	export let onShown: (event: Event) => void;

	$: eventAlert = event ? eventToAlert(event, alerts, isTest) : undefined;

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

	let durationTotal = 0;
	const finish = () => {
		if (event) {
			const neededForMinimum = Math.max(minimumDuration - durationTotal, 0)
			console.log('neededForMinimum', neededForMinimum)
			setTimeout(() => onShown(event!), neededForMinimum * 1000)
		}
	}


	$: audioUrl = eventAlert?.SoundUrl
	$: audioElement = new Audio(audioUrl)
	$: if(audioElement) {
		audioElement.autoplay = true;
		audioElement.onended = () => {
			durationTotal += audioElement.duration
			if(ttsUrl !== "") {
				ttsAudioElement.play()
			} else {
				finish()
			}
		}
	}
	$: ttsUrl = event?.TTSUrl
	$: ttsAudioElement = new Audio(ttsUrl)
	$: ttsAudioElement.onended = () => {
		durationTotal += ttsAudioElement.duration
		finish()
	}

	$: if (audioUrl === "") {
		if(ttsUrl !== "") {
			ttsAudioElement.play()
		} else {
			finish()
		}
	}

</script>

{#if eventAlert && event}
	<div transition:fade class="z-75 fixed inset-0 pointer-events-none mt-20 flex flex-col items-center animate__animated animate__jackInTheBox {!visible ? 'hidden' : ''}">
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