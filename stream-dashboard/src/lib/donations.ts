import {getDocs, doc, onSnapshot, query, where, collection} from "firebase/firestore";
import {auth, db} from "$lib/firebase";
import {get, type Writable, writable} from "svelte/store";
import {SendDonateRequest} from "../../../pb/functions_pb";
import {userStore} from "$lib/userStore";

export const donationsStore: Writable<SendDonateRequest[]> = writable([]);

const firstToLower = (str: string) => {
    return str[0].toLowerCase() + str.slice(1);
}

const keysToLower = (obj: any) => {
    const newObj: any = {};
    for(const key in obj) {
        newObj[firstToLower(key)] = obj[key];
    }
    return newObj;

}

const getDonationsSubscriber = () => {
    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    const donationsRef = collection(db, "donations");
    const q = query(donationsRef, where("RecipientId", "==", user.uid));

    console.log("Subscribing to donations");

    return onSnapshot(q, (snapshot) => {
        snapshot.docChanges().forEach((change) => {
            const data = keysToLower(change.doc.data());
            console.log(data)
            const sdReq = SendDonateRequest.fromJson(data);

            if(change.type === "added") {
                donationsStore.update((donations) => {
                    donations.push(sdReq);
                    return donations;
                })
            }
            if(change.type === "modified") {
                donationsStore.update((donations) => {
                    const index = donations.findIndex((d) => d.id === sdReq.id);
                    if(index === -1) return donations;
                    donations[index] = sdReq;
                    return donations;
                })
            }
            if(change.type === "removed") {
                donationsStore.update((donations) => {
                    return donations.filter((d) => d.id !== sdReq.id);
                })
            }
        })
    });
}

let unsubscribe: (() => void) | null = null;

userStore.subscribe((user) => {
    if(user) {
        unsubscribe = getDonationsSubscriber();
    } else {
        if(unsubscribe) {
            unsubscribe();
            unsubscribe = null;
        }
        donationsStore.set([]);
    }
})