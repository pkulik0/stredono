import {doc, setDoc, getDoc} from "firebase/firestore";
import {auth, db} from "$lib/firebase";

import {Alert} from "$lib/pb/alerts_pb";

interface UserDoc {
    doc: any,
    data: any
}

const getUserDoc = async (): Promise<UserDoc> => {
    const user = auth.currentUser;
    if (!user) {
        throw new Error("User not found");
    }

    const userDoc = doc(db, "users", user.uid);
    const userSnap = await getDoc(userDoc);
    if (!userSnap.exists()) {
        throw new Error("User not found: " + user.uid);
    }

    const userDocData = userSnap.data();
    if (!userDocData) {
        throw new Error("User not found: " + user.uid);
    }

    return {
        doc: userDoc,
        data: userDocData
    } as UserDoc;
}

export const createAlert = async (alert: Alert) => {
    const userDoc = await getUserDoc();
    const doc = userDoc.doc;
    const data = userDoc.data;

    const alerts = data.alerts || [];
    alerts.push(alert.toJson());

    await setDoc(doc, {alerts: alerts}, {merge: true});
}

export const getAlerts = async (): Promise<Alert[]> => {
    const userDoc = await getUserDoc();
    const data = userDoc.data;

    const alerts = data.alerts || [];
    return alerts.map((alert: any) => Alert.fromJson(alert));
}