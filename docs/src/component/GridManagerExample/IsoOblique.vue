<template>
	<div class="gex-isobl">
		<shima-grid-manager
			:xGrid="xGrid" :xDiv="xDiv" :xGutter="xGutter"
			:yGrid="yGrid" :yDiv="yDiv" :yGutter="yGutter"
			:t="type"
		>
			<svg class="gex-isobl__svg" version="1.1" xmlns="http://www.w3.org/2000/svg">
				<g :transform="transform">
					<g :transform="`translate(${(xGrid+xGutter) * 2},${-(xGrid+yGutter) * 1})`">
						<rect  x="0"   y="0"  :width="w" :height="h" class="gex-layer1" />
						<rect :x="w"   y="0"  :width="w" :height="h" class="gex-layer1" />
						<rect :x="w2"  y="0"  :width="w" :height="h" class="gex-layer1" />
						<rect :x="w3"  y="0"  :width="w" :height="h" class="gex-layer1" />
						<rect  x="0"  :y="h"  :width="w" :height="h" class="gex-layer2" />
						<g :transform="offset">
							<rect  x="0"  :y="h2" :width="w" :height="h" class="gex-layer3" />
							<rect :x="w"  :y="h2" :width="w" :height="h" class="gex-layer3" />
							<rect :x="w2" :y="h2" :width="w" :height="h" class="gex-layer3" />
						</g>
						<rect :x="w3" :y="h3" :width="w" :height="h" class="gex-layer2" />
						<rect  x="0"  :y="h4" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w"  :y="h4" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w2" :y="h4" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w3" :y="h4" :width="w" :height="h" class="gex-layer1" />
					</g>

					<g :transform="`translate(${(xGrid+xGutter) * 3},${-(xGrid+yGutter) * 1})`">
						<rect x="0"   y="0"  :width="w" :height="h" class="gex-layer1" />
						<rect x="0"  :y="h"  :width="w" :height="h" class="gex-layer1" />
						<rect x="0"  :y="h3" :width="w" :height="h" class="gex-layer1" />
						<rect x="0"  :y="h4" :width="w" :height="h" class="gex-layer1" />
						<g :transform="offset">
							<rect  x="0"  :y="h2" :width="w" :height="h" class="gex-layer3" />
							<rect :x="w"  :y="h2" :width="w" :height="h" class="gex-layer3" />
							<rect :x="w2" :y="h2" :width="w" :height="h" class="gex-layer3" />
						</g>
						<rect :x="w3"  y="0"  :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h"  :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h3" :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h4" :width="w" :height="h" class="gex-layer2" />
					</g>

					<g :transform="`translate(${(xGrid+xGutter) * 3},${-(xGrid+yGutter) * 2})`">
						<rect :x="0 " :y="h2" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w1" :y="h2" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w2" :y="h2" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w3" :y="h2" :width="w" :height="h" class="gex-layer1" />
					</g>

					<g :transform="`translate(${(xGrid+xGutter) * 4},${-(xGrid+yGutter) * 2})`">
						<rect x="0"  y="0"  :width="w" :height="h" class="gex-layer1" />
						<rect x="0" :y="h2" :width="w" :height="h" class="gex-layer1" />
						<rect x="0" :y="h3" :width="w" :height="h" class="gex-layer1" />
						<rect x="0" :y="h4" :width="w" :height="h" class="gex-layer1" />
						<g :transform="offset">
							<rect x="0"   :y="h"  :width="w" :height="h" class="gex-layer3" />
							<rect :x="w"  :y="h"  :width="w" :height="h" class="gex-layer3" />
							<rect :x="w2" :y="h"  :width="w" :height="h" class="gex-layer3" />
							<rect :x="w"  :y="h2" :width="w" :height="h" class="gex-layer3" />
						</g>
						<rect :x="w3"  y="0"  :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h2" :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h3" :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h4" :width="w" :height="h" class="gex-layer2" />
					</g>

					<g :transform="`translate(${(xGrid+xGutter) * 4},-${(xGrid+yGutter) * 3})`">
						<rect  x="0" :y="h"  :width="w" :height="h" class="gex-layer1" />
						<rect  x="0" :y="h2" :width="w" :height="h" class="gex-layer1" />
						<rect  x="0" :y="h3" :width="w" :height="h" class="gex-layer1" />
						<rect  x="0" :y="h4" :width="w" :height="h" class="gex-layer1" />
						<rect :x="w" :y="h2" :width="w" :height="h" class="gex-layer1" />
						<g :transform="offset">
							<rect  x="0"  y="0" :width="w" :height="h" class="gex-layer3" />
							<rect :x="w1" y="0" :width="w" :height="h" class="gex-layer3" />
							<rect :x="w2" y="0" :width="w" :height="h" class="gex-layer3" />
						</g>
						<rect :x="w3" :y="h"  :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h2" :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h3" :width="w" :height="h" class="gex-layer2" />
						<rect :x="w3" :y="h4" :width="w" :height="h" class="gex-layer2" />
						<rect :x="w2" :y="h2" :width="w" :height="h" class="gex-layer2" />
					</g>
				</g>
			</svg>
		</shima-grid-manager>
	</div>
</template>
<style lang="stylus">
@import "~style/settings"
@import "~style/mixins"


.gex-isobl
	&__svg
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
import { _ } from "src/util";
import ShimaGridManager from "src/component/GridManager/GridManager.vue";


export default {

	props: {
		type: { type: String, default: "isometric", validator: (value) => _.includes(["isometric", "oblique"], value) }
	},

	data() {
		return {
			xGrid: 100, xDiv: 4, xGutter: 16,
			yGrid: 100, yDiv: 5, yGutter: 16
		};
	},

	computed: {
		w() { return this.xGrid / this.xDiv; },
		h() { return this.yGrid / this.yDiv; },
		w1() { return this.w * 1; }, h1() { return this.h * 1; },
		w2() { return this.w * 2; }, h2() { return this.h * 2; },
		w3() { return this.w * 3; }, h3() { return this.h * 3; },
		w4() { return this.w * 4; }, h4() { return this.h * 4; },
		offset() { return `translate(${this.w / 2},0)`; },
		transform() {
			return this.type === "isometric"
				? "matrix(1,0.577,-1.732,1,0,0) scale(0.866,0.500)"
				: "matrix(1,1.000,-1.000,1,0,0) scale(0.707,0.707)"
			;
		}
	},

	components: {
		ShimaGridManager
	}

};
</script>
