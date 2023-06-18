<script lang="ts">
	import Text from '../components/Text.svelte';
	import { config, state } from '../lib/apicd';
	import { formatTime } from '../lib/time';

	let boldClass = 'text-3xl lg:text-6xl text-yellow-500 font-bold';
	let special: Record<string, string> = {};
	$: {
		special = {};
		special[$config.side_a] = 'bg-red-500 p-1';
		special[$config.side_b] = 'bg-blue-500 p-1';
	}
</script>

<main>
	<p><Text text={$state.phase} {special} /></p>
	<div class="topic">
		<p><Text text={$state.side_a} bold={boldClass} {special} /></p>
		<p class="side side-a">{$config.side_a}</p>
		<p>vs</p>
		<p class="side side-b">{$config.side_b}</p>
		<p><Text text={$state.side_b} bold={boldClass} {special} /></p>
	</div>
	<div class="countdown-container">
		{#each $state.timer as time, i}
			<p
				class="countdown"
				class:active={i === $state.active && !$state.paused}
				class:over={time === 0}
			>
				{formatTime(time)}
			</p>
		{/each}
	</div>
</main>

<style lang="postcss">
	:global(body) {
		@apply bg-center bg-cover text-white text-xl lg:text-5xl;
		background-image: url('/background.jpg');
	}

	main {
		@apply w-screen h-screen grid place-content-center text-center gap-16 pointer-events-none;
		@apply bg-black/50 lg:p-[10vw] xl:p-[20vw];
	}
	.topic {
		@apply flex gap-4 text-xl justify-center items-center w-full;
		@apply lg:gap-8 lg:text-4xl;
	}
	.side {
		@apply p-4 rounded-full shrink-0 text-2xl;
		@apply lg:p-8 lg:text-6xl;
	}
	.side-a {
		@apply bg-red-500;
	}
	.side-b {
		@apply bg-blue-500;
	}
	.countdown-container {
		@apply flex gap-6 justify-center;
		@apply lg:gap-16;
	}
	.countdown {
		@apply text-5xl lg:text-[12rem] text-gray-500;
	}
	.countdown.active {
		@apply text-white;
	}
	.countdown.over {
		@apply text-red-800;
	}
	.countdown.over.active {
		@apply text-red-500;
	}
</style>
