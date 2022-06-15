<template>
	<transition :duration="300" name="cube--fade">
		<svg v-if="!goodBye" :class="{'cube--logo': logo }" class="cube" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
			<g class="cube__transition">
				<g v-if="logo">
					<rect
						:style="eeStyle" class="cube__main"
						x="0" y="0" width="32" height="32" rx="0" ry="0"
						@mouseenter="updateCube" @mouseleave="updateCube" @click.prevent="deleteCube"
					/>
				</g>
				<g v-else transform="translate(0.5, 0.5)">
					<rect class="cube__main" x="2" y="2" width="27" height="27" rx="2" ry="2" />
				</g>

				<g v-if="t >= 0" :transform="gr ? 'rotate(45, 16, 16)' : null">
					<rect v-if="d[0] === 1" :transform="p[0].t" :x="p[0].x" :y="p[0].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[1] === 1" :transform="p[1].t" :x="p[1].x" :y="p[1].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[2] === 1" :transform="p[2].t" :x="p[2].x" :y="p[2].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[3] === 1" :transform="p[3].t" :x="p[3].x" :y="p[3].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[4] === 1" :transform="p[4].t" :x="p[4].x" :y="p[4].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[5] === 1" :transform="p[5].t" :x="p[5].x" :y="p[5].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[6] === 1" :transform="p[6].t" :x="p[6].x" :y="p[6].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[7] === 1" :transform="p[7].t" :x="p[7].x" :y="p[7].y" width="4"  height="4" class="cube__mini" />
					<rect v-if="d[8] === 1" :transform="p[8].t" :x="p[8].x" :y="p[8].y" width="4"  height="4" class="cube__mini" />

					<rect v-if="d[0] === 2" :x="p[0].x" :y="p[0].y" :width="8 + 4 * 2"  height="4" class="cube__mini" />
					<rect v-if="d[3] === 2" :x="p[3].x" :y="p[3].y" :width="8 + 4 * 2"  height="4" class="cube__mini" />
					<rect v-if="d[6] === 2" :x="p[6].x" :y="p[6].y" :width="8 + 4 * 2"  height="4" class="cube__mini" />

					<rect v-if="d[0] === 3" :height="8 + 4 * 2" :x="p[0].x" :y="p[0].y" width="4" class="cube__mini" />
					<rect v-if="d[1] === 3" :height="8 + 4 * 2" :x="p[1].x" :y="p[1].y" width="4" class="cube__mini" />
					<rect v-if="d[2] === 3" :height="8 + 4 * 2" :x="p[2].x" :y="p[2].y" width="4" class="cube__mini" />
				</g>
			</g>
		</svg>
	</transition>
</template>
<script>
import { _ } from "src/util";


