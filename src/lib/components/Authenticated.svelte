<script lang="ts">
	import { authUser } from '$lib/stores/user';
</script>

{#if $authUser}
	<slot />
{:else}
	<div class="h-screen flex items-center justify-center">
		<div class="flex flex-col items-center">
			<p>Authentication Required</p>
			Please login to continue
			<a href="/login" class="link">Login</a>
		</div>
	</div>
{/if}
