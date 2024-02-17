import {initializeApp} from "firebase/app";
import {getAuth} from "firebase/auth";
import {initializeFirestore, persistentLocalCache, CACHE_SIZE_UNLIMITED} from "firebase/firestore";
import {getMessaging} from "firebase/messaging";
import {initializeAppCheck, ReCaptchaEnterpriseProvider} from "firebase/app-check";
import firebaseConfig from "./firebaseWebConfig.json";
import appcheckConfig from "./firebaseAppCheck.json";

export const app = initializeApp(firebaseConfig);
initializeAppCheck(app, {
    provider: new ReCaptchaEnterpriseProvider(appcheckConfig.siteKey),
    isTokenAutoRefreshEnabled: true
})

export const auth = getAuth(app);

export const db = initializeFirestore(app, {
    localCache: persistentLocalCache({
        cacheSizeBytes: CACHE_SIZE_UNLIMITED
    }),
})

export const messaging = getMessaging(app);
