import { readable } from 'svelte/store';
import { auth, db } from '$lib/firebase/client/config';
import { doc, getDoc, setDoc } from 'firebase/firestore';
import type { User } from '$lib/models';

export const authUser = readable<any | null>(null, (set) => {
	const unsubscribe = auth.onAuthStateChanged(async (user) => {
		if (user) {
			set({
				uid: user.uid,
				displayName: user.displayName,
				photoURL: user.photoURL,
				email: user.email,
				phoneNumber: user.phoneNumber,
			});
			const userRef= doc(db, `users/${user.uid}`);
			const userDoc = await getDoc(userRef);
			if (!userDoc.exists()) {
				const saveUser: User = {
					uid: user.uid,
					phone: user.phoneNumber,
					employments: [],
				};

				await setDoc(userRef, saveUser);
			}
		} else {
			set(null);
		}
	});

	return unsubscribe;
});
