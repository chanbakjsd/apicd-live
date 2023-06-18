import { writable } from 'svelte/store';

export type Config = {
	side_a: string;
	side_b: string;
	phases: Phase[];
};

export type State = {
	phase: string;
	side_a: string;
	side_b: string;

	timer: number[];
	active: number;
	paused: boolean;
};

export type Phase = {
	text: string;
	time: number[];
};

export type Timer = {
	time: number;
	active: boolean;
};

export const config = writable<Config>({
	side_a: '',
	side_b: '',
	phases: []
});

export const state = writable<State>({
	phase: '',
	side_a: '',
	side_b: '',
	timer: [],
	active: -1,
	paused: true
});

export const connected = writable(false);

export const ws = writable<WebSocket>();

export const connect = () => {
	const url = location.protocol === 'https' ? `wss://${location.host}/ws` : `ws://${location.host}/ws`;
	const websocket = new WebSocket(url);
	websocket.onclose = () => {
		connected.set(false);
		console.warn('Websocket close');
		setTimeout(connect, 1000);
	};
	websocket.onerror = (err) => {
		connected.set(false);
		console.warn('Websocket error:', err);
		setTimeout(connect, 1000);
	};
	websocket.onmessage = (msg: MessageEvent) => {
		const decoded = (msg.data as string).split(' ', 2);
		if (decoded.length !== 2) {
			console.warn('expected space in event', decoded);
			return;
		}
		switch (decoded[0]) {
			case 'config':
				config.set(JSON.parse(decoded[1]));
				break;
			case 'state':
				state.set(JSON.parse(decoded[1]));
				break;
			default:
				console.warn('unknown message type', decoded[0]);
		}
	};
	ws.set(websocket);
};
