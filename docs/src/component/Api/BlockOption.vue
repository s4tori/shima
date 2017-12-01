<template>
	<li class="block-option">
		<code>{{title}}</code>
		<slot />
	</li>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.block-option
	code
		flex 0 0 shima-span()
		margin-right shima-gutter()
		shima-font-size(-2)
		text-align right
		color: $base-color["blue"]

	p
		shima-font-size(-2)

</style>
<script>
export default {

	props: {
		title: { type: String, default: "" }
	}

};
</script>
