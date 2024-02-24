import {db} from "$lib/ext/firebase/firebase";
import { User } from '$lib/pb/stredono_pb';
import {collection, doc, getDocs, query, where, onSnapshot, setDoc} from "firebase/firestore";
import {writable, type Writable} from "svelte/store";

export const userStore: Writable<User | undefined> = writable(undefined);

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
    await setDoc(doc(db, "users", user.Uid), user.toJson({
        useProtoFieldName: true
    }) as any);
}

export const getUserByUsername = async (username: string): Promise<User> => {
    const snapshot = await getDocs(query(collection(db, "users"), where("Username", "==", username)))
    if (snapshot.empty) throw new FetchError("User not found", 404)
    if (snapshot.size > 1) throw new FetchError("Multiple users found", 500)

    const userDoc = snapshot.docs[0]
    return User.fromJson(userDoc.data())
}