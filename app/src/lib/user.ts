import { auth, db } from '$lib/ext/firebase/firebase';
import { User } from '$lib/pb/user_pb';
import { terraformOutput } from '$lib/constants';
import axios from 'axios';
import { collection, doc, getDocs, query, where, getDoc, setDoc, limit } from 'firebase/firestore';
import {writable, type Writable} from "svelte/store";

export const userStore: Writable<User | undefined> = writable(undefined);

export const getUserByUid = async (uid: string) => {
    const d = await getDoc(doc(db, "users", uid))
    if (!d.exists()) throw new Error("No user found")
    return User.fromJson(d.data())
}

export const saveUser = async (user: User) => {
    const authUser = auth.currentUser;
    if (!authUser) throw new Error("User not authenticated")
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
    const snapshot = await getDocs(query(collection(db, "users"), where("Username", "==", username), limit(1)))
    if (snapshot.empty) throw new Error("No user found")

    const userDoc = snapshot.docs[0]
    return User.fromJson(userDoc.data())
}