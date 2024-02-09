import {type Writable, writable} from "svelte/store";
import type {User} from "firebase/auth";

export const userStore: Writable<User | null> = writable(null);
export const streamModeStore: Writable<boolean> = writable(localStorage.getItem("streamMode") === "true");

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