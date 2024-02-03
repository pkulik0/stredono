import {type Writable, writable} from "svelte/store";
import {auth} from "./firebase";
import {onAuthStateChanged} from "firebase/auth";
import type {User} from "firebase/auth";

export type OptionalUser = User | null;

export const userStore: Writable<OptionalUser> = writable(null);