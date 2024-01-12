<script lang="ts">
	import { currentStep, steps } from './onboarding.store';
</script>

<div class="mt-3 flex flex-1 justify-between">
	<a class="btn btn-secondary btn-sm btn-outline mr-1" href={steps[$currentStep - 1]}> Back </a>
	<a class="btn btn-primary btn-sm ml-1" href={steps[$currentStep + 1]}> Next </a>
</div>

