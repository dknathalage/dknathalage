import { writable } from 'svelte/store';
import type { Business } from '$lib/models';

export const businessInputStore = writable<Business | null>(null);