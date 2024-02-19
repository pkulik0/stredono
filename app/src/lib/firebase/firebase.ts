import {initializeApp} from "firebase/app";
import {getAuth} from "firebase/auth";
import {initializeFirestore, persistentLocalCache, CACHE_SIZE_UNLIMITED} from "firebase/firestore";
import {getMessaging} from "firebase/messaging";
import {initializeAppCheck, ReCaptchaEnterpriseProvider} from "firebase/app-check";
import TerraformOutput from "../terraform_output.json";

export const app = initializeApp(TerraformOutput.FirebaseWebappConfig);
initializeAppCheck(app, {
    provider: new ReCaptchaEnterpriseProvider(TerraformOutput.RecaptchaSiteKey),
    isTokenAutoRefreshEnabled: true
})

export const auth = getAuth(app);

export const db = initializeFirestore(app, {
    localCache: persistentLocalCache({
        cacheSizeBytes: CACHE_SIZE_UNLIMITED
    }),
})

export const messaging = getMessaging(app);
