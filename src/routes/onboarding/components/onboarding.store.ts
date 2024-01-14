import { writable } from 'svelte/store';
import type { Business } from '$lib/models';

export const businessOnboardingStore = writable<Business>({ name: '', uid: '' });
export const steps = ['/onboarding/name', '/onboarding/features', '/onboarding/submit', '/onboarding/done'];
export const currentStep = writable(0);