<script lang="ts">
	import { signin, signout } from '$lib/firebase/client/authentication';
	import { user } from '$lib/stores/user';
	import { add } from '$lib/collections/user';

	$user;

	async function signIn() {
		await signin();
		await add($user);
	}

	async function signOut() {
		await signOut();
	}
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

{#if $user}
	<button on:click={signout} class="btn btn-secondary">Signout</button>
	<p>Logged in as {$user.email}</p>
{:else}
	<button on:click={signIn} class="btn btn-primary">Signin</button>

	<p>Not logged in</p>
{/if}
