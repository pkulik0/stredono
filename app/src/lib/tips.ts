import { Tip, TipStatus } from '$lib/pb/tip_pb';
import {collection, onSnapshot, query, where, getDocs} from "firebase/firestore";
import {auth, db} from "$lib/ext/firebase/firebase";
import {userStore} from "$lib/user";
import {writable, type Writable} from "svelte/store";

export type TipsMap = {[key: number]: { tips: Tip[], startDate: Date, endDate: Date }}
const idToEntry: Map<string, { key: number, index: number }> = new Map();

export const tipsStore: Writable<TipsMap> = writable({});

export interface WebTip {
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

export const fetchOldTips = async (date: Date, page: number) => {
    if (page < 1) throw new Error("Invalid page number");

    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    const startDate = new Date(date);
    startDate.setDate(startDate.getDate() - 7 * page);

    const endDate = new Date(date);
    endDate.setDate(endDate.getDate() - 7 * (page - 1));

    const q = query(collection(db, "tips"),
        where("RecipientId", "==", user.uid),
        where("Timestamp", ">=", dateToFirestoreTimestamp(startDate)),
        where("Timestamp", "<", dateToFirestoreTimestamp(endDate))
    );
    const snapshot = await getDocs(q);

    tipsStore.update((pastTips) => {
        pastTips[page] = {
            tips: [],
            startDate: startDate,
            endDate: endDate
        }

        snapshot.forEach((doc) => {
            const tip = Tip.fromJson(doc.data());
            if(tip.Status < TipStatus.PAYMENT_SUCCESS) return;
            pastTips[page].tips.push(tip);
        });

        return pastTips
    })
}

const getTipsSubscriber = () => {
    const user = auth.currentUser;
    if(!user) throw new Error("Not logged in");

    const date = getWeekBeginning()
    const q = query(collection(db, "tips"),
        where("RecipientId", "==", user.uid),
        where("Timestamp", ">=", dateToFirestoreTimestamp(date))
    );

    return onSnapshot(q, (snapshot) => {
        tipsStore.update((tips) => {
            if(!tips[0]) {
                const endDate = new Date(date);
                endDate.setDate(endDate.getDate() + 7);

                tips[0] = {
                    tips: [],
                    startDate: date,
                    endDate: endDate
                }
            }

            snapshot.docChanges().forEach((change) => {
                const id = change.doc.id;

                const tip = Tip.fromJson(change.doc.data());
                if(tip.Status < TipStatus.PAYMENT_SUCCESS) return;

                if(change.type === "added" || change.type === "modified") {
                    const entry = idToEntry.get(id)
                    if (entry) {
                        tips[entry.key].tips[entry.index] = tip;
                        return;
                    }

                    const len = tips[0].tips.push(tip);
                    idToEntry.set(id, { key: 0, index: len - 1 });
                } else if(change.type === "removed") {
                    const entry = idToEntry.get(id);
                    if(entry) {
                        tips[entry.key].tips.splice(entry.index, 1);
                    }
                }
            })

            return tips;
        })

    });
}

let unsubscribe: (() => void) | null = null;

userStore.subscribe((user) => {
    try {
        if(user) {
            unsubscribe = getTipsSubscriber();
        } else {
            if(unsubscribe) {
                unsubscribe();
                unsubscribe = null;
            }
            tipsStore.set({});
        }
    } catch (e) {
        console.error(e);
    }
})