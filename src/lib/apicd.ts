import { writable } from 'svelte/store';

export type Config = {
	side_a: string;
	side_b: string;
	phases: Phase[];
};

export type State = {
	title: string;
	phase: string;
	side_a: string;
	side_b: string;
	side_a_name: string;
	side_b_name: string;

	timer: number[];
	active: number;
	paused: boolean;
};

export type Phase = {
	label: string;
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
	title: '',
	phase: '',
	side_a: '',
	side_b: '',
	side_a_name: '',
	side_b_name: '',
	timer: [],
	active: -1,
	paused: true
});

export const connected = writable(false);

export const ws = writable<WebSocket>();

export const connect = () => {
	const url = new URL("./ws", location.href);
	url.protocol = location.protocol === "https:" ? "wss:" : "ws:";
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
