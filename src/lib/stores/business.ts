import type { Readable } from "svelte/motion";
import { derived } from "svelte/store";
import { RtStore } from "$lib/stores/realtime";
import { authUser as user } from "$lib/stores/user";
import type { Business } from "$lib/models";

export const business: Readable<Business | null> = derived(user, ($user, set) => {
    if ($user && $user?.employments?.length > 0) {
        return RtStore<Business>(`businesses/${$user.employments[0].business.uid}`).subscribe(set);
    } else {
        set(null);
    }
});
