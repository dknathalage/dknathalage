<script lang="ts">
	import OnboardActions from '../components/OnboardActions.svelte';
	import { businessOnboardingStore } from '../components/onboarding.store';

	$: next = $businessOnboardingStore.name.length >= 3;
</script>

<div>
	<h2 class="text-3xl font-black mb-4">Name of the Business</h2>
	<input
		id="name"
		class="input input-bordered join-item"
		placeholder="Awesome Inc."
		bind:value={$businessOnboardingStore.name}
	/>
</div>
<span class="text-xs text-gray-500">
	This will be used as the name of your business on the platform.
</span>
{#if !($businessOnboardingStore.name.length >= 3)}
	<span class="text-xs text-error">Name must be at least 3 characters long.</span>
{/if}
<OnboardActions back={false} {next} />
