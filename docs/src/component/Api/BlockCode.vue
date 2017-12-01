<template>
	<li class="block-code">
		<h3>Examples</h3>
		<slot />
	</li>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


li.block-code
	background: $base-color["bg-code"]
	display flex
	flex-direction column
	padding shima-vr() 0
	margin-bottom 0

	h3
		margin-left shima-span(1, "wide")
		color: $base-color["blue"]

		+shima-mq("sm")
			margin-left shima-gutter()

	.block-option
		margin-bottom 0

	code
		font-family $base-font-code
		color #A4B1D8

	p
		font-family $base-font-code
		color: $base-color["blue"]

</style>
