import { readable } from 'svelte/store';
import { auth } from '$lib/firebase/client/config';

interface UserData {
	uid: string;
	displayName: string | null;
	photoURL: string | null;
	email: string | null;
}

export const userStore = readable<UserData | null>(null, (set) => {
	const unsubscribe = auth.onAuthStateChanged((user) => {
		if (user) {
			set({
				uid: user.uid,
				displayName: user.displayName,
				photoURL: user.photoURL,
				email: user.email
			});
		} else {
			set(null);
		}
	});

	return unsubscribe;
});
