import { auth, db } from '$lib/ext/firebase/firebase';
import { User } from '$lib/pb/user_pb';
import { terraformOutput } from '$lib/constants';
import axios from 'axios';
import {collection, doc, getDocs, query, where, onSnapshot, setDoc} from "firebase/firestore";
import {writable, type Writable} from "svelte/store";

export const userStore: Writable<User | null | undefined> = writable(undefined);

export class FetchError extends Error {
    constructor(message: string, public status: number) {
        super(message);
    }
}

export const getUserListener = async (uid: string) => {
    return onSnapshot(doc(db, "users", uid), (doc) => {
        if (!doc.exists()) return;
        userStore.set(User.fromJson(doc.data()));
    })
}

export const saveUser = async (user: User) => {
    const authUser = auth.currentUser;
    if (!authUser) throw new FetchError("User not logged in", 401);
    const token = await authUser.getIdToken();

    try {
        await axios.post(terraformOutput.FunctionsUrl + "/UserEdit", user.toBinary(), {
            headers: {
                'Content-Type': 'application/octet-stream',
                'Authorization': 'Bearer ' + token
            }
        })
    } catch (e) {
        console.error(e)
    }
}

export const getUserByUsername = async (username: string): Promise<User> => {
    const snapshot = await getDocs(query(collection(db, "users"), where("Username", "==", username)))
    if (snapshot.empty) throw new FetchError("User not found", 404)
    if (snapshot.size > 1) throw new FetchError("Multiple users found", 500)

    const userDoc = snapshot.docs[0]
    return User.fromJson(userDoc.data())
}