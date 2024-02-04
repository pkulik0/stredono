import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { getStorage } from "firebase/storage";
import { getFirestore } from "firebase/firestore";

const firebaseConfig = {
    apiKey: "AIzaSyDTNZ9x1GD5Y2Euvgmjwh5n70v7MDv9zPs",
    authDomain: "stredono-5ccdd.firebaseapp.com",
    databaseURL: "https://stredono-5ccdd-default-rtdb.europe-west1.firebasedatabase.app",
    projectId: "stredono-5ccdd",
    storageBucket: "stredono-5ccdd.appspot.com",
    messagingSenderId: "621885503876",
    appId: "1:621885503876:web:c7735a53013f52e23d05d1"
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export const storage = getStorage(app);
export const db = getFirestore(app);