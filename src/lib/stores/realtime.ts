import { auth, db } from '$lib/firebase/client/config';
import { readable, writable } from 'svelte/store';
import { doc, onSnapshot } from 'firebase/firestore';

/**
 * @param  {string} path document path or reference
 * @returns a store with realtime updates on document data
 */
export function RtStore<T>(
    path: string,
  ) {
    let unsubscribe: () => void;
  
    const docRef = doc(db, path);
  
    const { subscribe } = writable<T | null>(null, (set) => {
      unsubscribe = onSnapshot(docRef, (snapshot) => {
        set((snapshot.data() as T) ?? null);
      });
  
      return () => unsubscribe();
    });
  
    return {
      subscribe,
      ref: docRef,
      id: docRef.id,
    };
  }