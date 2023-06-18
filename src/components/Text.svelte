<script lang="ts">
	export let text: string;
	export let bold = '';
	export let special: Record<string, string> = {};

	let displays: Display[];
	interface Display {
		content: string;
		class: string;
	}

	const processSpecial = (text: Display, toCheck: string[]): Display[] => {
		for (let i = 0; i < toCheck.length; i++) {
			const split = text.content.split(toCheck[i]);
			if (split.length == 1) {
				continue;
			}
			const remaining = toCheck.slice(i + 1);
			const result: Display[] = [];
			for (let j = 0; j < split.length; j++) {
				const segment = {
					content: split[j],
					class: text.class
				};
				result.push(...processSpecial(segment, remaining));
				if (j !== split.length - 1) {
					result.push({
						content: toCheck[i],
						class: text.class + ' ' + special[toCheck[i]]
					});
				}
			}
			return result;
		}
		// No matches with special.
		return [text];
	};

	$: sortedKeys = Object.keys(special).sort((a, b) => {
		if (a.length - b.length != 0) return a.length - b.length;
		return a < b ? -1 : 1;
	});
	$: {
		let isBold = true;
		displays = [];
		for (let entry of text.split('*')) {
			isBold = !isBold;
			displays.push(
				...processSpecial(
					{
						content: entry,
						class: isBold ? bold : ''
					},
					sortedKeys
				)
			);
		}
	}
</script>

<p>
	{#each displays as display}
		<span class={display.class}>{display.content}</span>
	{/each}
</p>
