import { writable } from 'svelte/store';
import { auth } from '$lib/firebase/client/config';

interface UserData {
	uid: string;
	displayName: string | null;
	photoURL: string | null;
}

export const userStore = writable<UserData | null>(null);
auth.onAuthStateChanged(userStore.set);