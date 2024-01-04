import { auth } from '$lib/firebase/client/config';
// import { user } from "$lib/stores/user" auth state changes here when users are logged in from here
import { GoogleAuthProvider, signInWithRedirect, signOut } from 'firebase/auth';

export const signin = async function () {
	const provider = new GoogleAuthProvider();
	await signInWithRedirect(auth, provider);
};

export const signout = async function () {
	signOut(auth);
};
