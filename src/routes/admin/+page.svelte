<script lang="ts">
	import GenericControls from './GenericControls.svelte';
	import PhaseControls from './PhaseControls.svelte';
</script>

<svelte:head>
	<title>Admin Interface</title>
</svelte:head>

<main>
	<div class="preview">
		<p class="label">Live</p>
		<iframe src="./" title="Preview of page being shown to audience" />
	</div>
	<div class="divider" />
	<div class="admin">
		<p class="label">Controls</p>
		<GenericControls />
		<hr />
		<PhaseControls />
	</div>
</main>

<style lang="postcss">
	main {
		@apply text-xl text-center flex p-2 gap-2;
	}
	.preview {
		@apply w-[30vw] flex flex-col text-center gap-2 h-full place-items-stretch;
		min-height: calc(100vh - 1rem);
	}
	.divider {
		@apply border-l border-black;
	}
	.label {
		@apply text-3xl;
	}
	iframe {
		@apply h-full grow;
	}
	.admin {
		@apply grow flex flex-col items-center;
	}
	hr {
		@apply w-full border-black m-2;
	}
</style>
