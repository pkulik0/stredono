import {db} from "$lib/firebase/firebase";
import {collection, doc, getDocs, query, where, onSnapshot, setDoc} from "firebase/firestore";
import {User} from "$lib/pb/user_pb";
import {writable, type Writable} from "svelte/store";

export const userStore: Writable<User | undefined> = writable(undefined);

export class FetchError extends Error {
    constructor(message: string, public status: number) {
        super(message);
    }
}

export const getUserListener = async (uid: string) => {
    const usersDoc = doc(db, "users", uid);
    return onSnapshot(usersDoc, (doc) => {
        if (!doc.exists()) return;
        userStore.set(User.fromJson(doc.data()));
    })
}

export const saveUser = async (user: User) => {
    const userDoc = doc(db, "users", user.uid);
    await setDoc(userDoc, user.toJson() as any);
}

export const getUserByUsername = async (username: string): Promise<User> => {
    const q = query(collection(db, "users"), where("username", "==", username))
    const snapshot = await getDocs(q)
    if (snapshot.empty) throw new FetchError("User not found", 404)
    if (snapshot.size > 1) throw new FetchError("Multiple users found", 500)

    const userDoc = snapshot.docs[0]
    return User.fromJson(userDoc.data())
}