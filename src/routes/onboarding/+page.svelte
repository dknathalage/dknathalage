<script lang="ts">
	import { currentStep, steps } from './components/onboarding.store';
</script>

<div class="mt-5 flex flex-1 justify-between">
	<a class="btn btn-primary ml-1 text-xl" href={steps[$currentStep + 1]}> Let's Get Started </a>
</div>
