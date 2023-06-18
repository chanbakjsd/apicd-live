<script lang="ts">
	import Button from '../../components/Button.svelte';
	import TimeInput from '../../components/TimeInput.svelte';
	import { config } from '../../lib/apicd';
</script>

<table>
	{#each $config.phases as phase}
		<tr>
			<td>{phase.text}</td>
			<td>
				{#each phase.time as time}
					<TimeInput bind:value={time} />
				{/each}
				<Button>Go</Button>
			</td>
		</tr>
	{/each}
</table>

<style lang="postcss">
	table tr {
		@apply odd:bg-yellow-100;
	}
	td {
		@apply px-2 py-1;
	}
</style>
