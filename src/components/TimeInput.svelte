<script lang="ts">
	import { formatTime } from '../lib/time';

	export let value: number;
	let err = false;

	$: display = formatTime(value);

	const change = (e: Event) => {
		err = false;
		try {
			const text = (e.target as HTMLInputElement).value;
			const split = text.split(':');
			let val = -1;
			switch (split.length) {
				case 1:
					val = parseInt(split[0]);
					break;
				case 2:
					val = parseInt(split[0]) * 60 + parseInt(split[1]);
					break;
				default:
					break;
			}
			if (val <= 0) {
				err = true;
				return;
			}
			value = val;
		} catch (e) {
			console.log(e);
			err = true;
		}
	};
</script>

<input on:change={change} class:err value={display} />

<style lang="postcss">
	input {
		@apply border border-black w-16 mr-2 text-center;
	}
	input.err {
		@apply bg-red-200;
	}
</style>
