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
	const url = location.protocol === 'https:' ? `wss://${location.host}/ws` : `ws://${location.host}/ws`;
	const websocket = new WebSocket(url);
	websocket.onopen = () => {
		connected.set(true);
	};
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
		const data = msg.data as string;
		if (data.indexOf(" ") === -1) {
			console.warn('expected space in event', data);
			return;
		}
		const cmd = data.slice(0, data.indexOf(" "))
		const content = data.slice(data.indexOf(" "))
		switch (cmd) {
			case 'config':
				config.set(JSON.parse(content));
				break;
			case 'state':
				state.set(JSON.parse(content));
				break;
			default:
				console.warn('unknown message type', cmd);
		}
	};
	ws.set(websocket);
};
