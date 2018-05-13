<template>
	<div class="logo">
		<svg :viewBox="`0 0 ${(xGrid+xGutter) * 5 - 3 * w} ${yGrid}`" class="logo__svg" version="1.1" xmlns="http://www.w3.org/2000/svg">
			<g :transform="`translate(0,0)`">
				<cube :x="0"  :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w"  :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w2" :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0"  :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<g :transform="offset">
					<cube :x="0"  :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w"  :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w2" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				</g>
				<cube :x="w3" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0"  :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w"  :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w2" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
			</g>

			<g :transform="`translate(${(xGrid+xGutter) * 1},0)`">
				<cube :x="0"  :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0"  :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0"  :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0"  :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<g :transform="offset">
					<cube :x="0"  :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w"  :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w2" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				</g>
				<cube :x="w3" :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
			</g>

			<g :transform="`translate(${(xGrid+xGutter) * 2},0)`">
				<cube :x="0" :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h1" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
			</g>

			<g :transform="`translate(${(xGrid+xGutter) * 3 - 3 * w},0)`">
				<cube :x="0" :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<g :transform="offset">
					<cube :x="0"  :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w"  :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w2" :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w"  :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				</g>
				<cube :x="w3" :y="0"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
			</g>

			<g :transform="`translate(${(xGrid+xGutter) * 4 - 3 * w},0)`">
				<cube :x="0" :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="0" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<g :transform="offset">
					<cube :x="0"  :y="0" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w1" :y="0" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
					<cube :x="w2" :y="0" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				</g>
				<cube :x="w3" :y="h"  :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h3" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w3" :y="h4" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
				<cube :x="w2" :y="h2" :width="w" :height="h" :ee="ee" logo @deleted="inc" />
			</g>
		</svg>
	</div>
</template>
<style lang="stylus">
@import "~style/settings"
@import "~style/mixins"


.logo
	&__svg
		width 100%
		height 100%

		.cube__main
			transition all .3s ease-in-out

		//1, 4, ...
		svg:nth-child(3n+1) rect.cube__main
			stroke: alpha($base-color["api"], .8)
			fill: alpha($base-color["api"], .8)

		// 2, 5, ...
		svg:nth-child(3n+2) rect.cube__main
			stroke: alpha($base-color["api"],  1)
			fill: alpha($base-color["api"],  1)

		// 3, 6, ...
		svg:nth-child(3n) rect.cube__main
			stroke: alpha($base-color["api"], .5)
			fill: alpha($base-color["api"], .5)

		svg:hover rect.cube__main
			stroke: none
			fill: alpha(white, .5)

	rect
		fill yellow
		//stroke red
		//stroke-width 2px

	&__svg2
		display block
		absolute top right bottom left
		width 100%
		height 100%

		.gex-layer1
			fill red
			fill: $base-color["api"]

		.gex-layer2
			fill green
			fill: $base-color["api"]

		.gex-layer3
			fill blue
			fill: $base-color["default"]

</style>
<script>
import Cube from "src/component/Cube.vue";


export default {

	components: {
		Cube
	},

	props: {
		toggleEe: { type: Boolean, default: false }
	},

	data() {
		return {
			xGrid: 128, xDiv: 4, xGutter: 16,
			yGrid: 160, yDiv: 5, yGutter: 16,
			cube: "Cube",
			ee: this.forceEe,
			nbCubeDestroyed: 0,
			totalCubes: (13 + 11 + 5 + 12 + 13)
		};
	},

	computed: {
		w() { return this.xGrid / this.xDiv; },
		h() { return this.yGrid / this.yDiv; },
		w1() { return this.w * 1; }, h1() { return this.h * 1; },
		w2() { return this.w * 2; }, h2() { return this.h * 2; },
		w3() { return this.w * 3; }, h3() { return this.h * 3; },
		w4() { return this.w * 4; }, h4() { return this.h * 4; },
		offset() { return `translate(${this.w / 2},0)`; }
	},

	watch: {
		toggleEe: function () {
			this.ee = !this.ee;
		}
	},

	methods: {
		inc() {
			++this.nbCubeDestroyed;
			if (this.nbCubeDestroyed >= this.totalCubes) {
				this.nbCubeDestroyed = 0;
				this.ee = !this.ee;
			}
		}
	}

};
</script>
