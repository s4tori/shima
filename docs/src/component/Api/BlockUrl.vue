<template>
	<div class="block-url">
		<template v-if="url">{{ base }}</template>
		<slot />
	</div>
</template>
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
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.block-url
	position relative
	padding $base-spacing--sm $base-spacing
	margin-bottom $base-spacing
	shima-font-size(-2)
	background: $base-color["bg-url"]


	&::after
		shima-absolute top right bottom left
		pointer-events none
		content ""
		border-top 1px solid #CCCCCC

	em
		position relative
		display inline-block
		padding 0 $base-spacing--xs
		margin 0 $base-spacing--xs
		font-style normal
		background-color alpha($base-color["api"], 0.1)
		border-radius 5px
		transition all 300ms ease-in-out

		&::after
			shima-absolute top right bottom left
			content ""
			border 3px solid alpha($base-color["api"], 0.1)
			border-color #DDD
			border-radius 5px
			opacity 0
			transition all 300ms ease-in-out

		&:hover
			background transparent

			&:after
				opacity 1

</style>
