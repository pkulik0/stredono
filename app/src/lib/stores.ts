import {type Writable, writable} from "svelte/store";
import type {User} from "firebase/auth";
import {getUserDoc} from "$lib/alerts";

export const userStore: Writable<User | null> = writable(null);
export const usernameStore: Writable<string | null> = writable(null);

userStore.subscribe(async (value) => {
    if (!value) {
        usernameStore.set(null);
        return;
    }

    const userDoc = await getUserDoc();
    const username = userDoc.data.username;
    usernameStore.set(username);
});

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