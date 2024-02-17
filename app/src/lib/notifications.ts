import {type Writable, writable} from "svelte/store";

export type NotificationType = "success" | "error" | "warning" | "info";

let notificationId = 0;

export class Notification {
    constructor(
        public message: string,
        public duration: number = 5,
        public type: NotificationType = "info",
        public position: "top-right" | "top-left" | "bottom-right" | "bottom-left" = "bottom-right",
        public dismissible: boolean = true,
        public id: number = notificationId++,
    ) {}

    dismiss() {
        notificationsStore.update((notifications) => {
            return notifications.filter((n) => n.id !== this.id);
        });
    }
}

export const notificationsStore: Writable<Notification[]> = writable([]);

export const sendNotification = (notification: Notification) => {
    notificationsStore.update((notifications) => {
        return [...notifications, notification];
    });
}
