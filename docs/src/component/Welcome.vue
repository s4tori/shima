<template>
	<div class="welcome">
		<h2 id="welcome">
			<a href="#welcome">▨ Getting Started <icon i="anchor" /></a>
		</h2>

		<wrapper class="welcome__content" :grid="grid" top>
			<h3 class="welcome__title">Hi!</h3>
			<p>
				Shima helps you to render custom lovely grids 💖.
				You can use Shima as is without installing anything!
				Read further to:
			</p>
			<ul class="welcome__list">
				<li><a href="#grid"> ▨ <em>create</em></a> your own grid</li>
				<li><a href="#examples">▨ <em>see</em></a> few examples</li>
				<li><a href="#api"> ▨ <em>read</em></a> the API </li>
			</ul>

			<p>
				Do you want to display a grid with gutters? Simply add a line of CSS and you're done:
			</p>

			<shima-code class="welcome__code" :url="`${base}128x4x32/128x4x0`">
				<shima-checkbox v-model="grid" class="welcome__cb">Show grid</shima-checkbox>
			</shima-code>
		</wrapper>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"

.welcome
	margin-bottom $base-spacing--lg

	&__content
		position relative
		background white
		border-bottom 5px solid #EAEAEA
		color: $base-color["api"]
		overflow hidden

		p
			padding 0 $base-spacing
			margin-bottom $base-spacing

	&__title
		padding-top $base-spacing
		margin-left $base-spacing

	&__cb.shima-checkbox
		margin 0
		padding 0
		padding-top $base-spacing
		shima-font-size(-2, shima-vr())

	&__cb2
		position absolute
		right 0
		bottom 0
		z-index 1

		+shima-mq("sm")
			display none

	&__list
		reset-list()
		padding 0 $base-spacing $base-spacing $base-spacing

		li
			margin-left $base-spacing

			+shima-mq("sm")
				margin-left 0

		a
			reset-link()

			em
				text-decoration underline
				font-style normal

</style>
<script>
import { getEnvConfig } from "src/util";
import Wrapper          from "src/component/Wrapper.vue";
import ShimaCheckbox    from "src/component/Checkbox.vue";
import ShimaCode        from "src/component/GridManager/Code.vue";
import Icon             from "src/component/Icon.vue";


export default {

	data() {
		return {
			base: getEnvConfig("urlGrid"),
			grid: false
		};
	},

	components: {
		Wrapper,
		ShimaCheckbox,
		ShimaCode,
		Icon
	}

};
</script>
