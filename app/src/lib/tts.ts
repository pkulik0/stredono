import { db } from '$lib/ext/firebase/firebase';
import { Tier } from '$lib/pb/enums_pb';
import { Voice, Voices } from '$lib/pb/tts_pb';
import {getDoc, doc} from 'firebase/firestore';

export const getVoices = async () => {
	const docSnap = await getDoc(doc(db, 'tts/voices'));
	if (!docSnap.exists()) {
		throw new Error('No voices found');
	}

	let voicesPlus: Voice[] = [];
	let voicesBasic: Map<string, Voice[]> = new Map();

	Voices.fromJson(docSnap.data()).Voices.forEach(v => {
		if(v.Tier == Tier.PLUS) {
			voicesPlus.push(v)
			return
		}
		const code = v.Languages[0].split("-")[0]
		let voices = voicesBasic.get(code) || []
		voices.push(v)
		voicesBasic.set(code, voices)
	})

	return {voicesBasic, voicesPlus}
}