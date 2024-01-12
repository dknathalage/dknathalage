<script lang="ts">
	import { page } from '$app/stores';
	import AnimatedRoute from '$lib/components/AnimatedRoute.svelte';
	import { currentStep, steps } from './components/onboarding.store';

	$: {
		const currentPageId = $page.route.id || steps[0];
		const currentIndex = steps.indexOf(currentPageId);
		currentStep.set(currentIndex);
	}
</script>

<nav class="flex justify-center my-6">
	<ul class="steps">
		<a href="/onboarding/name" class="step" class:step-primary={$currentStep >= 0}>Business Name</a>
		<a href="/onboarding/features" class="step" class:step-primary={$currentStep >= 1}>
			Choose Username
		</a>
	</ul>
</nav>

<AnimatedRoute>
	<div class="mx-auto card">
		<div class="items-center text-center p-4 card-body">
			<slot />
		</div>
	</div>
</AnimatedRoute>
