<script lang="ts">
	import { userStore } from '$lib/stores/user';
</script>

{#if $userStore}
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
