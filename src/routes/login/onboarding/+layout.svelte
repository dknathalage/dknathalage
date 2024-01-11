<script lang="ts">
	import { page } from '$app/stores';
	import AnimatedRoute from '$lib/components/AnimatedRoute.svelte';
	import { writable } from 'svelte/store';

	const steps = ['/login/onboarding', '/login/onboarding/features'];
	const currentStep = writable(0);

	$: {
		const currentPageId = $page.route.id || steps[0];
		const currentIndex = steps.indexOf(currentPageId);
		currentStep.set(currentIndex);
	}
</script>

<nav class="flex justify-center my-6">
	<ul class="steps">
		<a href="/login/onboarding" class="step" class:step-primary={$currentStep >= 0}>Business Name</a
		>
		<a href="/login/onboarding/features" class="step" class:step-primary={$currentStep >= 1}>
			Choose Username
		</a>
		<!-- <a href="/login/onboard/photo" class="step" class:step-primary={$page.route.id?.includes('photo')}>
			Upload Photo
		</a> -->
	</ul>
</nav>

<AnimatedRoute>
	<div class="bg-neutral text-neutral-content mx-auto">
		<div class="items-center text-center">
			<slot />
		</div>
	</div>
	<div class="mt-5 flex flex-1 justify-between">
		<a class="btn btn-secondary btn-outline mr-1" href={steps[$currentStep - 1]}> Previous </a>
		<a class="btn btn-primary ml-1" href={steps[$currentStep + 1]}> Next </a>
	</div>
</AnimatedRoute>
