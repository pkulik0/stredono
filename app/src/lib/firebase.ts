import {initializeApp} from "firebase/app";
import {getAuth} from "firebase/auth";
import {getFirestore} from "firebase/firestore";
import {getDownloadURL, getMetadata, getStorage, ref, uploadBytes} from "firebase/storage";
import {getMessaging} from "firebase/messaging";
import firebaseConfig from "./firebaseWebConfig.json";

const app = initializeApp(firebaseConfig);

export const auth = getAuth(app);
export const db = getFirestore(app);
export const storage = getStorage(app);
export const messaging = getMessaging(app);

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