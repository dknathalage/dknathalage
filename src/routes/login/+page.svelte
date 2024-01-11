<script lang="ts">
	import { authUser } from '$lib/stores/user';
	import { signin } from '$lib/firebase/client/auth';
	import { GoogleAuthProvider } from 'firebase/auth';
	import { business } from '$lib/stores/business';
	import { afterUpdate } from 'svelte';

	if ($authUser) {
		if ($business) {
			afterUpdate(() => {
				window.location.href = '/app';
			});
		}
	}

	async function handleSignIn() {
		await signin(new GoogleAuthProvider());
	}
</script>

{#if $authUser}
	<div class="justify-normal md:justify-center">
		<h1 class="text-4xl">
			Welcome
			<span class="text-accent font-semibold">{$authUser.displayName}</span>
		</h1>
		{#if !$business}
			<br />
			<a href="/login/onboarding" class="link link-primary text-4xl font-black">
				Click to get started.
			</a>
		{/if}
	</div>
{:else}
	<button on:click={handleSignIn} class="btn btn-primary"> Google Signin </button>
{/if}
