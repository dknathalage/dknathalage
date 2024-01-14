<script lang="ts">
	import { GoogleAuthProvider } from 'firebase/auth';
	import { business } from '$lib/stores/business';
	import { afterUpdate } from 'svelte';
	import { user } from '$lib/repositories/user';

	if ($user && $business) {
		afterUpdate(() => {
			window.location.href = '/app';
		});
	}
</script>

{#if $user}
	<div class="justify-normal md:justify-center">
		<h1 class="text-4xl">
			Welcome
			<span class="text-accent font-semibold">{$user.displayName}</span>
		</h1>
		{#if !$business}
			<br />
			<a href="/onboarding" class="link link-primary text-4xl font-black">
				Click to get started.
			</a>
		{/if}
	</div>
	<button on:click={user.logOut} class="btn btn-primary"> Signout </button>
{:else}
	<button on:click={async () => await user.logIn(new GoogleAuthProvider())} class="btn btn-primary">
		Google Signin
	</button>
{/if}
