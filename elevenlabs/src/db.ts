import { InvocationType, printStats } from './main';
import { TTSKey } from './pb/stredono_pb';
import { getFirestore } from 'firebase-admin/firestore';
import { initializeApp, applicationDefault } from 'firebase-admin/app';

const db = getFirestore(initializeApp({
	credential: applicationDefault()
}));

export const saveKey = async (i: number, key: string) => {
	const ttsKey = new TTSKey();
	ttsKey.Key = key;
	ttsKey.LastUsed = BigInt(Math.floor(Date.now() / 1000))
	ttsKey.CharactersLimit = 10000 // default for free accounts
	ttsKey.CharactersLeft = ttsKey.CharactersLimit

	const docRef = await db.collection("elevenlabs").add(ttsKey.toJson({
		 useProtoFieldName: true,
		}) as any)

	console.log(i, `Wrote key (${docRef.id})`);
	printStats(InvocationType.SUCCESS)
}