<script lang="ts">
	export let disabled = false;
</script>

<button on:click {disabled}><slot /></button>

<style lang="postcss">
	button {
		@apply border border-black bg-gray-200 px-2 py-1;
	}
	button:disabled {
		@apply border-gray-400 bg-gray-100 text-gray-400 cursor-not-allowed;
	}
</style>