export default {

	props: {
		type: { type: Number,  default: null },
		logo: { type: Boolean, default: false },
		ee  : { type: Boolean, default: false }
	},

	emits: ["deleted"],

	data: function () {
		const layout = [
			{ pts: [1,1,1,1,1,1,1,1,1] },
			{ pts: [1,1,1,0,0,0,1,1,1] },
			{ pts: [1,0,1,1,0,1,1,0,1] },
			{ pts: [0,1,0,1,0,1,0,1,0] },
			{ pts: [1,0,1,0,0,0,1,0,1] },
			{ pts: [2,0,0,2,0,0,2,0,0] },
			{ pts: [3,3,3,0,0,0,0,0,0] },
			{ pts: [1,3,1,2,0,0,1,0,1] },
			{ pts: [1,3,1,1,0,1,1,0,1] },
			{ pts: [1,1,1,2,0,0,1,1,1] }
		];

		return {
			layout,
			// Select random shape on "mounted" event (Prevent "Hydration completed but contains mismatches.")
			// t: (this.type === null ? _.random(0, layout.length * 3) : this.type),
			t: (this.type === null ? -1 : this.type),
			goodBye: false
		};
	},

	computed: {
		d() {
			const type = this.t % this.layout.length;
			return this.layout[type].pts;
		},
		// global rotate
		gr() {
			return this.t / this.layout.length >= 1;
		},
		// local rotate
		lr() {
			return this.t / this.layout.length >= 2;
		},
		p() {
			const lr = this.lr;
			// offset, width, gutter
			const x = { o: 8, w: 4, g: 2 };
			const y = { o: 8, w: 4, g: 2 };
			x.wg = x.w + x.g; x.r = x.w / 2;
			y.wg = y.w + y.g; y.r = y.w / 2;

			// cells
			x.p1 = x.o;            y.p1 = y.o;
			x.p2 = x.o + x.wg;     y.p2 = y.o + y.wg;
			x.p3 = x.o + x.wg * 2; y.p3 = y.o + y.wg * 2;
			x.c1 = x.p1 + x.r;     y.c1 = y.p1 + y.r;
			x.c2 = x.p2 + x.r;     y.c2 = y.p2 + y.r;
			x.c3 = x.p3 + x.r;     y.c3 = y.p3 + y.r;

			return [
				{ x: x.p1, y: y.p1, cx: x.c1, cy: y.c1, t: (lr ? `rotate(45, ${x.c1}, ${y.c1})` : null) },
				{ x: x.p2, y: y.p1, cx: x.c2, cy: y.c1, t: (lr ? `rotate(45, ${x.c2}, ${y.c1})` : null) },
				{ x: x.p3, y: y.p1, cx: x.c3, cy: y.c1, t: (lr ? `rotate(45, ${x.c3}, ${y.c1})` : null) },
				{ x: x.p1, y: y.p2, cx: x.c1, cy: y.c2, t: (lr ? `rotate(45, ${x.c1}, ${y.c2})` : null) },
				{ x: x.p2, y: y.p2, cx: x.c2, cy: y.c2, t: (lr ? `rotate(45, ${x.c2}, ${y.c2})` : null) },
				{ x: x.p3, y: y.p2, cx: x.c3, cy: y.c2, t: (lr ? `rotate(45, ${x.c3}, ${y.c2})` : null) },
				{ x: x.p1, y: y.p3, cx: x.c1, cy: y.c3, t: (lr ? `rotate(45, ${x.c1}, ${y.c3})` : null) },
				{ x: x.p2, y: y.p3, cx: x.c2, cy: y.c3, t: (lr ? `rotate(45, ${x.c2}, ${y.c3})` : null) },
				{ x: x.p3, y: y.p3, cx: x.c3, cy: y.c3, t: (lr ? `rotate(45, ${x.c3}, ${y.c3})` : null) },
			];
		},
		eeStyle() {
			if (!this.ee) {
				return null;
			}

			const randClr = ["#EC5f67", "#F99157", "#FAC863", "#99C794", "#5FB3B3", "#C594C5", "#6699CC", "#AB7967"];
			const clr     = randClr[_.random(0, randClr.length - 1)];

			return {
				fill: clr,
				stroke: clr
			};
		}
	},

	watch: {
		ee: function (newValue, oldValue) {
			if (newValue !== oldValue) {
				this.goodBye = false;
			}
		}
	},

	mounted() {
		// Prevent "Hydration completed but contains mismatches."
		this.t === -1 && this.updateCube();
	},


	// ### [Methods] ########################################

	methods: {
		updateCube() {
			this.t = _.random(0, this.layout.length * 3);
		},

		deleteCube() {
			this.goodBye = true;
			this.$emit("deleted");
		}
	}

};
</script>
<style lang="stylus">
@import "~style/functions"


.cube
	position relative
	display inline-block
	width 32px
	height 32px
	opacity 1

	&__transition
		cursor pointer

	&--fade-enter-active &__transition, &--fade-leave-active &__transition
		transition opacity 0.3s ease-in-out, transform 0.3s ease-in-out
		transform scale(1)
		transform-origin center center

	&--fade-enter &__transition, &--fade-leave-to > &__transition
		opacity 0
		transform scale(0)

	&__main
		fill #DDD
		stroke none
		stroke-width 1px

	&__mini
		fill #678

	&--logo &__main
		fill: $base-color["api"]
		stroke-width 2px
		stroke: $base-color["api"]

		&:nth-child(2n)
			fill: $base-color["blue"]
			stroke: $base-color["blue"]

	&--logo &__mini
		pointer-events none
		fill #111

</style>
