<template>
	<div class="api">
		<h2 id="api">
			<a href="#api">▨ API <icon i="anchor" /></a>
		</h2>
		<wrapper :grid="grid" class="api__content" top>
			<shima-checkbox v-model="grid" class="api__cb">Show grid</shima-checkbox>
			<api-general />
			<api-axis />
			<api-color />
		</wrapper>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.api
	margin-bottom $base-spacing--xl

	&__content
		position relative
		background white
		border-bottom 5px solid #EAEAEA
		color: $base-color["api"]
		overflow hidden

	&__cb
		position absolute
		right $base-spacing
		top shima-gutter(1)
		z-index 1

		+shima-mq("sm")
			display none

	&__def
		position relative
		text-decoration underline
		color: $base-color["api"]
		outline none

		svg
			margin-left $base-spacing--xs

		&:after
			content ""
			absolute: top -5px right -12px bottom -5px left -12px
			border-radius 5px
			opacity 0
			transition all .3s ease-in-out

		&:hover:after
			opacity 1
			background alpha($base-color["api"], .1)

</style>
<script>
import ApiGeneral    from "./ApiGeneral.vue";
import ApiAxis       from "./ApiAxis.vue";
import ApiColor      from "./ApiColor.vue";
import Wrapper       from "src/component/Wrapper.vue";
import Icon          from "src/component/Icon.vue";
import ShimaCheckbox from "src/component/Checkbox.vue";


export default {

	components: {
		ApiGeneral,
		ApiAxis,
		ApiColor,
		Wrapper,
		Icon,
		ShimaCheckbox
	},

	data() {
		return {
			grid: false
		};
	}

};
</script>
