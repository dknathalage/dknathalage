<script lang="ts">
	import { signout } from '$lib/firebase/client/auth';
	import { auth } from '$lib/firebase/client/config';
	import Authenticated from '$lib/components/Authenticated.svelte';
</script>

<Authenticated>
	<div class="navbar bg-base-100">
		<div class="flex-1">
			<a class="btn btn-ghost text-xl">daisyUI</a>
		</div>
		<div class="flex-none">
			<button class="btn btn-square btn-ghost" on:click={async () => await signout(auth)}>
				Logout
			</button>
		</div>
	</div>
	<slot />
</Authenticated>
