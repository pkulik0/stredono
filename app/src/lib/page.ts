import {db, storage} from "$lib/firebase";
import {doc, getDoc} from "firebase/firestore";
import {getDownloadURL, ref} from "firebase/storage";

export class FetchError extends Error {
    constructor(message: string, public status: number) {
        super(message);
    }
}


export const getPageDocByUsername = async (username: string) => {
    const pageDoc = await getDoc(doc(db, "pages", username))
    if (!pageDoc.exists()) throw new FetchError("Page not found", 404)
    return pageDoc
}

export const getAvatarUrl = async (uid: string): Promise<string> => {
    const avRef = ref(storage, `users/${uid}/page/avatar`);
    return await getDownloadURL(avRef)
}