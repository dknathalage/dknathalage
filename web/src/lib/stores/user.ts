import { auth, db } from '$lib/firebase/client/config';
import { derived, writable } from 'svelte/store';
import type { Readable } from 'svelte/store';
import { doc, updateDoc, getDoc, setDoc } from 'firebase/firestore';
import { docStore } from '$lib/stores/realtime';

interface UserData {
	username: string;
	bio: string;
	photoURL: string;
	stripeId: string;
	stripeLink: string;
}

export const user = (function () {
	let unsubscribe: () => void;

	if (!auth || !globalThis.window) {
		console.warn('Auth is not initialized or not in browser');
		const { subscribe } = writable(null);
		return {
			subscribe
		};
	}

	const { subscribe } = writable(auth?.currentUser ?? null, (set) => {
		unsubscribe = auth.onAuthStateChanged(async (user) => {
			set(user);
			if (user) {
				const userRef = doc(db, `users/${user.uid}`);

				// check if user exists in database
				const exists = await getDoc(userRef).then((doc) => doc.exists());

				if (exists) {
					updateDoc(userRef, {
						name: user.displayName,
						email: user.email,
						photoURL: user.photoURL
					});
				} else {
					await setDoc(userRef, {
						name: user.displayName,
						email: user.email,
						photoURL: user.photoURL
					});
				}
			}
		});

		return () => unsubscribe();
	});

	return {
		subscribe
	};
})(); // note the () at the end, this is a self-invoking function. why js? why? :'

export const userData: Readable<UserData | null> = derived(user, ($user, set) => {
	if ($user) {
		return docStore<UserData>(`users/${$user.uid}`).subscribe(set);
	} else {
		set(null);
	}
});
