import { writable } from 'svelte/store';
import type { Business } from '$lib/models';

export const businessInputStore = writable<Business | null>(null);
export const steps = ['/onboarding/name', '/onboarding/features', '/onboarding/done'];
export const currentStep = writable(0);


