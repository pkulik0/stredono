import {collection, onSnapshot, query, where, getDocs} from "firebase/firestore";
import {auth, db} from "$lib/firebase/firebase";
import {DonateStatus, SendDonateRequest} from "$lib/pb/functions_pb";
import {userStore} from "$lib/user";
import {writable, type Writable} from "svelte/store";

export type DonationsMap = {[key: number]: { donate: SendDonateRequest[], startDate: Date, endDate: Date }}

export const donationStore: Writable<DonationsMap> = writable({});

export interface Donate {
    user: string;
    amount: number;
    currency: string;
    message: string;
}

const dateToFirestoreTimestamp = (date: Date) => {
    return Math.floor(date.getTime() / 1000);
}

const getWeekBeginning = (): Date => {
    const now = new Date();
    const day = now.getDay();
    const diff = now.getDate() - day + (day === 0 ? -6 : 1);
    now.setHours(0, 0, 0, 0);
    now.setDate(diff)
    return now;
}

export const fetchOldDonations = async (date: Date, page: number) => {
    if (page < 1) throw new Error("Invalid page number");

    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    const startDate = new Date(date);
    startDate.setDate(startDate.getDate() - 7 * page);

    const endDate = new Date(date);
    endDate.setDate(endDate.getDate() - 7 * (page - 1));

    const q = query(collection(db, "donations"),
        where("RecipientId", "==", user.uid),
        where("Timestamp", ">=", dateToFirestoreTimestamp(startDate)),
        where("Timestamp", "<", dateToFirestoreTimestamp(endDate))
    );
    const snapshot = await getDocs(q);

    donationStore.update((pastDonations) => {
        pastDonations[page] = {
            donate: [],
            startDate: startDate,
            endDate: endDate
        }

        snapshot.forEach((doc) => {
            const sdReq = SendDonateRequest.fromJson(doc.data());
            if(sdReq.Status < DonateStatus.PAYMENT_SUCCESS) return;
            pastDonations[page].donate.push(sdReq);
        });

        return pastDonations
    })
}

const getDonationsSubscriber = () => {
    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    const date = getWeekBeginning()
    const q = query(collection(db, "donations"),
        where("RecipientId", "==", user.uid),
        where("Timestamp", ">=", dateToFirestoreTimestamp(date))
    );

    return onSnapshot(q, (snapshot) => {
        donationStore.update((donations) => {
            if(!donations[0]) {
                const endDate = new Date(date);
                endDate.setDate(endDate.getDate() + 7);

                donations[0] = {
                    donate: [],
                    startDate: date,
                    endDate: endDate
                }
            }
            let latestDonations = donations[0].donate;

            snapshot.docChanges().forEach((change) => {
                const sdReq = SendDonateRequest.fromJson(change.doc.data());
                if(sdReq.Status < DonateStatus.PAYMENT_SUCCESS) return;

                if(change.type === "added" || change.type === "modified") {
                    const index = latestDonations.findIndex((d) => d.Id === sdReq.Id);
                    if(index === -1) {
                        latestDonations.push(sdReq);
                        return;
                    }
                    latestDonations[index] = sdReq;
                } else if(change.type === "removed") {
                    latestDonations =  latestDonations.filter((d) => d.Id !== sdReq.Id);
                }
            })

            return donations;
        })

    });
}

let unsubscribe: (() => void) | null = null;

userStore.subscribe((user) => {
    try {
        if(user) {
            unsubscribe = getDonationsSubscriber();
        } else {
            if(unsubscribe) {
                unsubscribe();
                unsubscribe = null;
            }
            donationStore.set({});
        }
    } catch (e) {
        console.error(e);
    }
})