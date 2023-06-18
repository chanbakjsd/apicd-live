const leftPad = (x: string, length: number): string => {
	while (x.length < length) {
		x = `0${x}`;
	}
	return x;
};
export const formatTime = (seconds: number): string => {
	const minute = leftPad(`${Math.floor(seconds / 60)}`, 2);
	const second = leftPad(`${seconds % 60}`, 2);
	return `${minute}:${second}`;
};
