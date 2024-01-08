import type { Readable } from "svelte/motion";
import { derived } from "svelte/store";
import { RtStore } from "./stores/realtime";
import { authUser as user } from "./stores/user";

interface BusinessMember {
    uid: string;
    role: string;
    user: User;
    business: Business;
}

export interface User{
    uid: string;
    phone: string | null;
    employments: BusinessMember[];
}

export interface Business {
    uid: string;
    name: string;
    employments: BusinessMember[];
}

export const business: Readable<Business | null> = derived(user, ($user, set) => {
    if ($user && $user?.employments?.length > 0) {
        return RtStore<Business>(`businesses/${$user.employments[0].business.uid}`).subscribe(set);
    } else {
        set(null);
    }
});
