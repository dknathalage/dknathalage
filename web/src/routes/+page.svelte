<script lang="ts">
	import { signin, signout } from '$lib/firebase/client/authentication';
	import { user } from '$lib/stores/user';
	import { create } from '$lib/payments/session';

	async function createSub() {
		await create($user?.uid);
	}
</script>

<div class="h-screen flex items-center justify-center">
	{#if $user}
		<div class="flex flex-col items-center">
			<img src={$user?.photoURL} alt="logo" class="w-16 h-16 rounded" />
			<div class="m-1">
				<button class="btn btn-primary m-1" on:click={createSub}>Personal</button>
				<button class="btn btn-primary m-1">Business</button>
				<button class="btn btn-primary m-1">Organisation</button>
			</div>
			<button on:click={signout} class="btn btn-outline m-1"
				>Sign Out from {$user?.displayName}</button
			>
		</div>
	{:else}
		<button on:click={signin} class="btn btn-primary">Signin</button>
	{/if}
</div>
