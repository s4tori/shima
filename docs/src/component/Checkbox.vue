<template>
	<p class="shima-checkbox">
		<input :id="cid" :checked="modelValue" type="checkbox" @change="updateValue" />
		<label :for="cid"><slot /></label>
	</p>
</template>
<script>
export default {

	props: {
		modelValue: { type: Boolean, default: false }
	},

	emits: ["update:modelValue"],

	data() {
		return {
			cid: "checkbox-" + this.$.uid
		};
	},

	// ### [Methods] ########################################

	methods: {
		updateValue() {
			this.$emit("update:modelValue", !this.modelValue);
		}
	}

};
</script>
<style lang="stylus">
@import "~style/functions"

.shima-checkbox

	[type="checkbox"]:not(:checked), [type="checkbox"]:checked
		position absolute
		left -9999px

		$cb-top    =  (shima-vr() / 4)
		$cb-width  = (shima-gutter() / 2)
		$cb-height = (shima-vr() / 2)
		$cb-offset = 3px

		& + label
			position relative
			display inline-block
			padding-left 25px
			line-height shima-vr()
			color: $base-color["api"]
			cursor pointer
			user-select none

		& + label:before
			position absolute
			top $cb-top
			left 0
			width $cb-width
			height $cb-height
			content ""
			border 1px solid #999
			border-radius 3px
			box-shadow 0 1px 0 rgba(0, 0, 0, 0.3)
			transition all 0.3s

		& + label:after
			position absolute
			top $cb-top + $cb-offset
			left $cb-offset
			width $cb-width - (2 * $cb-offset)
			height $cb-height - (2 * $cb-offset)
			content ""
			background-color #999
			border 1px solid #999
			border-radius 2px
			transition all 0.3s

		&:hover + label:before, &:hover + label:after
			border-color: $base-color["cb-active"]

		&:hover + label:after
			background-color: $base-color["cb-active"]

	[type="checkbox"]:not(:checked) + label:after
		opacity 0
		transform scale(1) rotate(0deg)

	[type="checkbox"]:checked + label:after
		opacity 1
		transform scale(1) rotate(0deg)

</style>
