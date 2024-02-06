import {type Writable, writable} from "svelte/store";
import type {User} from "firebase/auth";
import {SendDonateRequest} from "$lib/pb/functions_pb";

export const userStore: Writable<User | null> = writable(null);
export const streamModeStore: Writable<boolean> = writable(localStorage.getItem("streamMode") === "true");
export const donationsStore: Writable<SendDonateRequest[]> = writable([]);

streamModeStore.subscribe((value) => {
    localStorage.setItem("streamMode", value.toString());
})

//----

export const emailStore = writable(localStorage.getItem("email") || "");
export const usernameStore = writable((localStorage.getItem("username") || ""));

emailStore.subscribe(value => {
    localStorage.setItem("email", value);
});

usernameStore.subscribe(value => {
    localStorage.setItem("username", value);
});