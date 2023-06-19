<script lang="ts">
	import Button from '../../components/Button.svelte';
	import { state, ws } from '../../lib/apicd';

	let title: HTMLInputElement;
	let sideA: HTMLInputElement;
	let sideB: HTMLInputElement;
	let sideAName: HTMLInputElement;
	let sideBName: HTMLInputElement;

	const changeText = (cmd: string, e: HTMLInputElement) => () => {
		$ws.send(cmd + ' ' + e.value);
	};

	const send = (cmd: string) => () => {
		$ws.send(cmd);
	};

	const changeTeamName = () => {
		changeText("sideAName", sideAName)()
		changeText("sideBName", sideBName)()
	}

	const changePrompt = () => {
		changeText("sideA", sideA)()
		changeText("sideB", sideB)()
	}
</script>

<table>
	<tr>
		<td>Title</td>
		<td>
			<input bind:this={title} placeholder="Title" value={$state.title} />
			<Button on:click={changeText('title', title)}>Change</Button>
		</td>
	</tr>
	<tr>
		<td>Team</td>
		<td>
			<div class="inputs">
			<div class="stack">
			<input bind:this={sideAName} placeholder="Side A" value={$state.side_a_name} />
			<input bind:this={sideBName} placeholder="Side B" value={$state.side_b_name} />
			</div>
			<Button on:click={changeTeamName}>Change</Button>
			</div>
		</td>
	</tr>
	<tr>
		<td>Prompt</td>
		<td>
			<div class="inputs">
			<div class="stack">
			<input bind:this={sideA} placeholder="Side A" value={$state.side_a} />
			<input bind:this={sideB} placeholder="Side B" value={$state.side_b} />
			</div>
			<Button on:click={changePrompt}>Change</Button>
			</div>
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
	.inputs {
		@apply flex gap-2;
	}
	.stack {
		@apply inline-block flex flex-col gap-1;
	}
	input {
		@apply border border-black;
	}
</style>
