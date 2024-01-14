import { writable } from "svelte/store";
import { signInWithPopup, signOut } from 'firebase/auth';
import { auth, db } from "$lib/firebase/client/config";
import type { AuthProvider } from '@firebase/auth';
import { doc, getDoc, setDoc } from "@firebase/firestore";

export interface User {
    uid: string;
    displayName?: string | null;
    photoURL?: string | null;
    email?: string | null;
    phoneNumber?: string | null;
}

const createUserStore = function () {
    const { subscribe } = writable<User | null>(null, (set) => {
        const unsubscribe = auth.onAuthStateChanged(async (user) => {
            if (user) {
                const userRef = doc(db, `users/${user.uid}`);
                const userDoc = await getDoc(userRef);
                if (!userDoc.exists()) {

                    const saveUser: User = {
                        uid: user.uid,
                        email: user.email,
                        photoURL: user.photoURL,
                        displayName: user.displayName,
                        phoneNumber: user.phoneNumber,
                    };

                    await setDoc(userRef, saveUser);
                }

                await getDoc(userRef).then((doc) => {
                    set(doc.data() as User);
                });
            } else {
                set(null);
            }
        });

        return unsubscribe;
    });

    function logIn(provider: AuthProvider) {
        return signInWithPopup(auth, provider);
    }

    function logOut() {
        return signOut(auth);
    }

    return {
        logOut,
        logIn,
        subscribe,
    }
}

export const user = createUserStore();