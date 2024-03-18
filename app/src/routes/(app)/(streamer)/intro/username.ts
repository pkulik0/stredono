import { auth, db } from '$lib/ext/firebase/firebase';
import { query, getDocs, limit, collection, where } from 'firebase/firestore';
import { saveUser, userStore } from '$lib/user';
import { get } from 'svelte/store';

export const saveUsername = async (username: string) => {
	const user = get(userStore);
	if(!user) throw new Error("No user found");

	const authUser = auth.currentUser;
	if (!authUser) throw new Error("User not authenticated")

	user.Username = username;
	await saveUser(user);
}