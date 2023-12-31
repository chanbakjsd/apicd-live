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

<svelte:head>
	<title>Audience View</title>
</svelte:head>

<main>
	<div class="title">
		<img src="logo.png" alt="Logo of APICD" />
		<p>{$state.title}</p>
	</div>
	<div class="teams">
		<p>{$state.side_a_name}</p>
		<p>{$state.side_b_name}</p>
	</div>
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
				class:over={time <= 30}
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
		@apply w-screen h-screen flex flex-col items-center justify-center text-center gap-8 pointer-events-none;
		@apply bg-black/50 lg:gap-10 xl:gap-12 2xl:gap-16 xl:px-[5vw] 2xl:px-[10vw];
	}
	.title {
		@apply flex flex-col items-center;
	}
	img {
		@apply w-24 h-24;
	}
	.teams {
		@apply flex gap-8 lg:gap-16 text-2xl xl:text-4xl 2xl:text-6xl;
	}
	.topic {
		@apply flex gap-4 text-xl justify-center items-center w-full;
		@apply lg:gap-6 lg:text-3xl 2xl:gap-8 2xl:text-4xl;
	}
	.side {
		@apply p-4 rounded-full shrink-0 text-2xl;
		@apply lg:p-6 lg:text-4xl 2xl:p-8 2xl:text-6xl;
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
		@apply text-4xl xl:text-9xl 2xl:text-[10rem] 2xl:leading-[12rem] text-gray-500;
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
