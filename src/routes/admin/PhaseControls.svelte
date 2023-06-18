<script lang="ts">
	import Button from '../../components/Button.svelte';
	import TimeInput from '../../components/TimeInput.svelte';
	import { config, ws } from '../../lib/apicd';

	const go = (i: number) => () => {
		$ws.send(`go ${i}${$config.phases[i].time.map((x) => ` ${x}`)}`);
	};
</script>

<table>
	{#each $config.phases as phase, i}
		<tr>
			<td>{phase.text}</td>
			<td>
				{#each phase.time as time}
					<TimeInput bind:value={time} />
				{/each}
				<Button on:click={go(i)}>Go</Button>
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
