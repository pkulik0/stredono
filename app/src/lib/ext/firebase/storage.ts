import { isLocal } from '$lib/constants';
import {getDownloadURL, getMetadata, getStorage, ref, uploadBytes, connectStorageEmulator} from "firebase/storage";
import {auth, app} from "$lib/ext/firebase/firebase";

export const storage = getStorage(app);
if(isLocal) {
    connectStorageEmulator(storage, 'localhost', 30506);
    console.log("Storage emulator connected");
}

export const uploadToStorage = async (folder: string, name: string, file: File, overwrite: boolean): Promise<string> => {
    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    if(!folder || !name || !file) throw new Error("Invalid path, name or file");

    let fullPath = "users/" + user.uid + "/" + folder;
    if(!folder.endsWith("/")) fullPath += "/";
    fullPath += name;

    const storageRef = ref(storage, fullPath);
    if(!overwrite) {
        try {
            const metadata = await getMetadata(storageRef);
            if(metadata.size > 0) {
                throw new Error("File already exists");
            }
        } catch (e: any) {
            if(e.code !== "storage/object-not-found") throw e;
        }
    }

    const bytes = await file.arrayBuffer();
    const snapshot = await uploadBytes(storageRef, new Uint8Array(bytes));
    return await getDownloadURL(snapshot.ref);
}