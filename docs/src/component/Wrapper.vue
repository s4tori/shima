<template>
	<div :class="{ 'wrapper__gutter': gutter }" class="wrapper">
		<div :class="{ 'wrapper__page': page }">
			<div v-if="grid" :class="{ 'wrapper__grid--dark': dark, 'wrapper__grid--top': top }" class="wrapper__grid">
				<slot />
			</div>
			<slot v-else />
		</div>
	</div>
</template>
<script setup>
defineProps({
	page  : { type: Boolean, default: false },
	gutter: { type: Boolean, default: false },
	grid  : { type: Boolean, default: false },
	dark  : { type: Boolean, default: false },
	top   : { type: Boolean, default: false }
});
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.wrapper
	margin 0 auto

	&__page
		max-width shima-span(6, "narrow")
		margin 0 auto

	&__gutter
		margin 0 $base-spacing--sm

	&__grid
		position relative
		background-image url($base-url + "128x4x32/128x4x0")

		&::after
			shima-absolute top right bottom left
			display none
			pointer-events none
			content ""
			background-image url($base-url + "128x4x32/128x4x0")
			// background-image url($base-url + "128x4x32x0-0-0-.2x0-0-0-.1/128x4x0x0-0-0-.2x0-0-0-.1")

		&--dark
			background-image url($base-url + "128x4x32x255-255-255-.5x255-255-255-.1/128x4x0x255-255-255-.5x255-255-255-.1")

			&::after
				background-image url($base-url + "128x4x32x255-255-255-.5x255-255-255-.1/128x4x0x255-255-255-.5x255-255-255-.1")

		&--top
			background none

			&::after
				display block

</style>
