import axios from "axios";
import {PUBLIC_FUNC_LINK} from "$env/static/public";
import {auth} from "$lib/firebase/firebase";

export const getTwitchAuthUrl = async (): Promise<string> => {
    const user = auth.currentUser;
    if (!user) {
        throw new Error("Not logged in");
    }
    const token = await user.getIdToken(true);

    const url = PUBLIC_FUNC_LINK + "connectTwitch?redirect=" + encodeURIComponent(window.location.href);
    const headers = {
        "Authorization": "Bearer " + token
    }
    const res = await axios.get(url, {headers});
    if (!res.data) {
        throw new Error("No data");
    }

    return res.data;
}