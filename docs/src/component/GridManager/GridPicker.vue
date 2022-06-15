<template>
	<div class="shima-gp">
		<div class="shima-gp__input" @click="updateValue">
			<svg class="shima-gp__svg" xmlns="http://www.w3.org/2000/svg">
				<defs>
					<pattern :id="pid" :patternTransform="pt" :width="w" :height="h" x="0" y="0" patternUnits="userSpaceOnUse">
						<g transform="translate(0.5, 0.5)">
							<path :d="`M 0 0 L ${w} 0 M 0 0 L 0 ${h}`" />
						</g>
					</pattern>
				</defs>
				<rect :fill="`url(#${pid})`" x="0" y="0" width="100%" height="100%" />
			</svg>
		</div>
	</div>
</template>
<script>
export default {

	props: {
		modelValue: { type: String,  default: "" }
	},

	emits: ["update:modelValue"],

	data: function () {
		return {
			pid: "pattern-grid-" + this.$.uid,
			types: ["standard", "oblique", "isometric", "2-1-isometric", "trimetric-left", "trimetric-right"],
			current: 0,
			w: 20,
			h: 20
		};
	},

	computed: {
		pt() {
			switch (this.modelValue) {
				case "standard"       : return null;
				case "oblique"        : return "matrix(1,1.000,-1.000,1,0,0) scale(0.707,0.707)";
				case "isometric"      : return "matrix(1,0.577,-1.732,1,0,0) scale(0.866,0.500)";
				case "2-1-isometric"  : return "matrix(1,0.510,-1.963,1,0,0) scale(0.891,0.454)";
				case "trimetric-left" : return "matrix(1,-0.158,0.700,1,0,0) scale(0.988,0.819)";
				case "trimetric-right": return "matrix(1,0.158,-0.466,1,0,0) scale(0.988,0.906)";
				default               : return null;
			}
		}
	},


	// ### [Methods] ########################################

	methods: {
		updateValue() {
			this.$emit("update:modelValue", this.types[++this.current % this.types.length]);
		}
	}

};
</script>
<style lang="stylus">
@import "~style/functions"


.shima-gp
	position relative

	&__input
		display block
		height shima-vr(2) + $base-spacing--sm
		overflow hidden
		cursor pointer
		background white
		border 1px solid #EEEEEE
		border-radius 3px

	&__svg
		width 100%
		height 100%

		path
			fill none
			stroke #DDD
			stroke-width 1px

</style>
