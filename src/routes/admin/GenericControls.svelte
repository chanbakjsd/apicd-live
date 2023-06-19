<script lang="ts">
	import Button from '../../components/Button.svelte';
	import { state, ws } from '../../lib/apicd';

	let sideA: HTMLInputElement;
	let sideB: HTMLInputElement;
	let title: HTMLInputElement;

	const changeText = (cmd: string, e: HTMLInputElement) => () => {
		$ws.send(cmd + ' ' + e.value);
	};

	const send = (cmd: string) => () => {
		$ws.send(cmd);
	};
</script>

<table>
	<tr>
		<td>Title</td>
		<td>
			<input bind:this={title} value={$state.title} />
			<Button on:click={changeText('title', title)}>Change</Button>
		</td>
	</tr>
	<tr>
		<td>Side A</td>
		<td>
			<input bind:this={sideA} value={$state.side_a} />
			<Button on:click={changeText('sideA', sideA)}>Change</Button>
		</td>
	</tr>
	<tr>
		<td>Side B</td>
		<td>
			<input bind:this={sideB} value={$state.side_b} />
			<Button on:click={changeText('sideB', sideB)}>Change</Button>
		</td>
	</tr>
	<tr>
		<td>Timer</td>
		<td>
			<Button on:click={send('togglePause')}>Pause/Resume</Button>
			<Button on:click={send('switchSide')} disabled={$state.timer.length <= 1}>Switch Side</Button>
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
