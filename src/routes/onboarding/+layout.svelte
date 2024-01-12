<script lang="ts">
	import { page } from '$app/stores';
	import AnimatedRoute from '$lib/components/AnimatedRoute.svelte';
	import { currentStep, steps } from './components/onboarding.store';

	$: {
		const currentIndex = steps.indexOf($page.route.id || steps[0]);
		currentStep.set(currentIndex);
	}
</script>

<nav class="flex justify-center my-6">
	<ul class="steps">
		{#each steps as step, index}
			<a href={step} class="step" class:step-primary={$currentStep >= index}> </a>
		{/each}
	</ul>
</nav>

<AnimatedRoute>
	<div class="mx-auto card">
		<div class="items-center text-center p-4 card-body">
			<slot />
		</div>
	</div>
</AnimatedRoute>
