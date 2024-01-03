<script lang="ts">
	import { signin, signout } from '$lib/firebase/client/authentication';
	import { user } from '$lib/stores/user';
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<button on:click={signin} class="btn btn-primary">Signin</button>
<button on:click={signout} class="btn btn-secondary">Signout</button>
{#if $user}
	<p>Logged in as {$user.email}</p>
{:else}
	<p>Not logged in</p>
{/if}
