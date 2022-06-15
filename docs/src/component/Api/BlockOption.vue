<template>
	<li class="block-option">
		<code>{{ title }}</code>
		<slot />
	</li>
</template>
<script setup>
defineProps({
	title: { type: String, default: "" }
});
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.block-option
	code
		flex 0 0 shima-span()
		margin-right shima-gutter()
		shima-font-size(-2)
		color: $base-color["blue"]
		text-align right

	p
		shima-font-size(-2)

</style>
