import {collection, onSnapshot, query, where} from "firebase/firestore";
import {auth, db} from "$lib/firebase";
import {DonateStatus, SendDonateRequest} from "$lib/pb/functions_pb";
import {donationsStore, userStore} from "$lib/stores";

export interface Donate {
    user: string;
    amount: number;
    currency: string;
    message: string;
}

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

const getWeekBeginning = (): Date => {
    const now = new Date();
    const day = now.getDay();
    const diff = now.getDate() - day + (day === 0 ? -6 : 1);
    now.setHours(0, 0, 0, 0);
    now.setDate(diff)
    return now;
}

const getDonationsSubscriber = () => {
    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    const donationsRef = collection(db, "donations");
    const q = query(donationsRef,
        where("RecipientId", "==", user.uid),
        where("Timestamp", ">=", getWeekBeginning().getTime() / 1000)
    );

    console.log("Subscribing to donations");

    return onSnapshot(q, (snapshot) => {
        console.log("Donations snapshot");

        donationsStore.update((donations) => {
            snapshot.docChanges().forEach((change) => {
                console.log("Donation change", change.type, change.doc.data());

                const sdReq = SendDonateRequest.fromJson(keysToLower(change.doc.data()));
                if(sdReq.status < DonateStatus.PAYMENT_SUCCESS) return;

                if(change.type === "added" || change.type === "modified") {
                    const index = donations.findIndex((d) => d.id === sdReq.id);
                    if(index === -1) {
                        donations.push(sdReq);
                        return;
                    }
                    donations[index] = sdReq;
                } else if(change.type === "removed") {
                    donations =  donations.filter((d) => d.id !== sdReq.id);
                }
            })

            return donations;
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