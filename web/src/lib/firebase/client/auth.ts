import { auth } from '$lib/firebase/client/config';
import { signInWithRedirect, signOut } from 'firebase/auth';
import type { AuthProvider } from '@firebase/auth';
import type { Auth } from 'firebase/auth';

export const signin = async function (provider: AuthProvider) {
	await signInWithRedirect(auth, provider);
};

export const signout = async function (auth: Auth) {
	await signOut(auth);
};
