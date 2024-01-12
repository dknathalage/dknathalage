<script lang="ts">
	import { currentStep, steps } from './onboarding.store';

	export let back = true;
	export let next = true;
</script>

{#if back || next}
	<div class="mt-3 flex flex-1 justify-between">
		{#if back}
			<a class="btn btn-secondary btn-sm btn-outline mr-1" href={steps[$currentStep - 1]}>
				<svg
					class="w-3 h-3"
					aria-hidden="true"
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 8 14"
				>
					<path
						stroke="currentColor"
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M7 1 1.3 6.326a.91.91 0 0 0 0 1.348L7 13"
					/>
				</svg>
				Back
			</a>
		{/if}
		{#if next}
			<a class="btn btn-primary btn-sm ml-1" href={steps[$currentStep + 1]}>
				Next
				<svg
					class="w-3 h-3"
					aria-hidden="true"
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 8 14"
				>
					<path
						stroke="currentColor"
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="m1 13 5.7-5.326a.909.909 0 0 0 0-1.348L1 1"
					/>
				</svg>
			</a>
		{/if}
	</div>
{/if}
