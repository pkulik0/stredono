<script lang="ts">
	import { eventToAlert } from '$lib/alerts';
	import { Alert, Alignment, AnimationType, Event, Position, Speed } from '$lib/pb/stredono_pb';
	import 'animate.css';
	export let alerts: Alert[];

	export let visible: boolean = true;
	export let event: Event|undefined;
	$: eventAlert = event ? eventToAlert(event, alerts) : undefined;

	$: message = eventAlert ? eventAlert.Message : '';
	$: if(event && eventAlert) {
		Object.entries(event.Data).forEach(([key, value]) => {message = message.replace(`[${key}]`, `<span style="color: ${eventAlert.AccentColor};">${value}</span>`);});
	}
	$: header = message.split('\n')[0];
	$: content = event ? event.Data.message : '';

	$: alignment = eventAlert ? eventAlert.Alignment : Alignment.CENTER;
	$: alignmentClass = (() => {
		switch(alignment) {
			case Alignment.START: return 'text-start';
			case Alignment.CENTER: return 'text-center';
			case Alignment.END: return 'text-end';
			case Alignment.JUSTIFY: return 'text-justify';
		}
	})();

	$: textPosition = eventAlert ? eventAlert.TextPosition : Position.TOP;
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

	let timeout: any;
	$: {
		clearTimeout(timeout);
		if(event) {
			timeout = setTimeout(() => {
				event = event;
			}, 10000);
		}
	}

	let audioElement: HTMLAudioElement;
	let audioTimeout: any;
	let lastUrl = '';
	$: {
		if(eventAlert && eventAlert.SoundUrl) {
			if(lastUrl !== eventAlert.SoundUrl) {
				lastUrl = eventAlert.SoundUrl;
				clearTimeout(audioTimeout);
				audioTimeout = setTimeout(() => {
					audioElement.play();
				}, 200);
			}
		}
	}

</script>

{#if eventAlert && event}
	<div class="z-75 fixed inset-0 pointer-events-none mt-20 flex flex-col items-center animate__animated animate__jackInTheBox {!visible ? 'hidden' : ''}">
		<div class="flex items-center  {alignmentClass} {containerClass} {animationClass}" style="color: {eventAlert.TextColor}">
			{#if eventAlert.GifUrl}
				<img src={eventAlert.GifUrl} alt="" class="rounded-xl shadow-2xl {imgClass} max-w-sm" />
			{/if}

			<div class="flex flex-col space-y-2 max-w-lg">
				<h5 class="font-extrabold text-2xl pt-2 {alignment === Alignment.JUSTIFY ? 'text-center' : ''}">{@html header}</h5>
				<p class="text-lg">{@html content}</p>
			</div>
		</div>
	</div>
	<audio src={eventAlert.SoundUrl} bind:this={audioElement} class="hidden" />
{/if}
