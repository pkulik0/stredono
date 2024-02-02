import {type Writable, writable} from "svelte/store";
import {auth} from "./firebase";
import {onAuthStateChanged} from "firebase/auth";
import type {User} from "firebase/auth";

type OptionalUser = User | null;

export const userStore: Writable<OptionalUser> = writable(null);

onAuthStateChanged(auth, (u: OptionalUser) => {
    if(u) {
        userStore.set(u)
    } else {
        userStore.set(null)
    }
});