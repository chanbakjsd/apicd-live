<script lang="ts">
	import Button from '../../components/Button.svelte';
	import TimeInput from '../../components/TimeInput.svelte';
	import { config } from '../../lib/apicd';

	let phases = [
		{ text: '正方一辩立论', time: [210] },
		{ text: '反方一辩立论', time: [210] },
		{ text: '反方三辩质询正方一辩', time: [120] },
		{ text: '正方三辩质询反方一辩', time: [120] },
		{ text: '正方二辩驳论', time: [180] },
		{ text: '反方二辩驳论', time: [180] },
		{ text: '正方三辩申论', time: [180] },
		{ text: '反方三辩申论', time: [180] },
		{ text: '反方二辩盘问正方一、三、四辩', time: [150] },
		{ text: '正方二辩盘问反方一、三、四辩', time: [150] },
		{ text: '正方四辩小结', time: [90] },
		{ text: '反方四辩小结', time: [90] },
		{ text: '自由辩论', time: [240, 240] },
		{ text: '正方四辩总结陈词', time: [150] },
		{ text: '反方四辩总结陈词', time: [150] }
	];
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
