import type { Readable } from "svelte/motion";
import { derived } from "svelte/store";
import { RtStore } from "$lib/stores/realtime";
import { authUser as user } from "$lib/stores/user";
import type { Business } from "$lib/models";
import { db } from "$lib/firebase/client/config";
import { doc, collection, setDoc } from "firebase/firestore";

export const business: Readable<Business | null> = derived(user, ($user, set) => {
    if ($user && $user?.employments?.length > 0) {
        return RtStore<Business>(`businesses/${$user.employments[0].business.uid}`).subscribe(set);
    } else {
        set(null);
    }
});

export const createBusiness = (business: Business): Promise<any> => {
    const businessRef = collection(db, "businesses")
    const businessDocRef = doc(businessRef)
    return setDoc(businessDocRef, business)
}