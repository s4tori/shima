<template>
	<p class="shima-checkbox">
		<input  :id="cid" type="checkbox" @change="updateValue" :checked="value" />
		<label :for="cid"><slot></slot></label>
	</p>
</template>
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
			padding-left 25px
			cursor pointer
			user-select none
			line-height shima-vr()
			display inline-block
			color: $base-color["api"]

		& + label:before
			content ""
			position absolute
			left 0
			top $cb-top
			width $cb-width
			height $cb-height
			border 1px solid #a4b1d8
			border 1px solid #999
			border-radius 3px
			//box-shadow inset 0 1px 3px rgba(0,0,0,.3)
			box-shadow 0 1px 0px rgba(0,0,0,.3)
			transition all .3s

		& + label:after
			content ""
			position absolute
			top: $cb-top + $cb-offset
			left $cb-offset
			width: $cb-width - (2 * $cb-offset)
			height $cb-height - (2 * $cb-offset)
			border 1px solid #999
			background-color #999
			border-radius 2px
			//box-shadow inset 3px 3px 0px rgba(0,0,0,.3)
			transition all .3s

		&:hover + label:before, &:hover + label:after
			border-color: $base-color["cb-active"]

		&:hover + label:after
			background-color: $base-color["cb-active"]

	[type="checkbox"]:not(:checked) + label:after
		opacity 0
		transform scale(1) rotate(0deg)

	[type="checkbox"]:checked + label:after
		opacity 1
		transform scale(1) rotate(225deg)
		transform scale(1) rotate(0deg)

</style>
<script>
export default {

	props: {
		value: { type: Boolean, default: false }
	},

	data() {
		return {
			cid: "checkbox-" + this._uid
		};
	},


	// ### [Methods] ########################################

	methods: {
		updateValue() {
			this.$emit("input", !this.value);
		}
	}

};
</script>
