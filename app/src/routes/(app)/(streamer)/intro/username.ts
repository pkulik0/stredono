import { auth, db } from '$lib/ext/firebase/firebase';
import { query, getDocs, limit, collection, where } from 'firebase/firestore';
import { saveUser, userStore } from '$lib/user';
import { get } from 'svelte/store';

export const checkIfUnique = async (username: string) => {
	if(!username) return true;

	const qSnap = await getDocs(query(collection(db, "users"), where("Username", "==", username), limit(1)));
	return qSnap.empty;
}

export const saveUsername = async (username: string) => {
	const user = get(userStore);
	if(!user) throw new Error("No user found");

	const authUser = auth.currentUser;
	if (!authUser) throw new Error("User not authenticated")

	user.Username = username;
	await saveUser(user);
}