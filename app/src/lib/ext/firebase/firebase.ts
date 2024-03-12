import { isLocal } from '$lib/constants';
import {initializeApp} from "firebase/app";
import {getAuth, connectAuthEmulator} from "firebase/auth";
import {initializeFirestore, persistentLocalCache, CACHE_SIZE_UNLIMITED, connectFirestoreEmulator, persistentMultipleTabManager} from "firebase/firestore";
import {initializeAppCheck, ReCaptchaEnterpriseProvider} from "firebase/app-check";
import {getDatabase, connectDatabaseEmulator} from "firebase/database";
import TerraformOutput from "../../terraform_output.json";

export const app = initializeApp(TerraformOutput.FirebaseWebappConfig);
initializeAppCheck(app, {
    provider: new ReCaptchaEnterpriseProvider(TerraformOutput.RecaptchaSiteKey),
    isTokenAutoRefreshEnabled: true
})

export const auth = getAuth(app);
if(isLocal) {
    connectAuthEmulator(auth, 'http://localhost:30501');
    console.log("Auth emulator connected");
}

export const db = initializeFirestore(app, {
    localCache: persistentLocalCache({
        cacheSizeBytes: CACHE_SIZE_UNLIMITED,
        tabManager: persistentMultipleTabManager()
    }),
})
if(isLocal) {
    connectFirestoreEmulator(db, 'localhost', 30502);
    console.log("Firestore emulator connected");
}

export const rtdb = getDatabase(app);
if(isLocal) {
    connectDatabaseEmulator(rtdb, 'localhost', 30503);
    console.log("RTDB emulator connected");
}
