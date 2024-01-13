<script lang="ts">
	import OnboardActions from '../components/OnboardActions.svelte';
	import { businessOnboardingStore } from '../components/onboarding.store';
</script>

<h2 class="text-3xl font-black mb-4">You're all set!</h2>

<span>{$businessOnboardingStore.name}</span>

<a class="btn btn-primary btn-xl" href="/app"> Continue to App </a>
<OnboardActions next={false} />
