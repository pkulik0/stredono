import { writable, type Writable } from 'svelte/store';

export const dashboardUidStore: Writable<string> = writable('');