import { db } from '$lib/firebase/client/config';
import { doc, setDoc, onSnapshot, collection } from 'firebase/firestore';
import { initialiseStripe } from '$lib/payments/initialise';

export async function create(uid: string) {
    // Create a reference to the Firestore document
    const docRef = doc(collection(db, 'users', uid, 'checkout_sessions'));

    console.log("done 1")

    // Set the document data
    await setDoc(docRef, {
        price: 'price_1OUjxDKjvXvpGllS8ZTHThE2',
        success_url: window.location.origin,
        cancel_url: window.location.origin,
    });

    console.log("done 2")

    // Listen to the document
    await onSnapshot(docRef, async (doc) => {
        const data = doc.data();
        const sessionId = data?.sessionId;
        if (sessionId) {
            const stripe = await initialiseStripe();
            stripe?.redirectToCheckout({ sessionId });
        }
    });
}