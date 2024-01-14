<script lang="ts">
	import OnboardActions from '../components/OnboardActions.svelte';
	import { createBusiness } from '$lib/stores/business';
	import { businessOnboardingStore } from '../components/onboarding.store';
	import { authUser } from '$lib/stores/user';

	async function handleCreateBusiness() {
		await createBusiness($businessOnboardingStore, $authUser);
	}
</script>

<button class="btn btn-primary" on:click={handleCreateBusiness}> Submit </button>
<OnboardActions />
