import { db } from '$lib/ext/firebase/firebase';
import {getDoc, doc} from 'firebase/firestore';

export const getVoices = async () => {
	const docSnap = await getDoc(doc(db, 'tts/voices'));
	if (docSnap.exists()) {
		return docSnap.data();
	} else {
		console.log('No such document!');
	}
}