<script lang="ts">
	import Button from '../../components/Button.svelte';
	import { state } from '../../lib/apicd';
</script>

<table>
	<tr>
		<td>Side A</td>
		<td><input value={$state.side_a} /></td>
	</tr>
	<tr>
		<td>Side B</td>
		<td><input value={$state.side_b} /></td>
	</tr>
	<tr>
		<td>Timer</td>
		<td>
			<Button>Pause/Resume</Button>
			<Button disabled={$state.timer.length <= 1}>Switch Side</Button>
		</td>
	</tr>
</table>

<style lang="postcss">
	td {
		@apply px-2 py-1;
	}
	input {
		@apply border border-black;
	}
</style>
