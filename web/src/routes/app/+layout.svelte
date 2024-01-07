<script lang="ts">
	import { signout } from '$lib/firebase/client/auth';
	import { auth } from '$lib/firebase/client/config';
	import Authenticated from '$lib/components/Authenticated.svelte';
	import { writable } from 'svelte/store';

	const active = writable('Home');
</script>

<Authenticated>
	<div class="flex-1 navbar bg-base-100 fixed sm:relative bottom-0 sm:bottom-auto justify-center">
		<div class="md:flex-1 hidden sm:block">
			<a class="btn btn-ghost text-lg" on:click={() => active.set('Home')} href="/app">
				[Business Name] / {$active}
			</a>
		</div>
		<div class="items-center">
			<ul class="menu bg-base-300 menu-horizontal rounded-box shadow-xl">
				<li>
					<a class:active={$active === 'Home'} on:click={() => active.set('Home')} href="/app">
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 20 20"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M3 8v10a1 1 0 0 0 1 1h4v-5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5h4a1 1 0 0 0 1-1V8M1 10l9-9 9 9"
							/>
						</svg>
						<div class="hidden md:block">
							Home
							<!-- <span class="badge badge-sm">99+</span> -->
						</div>
					</a>
				</li>
				<li>
					<a
						class:active={$active === 'People'}
						on:click={() => active.set('People')}
						href="/app/people"
					>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 20 20"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M4.333 6.764a3 3 0 1 1 3.141-5.023M2.5 16H1v-2a4 4 0 0 1 4-4m7.379-8.121a3 3 0 1 1 2.976 5M15 10a4 4 0 0 1 4 4v2h-1.761M13 7a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm-4 6h2a4 4 0 0 1 4 4v2H5v-2a4 4 0 0 1 4-4Z"
							/>
						</svg>
						<div class="hidden md:block">People</div>
					</a>
				</li>
				<li>
					<a
						class:active={$active === 'Transactions'}
						on:click={() => active.set('Transactions')}
						href="/app/transactions"
					>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="currentColor"
							viewBox="0 0 19 20"
						>
							<path
								d="M18.972.863a.913.913 0 0 0-.041-.207.956.956 0 0 0-.107-.19 1.01 1.01 0 0 0-.065-.116c-.008-.01-.02-.013-.028-.022a1.008 1.008 0 0 0-.174-.137 1.085 1.085 0 0 0-.141-.095 1.051 1.051 0 0 0-.171-.047.985.985 0 0 0-.207-.041C18.025.007 18.014 0 18 0h-3.207a1 1 0 1 0 0 2h.5l-4.552 3.9-3.5-.874a1 1 0 0 0-.867.189l-5 4a1 1 0 0 0 1.25 1.562L7.238 7.09l3.52.88a1 1 0 0 0 .892-.211L17 3.173v1.034a1 1 0 0 0 2 0V1a.9.9 0 0 0-.028-.137ZM13.5 9a5.5 5.5 0 1 0 0 11 5.5 5.5 0 0 0 0-11Zm.24 4.591a3.112 3.112 0 0 1 1.935 1.374 2.036 2.036 0 0 1 .234 1.584 2.255 2.255 0 0 1-1.374 1.469.982.982 0 0 1-1.953.09 2.943 2.943 0 0 1-1.475-.92 1 1 0 0 1 1.536-1.283.953.953 0 0 0 .507.29.778.778 0 0 0 .831-.18 1.108 1.108 0 0 0-.714-.481 3.105 3.105 0 0 1-1.934-1.374 2.042 2.042 0 0 1-.233-1.584 2.264 2.264 0 0 1 1.45-1.493v-.03a1 1 0 0 1 2 0c.517.159.98.457 1.337.862a1.002 1.002 0 1 1-1.524 1.3.962.962 0 0 0-.507-.286.775.775 0 0 0-.829.18 1.113 1.113 0 0 0 .713.482ZM6 20a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1Zm-4 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1Z"
							/>
						</svg>
						<div class="hidden md:block">
							Transactions
							<!-- <span class="badge badge-sm badge-success">+99</span> -->
						</div>
					</a>
				</li>
				<li>
					<a
						class:active={$active === 'Notifications'}
						on:click={() => active.set('Notifications')}
						href="/app/notifications"
					>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 16 21"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M8 3.464V1.1m0 2.365a5.338 5.338 0 0 1 5.133 5.368v1.8c0 2.386 1.867 2.982 1.867 4.175C15 15.4 15 16 14.462 16H1.538C1 16 1 15.4 1 14.807c0-1.193 1.867-1.789 1.867-4.175v-1.8A5.338 5.338 0 0 1 8 3.464ZM4.54 16a3.48 3.48 0 0 0 6.92 0H4.54Z"
							/>
						</svg>
						<div class="hidden md:block">
							Notifications
							<!-- <span class="badge badge-sm badge-warning">NEW</span> -->
						</div>
					</a>
				</li>
				<li>
					<a
						class:active={$active === 'Preferences'}
						on:click={() => active.set('Preferences')}
						href="/app/preferences"
					>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 20 20"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M7.75 4H19M7.75 4a2.25 2.25 0 0 1-4.5 0m4.5 0a2.25 2.25 0 0 0-4.5 0M1 4h2.25m13.5 6H19m-2.25 0a2.25 2.25 0 0 1-4.5 0m4.5 0a2.25 2.25 0 0 0-4.5 0M1 10h11.25m-4.5 6H19M7.75 16a2.25 2.25 0 0 1-4.5 0m4.5 0a2.25 2.25 0 0 0-4.5 0M1 16h2.25"
							/>
						</svg>
						<div class="hidden md:block">Preferences</div>
					</a>
				</li>
			</ul>
			<ul class="menu bg-base-300 lg:menu-horizontal rounded-box ml-2 hidden md:block shadow-xl">
				<li>
					<button
						on:click={() => {
							signout(auth);
						}}
					>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 18 15"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M1 7.5h11m0 0L8 3.786M12 7.5l-4 3.714M12 1h3c.53 0 1.04.196 1.414.544.375.348.586.82.586 1.313v9.286c0 .492-.21.965-.586 1.313A2.081 2.081 0 0 1 15 14h-3"
							/>
						</svg>
					</button>
				</li>
			</ul>
		</div>
	</div>
	<div class="m-5">
		<slot />
	</div>
</Authenticated>
