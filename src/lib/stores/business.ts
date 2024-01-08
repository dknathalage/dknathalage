import { readable } from 'svelte/store';
import { db } from '$lib/firebase/client/config';
import { collection } from '@firebase/firestore';

interface BusinessOwner {
    uid: string;
    user_id: string;
    business_id: string;
}

interface BusinessData {
    uid: string;
    name: string;
}
