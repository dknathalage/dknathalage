<script lang="ts">
	import OnboardActions from '../components/OnboardActions.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h2 class="text-3xl font-black mb-4">You're all set!</h2>
<a class="btn btn-primary btn-xl" href="/app"> Continue to App </a>
<OnboardActions next={false} />
