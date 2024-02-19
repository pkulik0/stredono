import Emotes from "./emotes.json";

const emitterPostions = [
	{
		x: -5,
		y: 10,
	},
	{
		x: 100,
		y: 80,
	},
	{
		x: 50,
		y: 0
	},
	{
		x: -5,
		y: 80,
	},
	{
		x: 100,
		y: 0,
	}
]

const emitter = (x: number, y: number, intensity: number) => ({
	life: {
		wait: false,
	},
	rate: {
		quantity: 1,
		delay: 3 / (intensity / 100) + Math.random() * 2,
	},
	particles: {
		opacity: {
			value: 0.5,
		},
		shape: {
			type: "images",
			options: {
				images: Emotes,
			},
		},
		size: {
			value: 28,
		},
		move: {
			speed: 8,
			outModes: {
				default: "none",
				right: "destroy",
			},
			straight: true,
		},
		zIndex: {
			value: 0,
		},
		rotate: {
			value: {
				min: 0,
				max: 360,
			},
			animation: {
				enable: true,
				speed: Math.random() * 6 + 5,
				sync: true,
			},
		},
	},
	position: {
		x: x,
		y: y,
	},
});

export const particlesConfig = (intensity: number) => ({
	name: "Stredono",
	particles: {
		number: {
			value: 85 * (intensity / 100),
		},
		links: {
			distance: 75,
			enable: true,
		},
		move: {
			enable: true,
		},
		size: {
			value: 1,
		},
		shape: {
			type: "circle",
		},
	},
	emitters: emitterPostions.map(({ x, y }) => emitter(x, y, intensity))
});