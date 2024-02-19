import axios from "axios";
import {auth} from "$lib/firebase/firebase";

export const getTwitchAuthUrl = async (): Promise<string> => {
    const user = auth.currentUser;
    if (!user) {
        throw new Error("Not logged in");
    }
    const token = await user.getIdToken(true);

    const url = "connectTwitch?redirect=" + encodeURIComponent(window.location.href);
    const headers = {
        "Authorization": "Bearer " + token
    }
    const res = await axios.get(url, {headers});
    if (!res.data) {
        throw new Error("No data");
    }

    return res.data;
}