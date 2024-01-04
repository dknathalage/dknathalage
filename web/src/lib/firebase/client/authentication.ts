import { auth } from '$lib/firebase/client/config';
import { GoogleAuthProvider, signInWithPopup, signOut } from 'firebase/auth';

export const signin = async function () {
	const provider = new GoogleAuthProvider();
	await signInWithPopup(auth, provider);
};

export const signout = async function () {
	signOut(auth);
};
