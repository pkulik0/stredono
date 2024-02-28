import { isLocal } from '$lib/constants';
import {initializeApp} from "firebase/app";
import {getAuth, connectAuthEmulator} from "firebase/auth";
import {initializeFirestore, persistentLocalCache, CACHE_SIZE_UNLIMITED, connectFirestoreEmulator} from "firebase/firestore";
import {getMessaging} from "firebase/messaging";
import {initializeAppCheck, ReCaptchaEnterpriseProvider} from "firebase/app-check";
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
        cacheSizeBytes: CACHE_SIZE_UNLIMITED
    }),
})
if(isLocal) {
    connectFirestoreEmulator(db, 'localhost', 30502);
    console.log("Firestore emulator connected");
}

export const messaging = getMessaging(app);
