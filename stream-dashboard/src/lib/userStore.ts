import {type Writable, writable} from "svelte/store";
import {auth} from "./firebase";
import {onAuthStateChanged} from "firebase/auth";
import type {User} from "firebase/auth";

export const userStore: Writable<User | null> = writable(null);