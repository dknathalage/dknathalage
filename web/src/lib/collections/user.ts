import { doc, writeBatch } from 'firebase/firestore';
import { db } from '../firebase/client/config';
import type { User } from 'firebase/auth';

export const add = async function (user: User) {
	if (!user) {
		return;
	}

	console.log(user);

	const batch = writeBatch(db);
	batch.set(doc(db, 'users', user!.uid), {
		uid: user.uid,
		email: user.email,
		name: user.displayName,
		photoURL: user.photoURL
	});
	await batch.commit();
};
