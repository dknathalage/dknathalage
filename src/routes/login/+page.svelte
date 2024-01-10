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
	<h1 class="text-4xl">
		Welcome {$authUser.displayName}
	</h1>
	{#if !$business}
		<br />
		<span class="text text-accent text-4xl"> Let's setup your business </span>
		<br />
		<a href="/login/onboard" class="link link-primary text-4xl"> Click to get started </a>
	{/if}
{:else}
	<button on:click={handleSignIn} class="btn btn-primary"> Google Signin </button>
{/if}
