import {type Writable, writable} from "svelte/store";

export const streamModeStore: Writable<boolean> = writable(localStorage.getItem("streamMode") === "true");

streamModeStore.subscribe((value) => {
    localStorage.setItem("streamMode", value.toString());
})

//----

export const emailStore = writable(localStorage.getItem("email") || "");
export const senderStore = writable((localStorage.getItem("sender") || ""));

emailStore.subscribe(value => {
    localStorage.setItem("email", value);
});

senderStore.subscribe(value => {
    localStorage.setItem("sender", value);
});