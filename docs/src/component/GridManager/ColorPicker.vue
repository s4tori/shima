<template>
	<div class="shima-cp">
		<input
			class="shima-cp__input"
			readonly type="text"
			@focus="showPicker"
			:style="'background-color: ' + this.rtColor"
		/>
		<sketch-picker
			class="shima-cp__picker"
			v-if="displayPicker"
			:value="colors"
			@input="updateValueFromPicker"
		/>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"


.shima-cp
	position relative
	height shima-vr()
	margin-bottom $base-spacing--sm
	background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAMElEQVQ4T2N89uzZfwY8QFJSEp80A+OoAcMiDP7//483HTx//hx/Ohg1gIFx6IcBALl+VXknOCvFAAAAAElFTkSuQmCC");
	border 1px solid #EEEEEE
	border-radius 3px

	&__input
		display block
		border none
		background #00D99E
		outline none
		cursor pointer
		width 100%
		border-radius 3px
		overflow: hidden
		height: shima-vr()
		line-height: shima-vr()

	&__picker
		position absolute !important
		top shima-vr() + $base-spacing--xs
		right 0
		z-index 1

		input
			line-height 12px

</style>
<script>
import { Sketch } from "vue-color";
import { _ }      from "src/util";


export default {

	// https://vuejs.org/v2/guide/components.html#Form-Input-Components-using-Custom-Events
	// We use "v-model=" instead of ":color.sync="
	model: {
		prop: "color",
		event: "input"
	},

	props: {
		color: { type: String, default: "#194d33" }
	},

	data: function () {
		const cols = this.color ? _.split(this.color, "-", 4) : [""];
		const c    = cols.length > 1
			? `rgba(${cols[0]},${cols[1]},${cols[2]},${cols[3]})`
			: "#" + this.color
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

	components: {
		"sketch-picker": Sketch
	},


	// ### [Methods] ########################################

	methods: {
		showPicker() {
			document.addEventListener("click", this.documentClick);
			this.displayPicker = true;
		},

		hidePicker() {
			// https://vuejs.org/v2/guide/components.html#sync-Modifier
			// this.$emit("update:color", this.getGridColor(this.colors));
			this.$emit("input", this.getGridColor(this.colors));
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
			if (_.isString(c)) { return this.color; }
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
