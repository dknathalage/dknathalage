<script lang="ts">
	import { authUser } from '$lib/stores/user';
	import { signin, signout } from '$lib/firebase/client/auth';
	import { GoogleAuthProvider } from 'firebase/auth';
	import { auth } from '$lib/firebase/client/config';

	async function handleSignIn() {
		await signin(new GoogleAuthProvider());
		window.location.href = '/app';
	}
</script>

{#if $authUser}
	<h1 class="text-2xl">
		Welcome {$authUser.displayName}
	</h1>

	<a href="/app" class="btn btn-primary m-1"> Home </a>
	<button on:click={async () => signout(auth)} class="btn btn-primary btn-outline m-1">
		Logout
	</button>
{:else}
	<button on:click={handleSignIn} class="btn btn-primary"> Google Signin </button>
{/if}
