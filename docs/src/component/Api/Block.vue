<template>
	<div :id="id" class="block">
		<h3 class="block__title"><a :href="'#' + id">{{ title }} <icon i="anchor" /></a></h3>
		<div class="block__content">
			<ul class="block__list">
				<slot />
			</ul>
		</div>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"

.block
	&__title
		padding-top $base-spacing
		margin-left $base-spacing

		+shima-mq("sm")
			margin-left $base-spacing--sm

	&__content
		shima-font-size(-2)

	&__list
		reset-list()

		li
			display flex
			flex-direction row
			margin-bottom shima-vr()

			+shima-mq("sm")
				padding 0 $base-spacing--sm
				display block

	em
		text-decoration none
		color: $base-color["blue"]
		font-family $base-font-code
		font-style normal
		line-height $base-line-height - 2px // font-family fix

</style>
<script>
import Icon from "src/component/Icon.vue";


export default {

	props: {
		id   : { type: String, default: "" },
		title: { type: String, default: "" }
	},

	components: {
		Icon
	}

};
</script>
