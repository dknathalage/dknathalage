<script lang="ts">
	import { page } from '$app/stores';
	import AnimatedRoute from '$lib/components/AnimatedRoute.svelte';
</script>

<nav class="flex justify-center my-6">
	<ul class="steps">
		<a href="/login/onboard" class="step step-primary">Business Name</a>
		<a
			href="/login/onboard/features"
			class="step"
			class:step-primary={$page.route.id?.match(/photo|features/g)}
		>
			Choose Username
		</a>
		<!-- <a href="/login/onboard/photo" class="step" class:step-primary={$page.route.id?.includes('photo')}>
			Upload Photo
		</a> -->
	</ul>
</nav>

<AnimatedRoute>
	<main class="card w-4/6 bg-neutral text-neutral-content mx-auto">
		<div class="card-body items-center text-center">
			<slot />
		</div>
	</main>
</AnimatedRoute>
