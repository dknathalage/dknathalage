import type { Readable } from "svelte/motion";
import { derived } from "svelte/store";
import { authUser as user } from "$lib/stores/user";
import type { Business, User, Employment } from "$lib/models";
import { db } from "$lib/firebase/client/config";
import { doc, collection, writeBatch } from "firebase/firestore";

export const business: Readable<Business | null> = derived(user, ($user, set) => {
    if ($user && $user?.employments?.length > 0) {
        return set($user.employments[0].business);
    } else {
        set(null);
    }
});

export const createBusiness = async (business: Business, user: User): Promise<any> => {


    const batch = writeBatch(db)

    const employment: Employment = {
        role: "owner",
        user: {
            uid: user.uid,
            phone: user.phone,
        },
        business: {
            uid: "",
            name: business.name,
        },
    }

    // const businessRef = doc(db, "businesses", business.uid)
    // const userRef = doc(db, "users", user.uid)

    // business collection
    const businessesRef = doc(collection(db, "businesses"))
    batch.set(businessesRef, business)

    const businessRef = doc(collection(db, "businesses"), businessesRef.id)

    // user employments collection
    batch.set(doc(collection(db, "users", user.uid, "employments")), employment)

    // // business employments collection
    batch.set(doc(collection(db, "businesses", businessRef.id, "employments")), employment)

    batch.commit()
}