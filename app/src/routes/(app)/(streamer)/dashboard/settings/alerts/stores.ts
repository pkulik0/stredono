import { getDefaultAlert } from '$lib/alerts';
import { Alert, Alignment, AnimationType, Position, Speed } from '$lib/pb/alert_pb';
import { EventType } from '$lib/pb/event_pb';
import { writable, type Writable } from 'svelte/store';

export const alertStore: Writable<Alert> = writable(getDefaultAlert(EventType.TIP));
export const drawerHiddenStore: Writable<boolean> = writable(true);