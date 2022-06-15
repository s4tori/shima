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
<script setup>
import { ref }       from "vue";
import ApiGeneral    from "./ApiGeneral.vue";
import ApiAxis       from "./ApiAxis.vue";
import ApiColor      from "./ApiColor.vue";
import Wrapper       from "src/component/Wrapper.vue";
import Icon          from "src/component/Icon.vue";
import ShimaCheckbox from "src/component/Checkbox.vue";


const grid = ref(false);
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.api
	margin-bottom $base-spacing--xl

	&__content
		position relative
		overflow hidden
		color: $base-color["api"]
		background white
		border-bottom 5px solid #EAEAEA

	&__cb
		position absolute
		top shima-gutter(1)
		right $base-spacing
		z-index 1

		+shima-mq("sm")
			display none

	&__def
		position relative
		color: $base-color["api"]
		text-decoration underline
		outline none

		svg
			margin-left $base-spacing--xs

		&:after
			shima-absolute top -5px right -12px bottom -5px left -12px
			content ""
			border-radius 5px
			opacity 0
			transition all 0.3s ease-in-out

		&:hover:after
			background alpha($base-color["api"], 0.1)
			opacity 1

</style>
