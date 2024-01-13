<script lang="ts">
	import OnboardActions from '../components/OnboardActions.svelte';
	import { createBusiness } from '$lib/stores/business';
	import { businessOnboardingStore } from '../components/onboarding.store';

	async function handleCreateBusiness() {
		const results = await createBusiness($businessOnboardingStore);
		console.log(results);
	}
</script>

<button class="btn btn-primary" on:click={handleCreateBusiness}> Submit </button>
<OnboardActions />
