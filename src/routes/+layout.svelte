<script lang="ts">
	import { onMount } from 'svelte';
	import '../app.css';
	import { connect, connected } from '../lib/apicd';

	onMount(() => {
		connect();
	});
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="" />
	<link
		href="https://fonts.googleapis.com/css2?family=ZCOOL+XiaoWei&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<slot />
{#if !$connected}
	<div class="warning">⚠️</div>
{/if}

<style lang="postcss">
	.warning {
		@apply absolute bottom-2 right-2 text-5xl animate-bounce;
	}
</style>
