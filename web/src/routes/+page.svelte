<script lang="ts">
	import { signin, signout } from '$lib/firebase/client/authentication';
	import { user } from '$lib/stores/user';
	import { add } from '$lib/collections/user';
	import { create} from '$lib/payments/session';

	$user;

	async function signIn() {
		await signin();
		add($user);
	}

	async function signOut() {
		await signOut();
	}

	async function createSub() {
		console.log($user);
		await create($user.uid);
	}
</script>

<div class="h-screen flex items-center justify-center">
{#if $user}
	<div class="flex flex-col items-center">
		<div class="m-1">
			<button class="btn btn-primary m-1" on:click={createSub}>Personal</button>
			<button class="btn btn-primary m-1">Business</button>
			<button class="btn btn-primary m-1">Organisation</button>
		</div>
		<button on:click={signout} class="btn btn-outline m-1">Sign Out from {$user.email}</button>
	</div>
{:else}
	<button on:click={signIn} class="btn btn-primary">Signin</button>
{/if}
</div>
