import { db } from '$lib/firebase/client/config';
import { doc, setDoc, collection, onSnapshot } from 'firebase/firestore';
import { initialiseStripe } from '$lib/payments/initialise';

export async function create(uid: string) {
    console.log('create session');

    // Create a reference to the Firestore document
    const docRef = doc(collection(db, 'users', uid, 'checkout_sessions'));

    // Set the document data
    await setDoc(docRef, {
        price: 'price_1OUgB5KjvXvpGllS4U5odCgw',
        success_url: window.location.origin,
        cancel_url: window.location.origin,
    });

    // Listen to the document
    onSnapshot(docRef, async (doc) => {
        const data = doc.data();
        const sessionId = data?.sessionId;
        if (sessionId) {
            const stripe = await initialiseStripe();
            stripe?.redirectToCheckout({ sessionId });
        }
    });

}