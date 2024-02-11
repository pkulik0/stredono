import {db, storage} from "$lib/firebase";
import {collection, doc, getDoc, setDoc} from "firebase/firestore";
import {getDownloadURL, ref} from "firebase/storage";
import {Profile} from "$lib/pb/profile_pb";

export class FetchError extends Error {
    constructor(message: string, public status: number) {
        super(message);
    }
}

export const getProfileByUsername = async (username: string): Promise<Profile> => {
    const pageDoc = await getDoc(doc(db, "pages", username))
    if (!pageDoc.exists()) throw new FetchError("Page not found", 404)
    return Profile.fromJson(pageDoc.data())
}

export const saveProfileToDb = async (username: string, profile: Profile) => {
    await setDoc(doc(db, "pages", username), profile.toJson() as object)
}

export const getAvatarUrl = async (uid: string): Promise<string> => {
    const avRef = ref(storage, `users/${uid}/page/avatar`);
    return await getDownloadURL(avRef)
}