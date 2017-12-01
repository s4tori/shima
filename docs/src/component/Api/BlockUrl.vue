<template>
	<div class="block-url">
		<template v-if="url">{{base}}</template>
		<slot />
	</div>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.block-url
	position relative
	padding $base-spacing--sm $base-spacing
	background: $base-color["bg-url"]
	shima-font-size(-2)
	margin-bottom $base-spacing

	&::after
		content ""
		absolute: top right bottom left
		border-top 1px solid #CCCCCC
		pointer-events none

	em
		position relative
		display inline-block
		margin 0 $base-spacing--xs
		padding 0 $base-spacing--xs
		font-style normal
		background-color alpha($base-color["api"], .1)
		border-radius 5px
		transition all 300ms ease-in-out

		&::after
			content ""
			absolute: top right bottom left
			border 3px solid alpha($base-color["api"], .1)
			border-color #DDD
			border-radius 5px
			transition all 300ms ease-in-out
			opacity 0

		&:hover
			background transparent

			&:after
				opacity 1

</style>
<script>
import { getEnvConfig } from "src/util";


export default {

	props: {
		url: { type: Boolean, default: false },
	},

	data() {
		return {
			base: getEnvConfig("urlGrid")
		};
	}

};
</script>
