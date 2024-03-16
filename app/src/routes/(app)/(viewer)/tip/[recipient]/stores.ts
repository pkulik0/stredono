import {writable} from "svelte/store";

export const emailStore = writable(localStorage.getItem("email") || "");
export const senderStore = writable((localStorage.getItem("sender") || ""));

emailStore.subscribe(value => {
    localStorage.setItem("email", value);
});

senderStore.subscribe(value => {
    localStorage.setItem("sender", value);
});