<template>
	<div class="shima-cp">
		<input
			:style="'background-color: ' + rtColor"
			class="shima-cp__input"
			readonly type="text"
			@focus="showPicker"
		/>
		<sketch-picker
			v-if="displayPicker"
			:model-value="colors"
			class="shima-cp__picker"
			@update:model-value="updateValueFromPicker"
		/>
	</div>
</template>
<script>
import { Sketch } from "@ckpack/vue-color";
import { _ }      from "src/util";


export default {

	components: {
		"sketch-picker": Sketch
	},

	props: {
		modelValue: { type: String, default: "#00D99E" }
	},

	emits: ["update:modelValue"],

	data: function () {
		const cols = this.modelValue ? _.split(this.modelValue, "-", 4) : [""];
		const c    = cols.length > 1
			? `rgba(${cols[0]},${cols[1]},${cols[2]},${cols[3]})`
			 : "#" + this.modelValue //: "#00D99E"
		;

		return {
			displayPicker: false,
			colors: c // string color by default (vue-color will convert the value to plain object)
		};
	},

	computed: {
		rtColor() {
			return this.getValidColor(this.colors);
		}
	},


	// ### [Methods] ########################################

	methods: {
		showPicker() {
			document.addEventListener("click", this.documentClick);
			this.displayPicker = true;
		},

		hidePicker() {
			this.$emit("update:modelValue", this.getGridColor(this.colors));
			document.removeEventListener("click", this.documentClick);
			this.displayPicker = false;
		},

		documentClick(e) {
			const el = this.$el, target = e.target;
			if (el !== target && !el.contains(target)) {
				this.hidePicker();
			}
		},

		getValidColor(colors) {
			const c = colors;
			// If vue-color hasn't called "updateValueFromPicker" fonction yet
			if (_.isString(c)) { return c; }
			return (c.a < 1 && c.rgba) ? `rgba(${c.rgba.r},${c.rgba.g},${c.rgba.b},${c.a})` : c.hex;
		},

		getGridColor(colors) {
			const c = colors;
			// If vue-color hasn't called "updateValueFromPicker" fonction yet
			if (_.isString(c)) { return this.modelValue; }
			return (c.a < 1 && c.rgba) ? `${c.rgba.r}-${c.rgba.g}-${c.rgba.b}-${_.round(c.a, 1)}` : c.hex.substring(1);
		},

		updateValueFromPicker(value) {
			this.colors = value;
		},

		beforeDestroy() {
			this.hidePicker();
		}
	}

};
</script>
<style lang="stylus">
@import "~style/functions"


.shima-cp
	position relative
	height shima-vr()
	margin-bottom $base-spacing--sm
	background-image url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAMElEQVQ4T2N89uzZfwY8QFJSEp80A+OoAcMiDP7//483HTx//hx/Ohg1gIFx6IcBALl+VXknOCvFAAAAAElFTkSuQmCC")
	border 1px solid #EEEEEE
	border-radius 3px

	&__input
		display block
		width 100%
		height shima-vr()
		overflow hidden
		line-height shima-vr()
		cursor pointer
		background #00D99E
		border none
		border-radius 3px
		outline none

	&__picker
		position absolute !important
		top shima-vr() + $base-spacing--xs
		right 0
		z-index 1

		input
			line-height 12px

</style>
