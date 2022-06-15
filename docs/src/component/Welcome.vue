<template>
	<div class="welcome">
		<h2 id="welcome">
			<a href="#welcome">▨ Getting Started <icon i="anchor" /></a>
		</h2>

		<wrapper :grid="grid" class="welcome__content" top>
			<h3 class="welcome__title">Hi!</h3>
			<p>
				Shima helps you to render custom lovely grids 💖.
				You can use Shima as is without installing anything!
				Refer to the following for more details:
			</p>
			<ul class="welcome__list">
				<li><a href="#grid"> ▨ <em>create</em></a> your own grid</li>
				<li><a href="#examples">▨ <em>Look at</em></a> some examples</li>
				<li><a href="#api"> ▨ <em>read</em></a> the API</li>
			</ul>

			<p>
				Do you want to display a grid with gutters? Simply add a line of CSS and you're done:
			</p>

			<shima-code :url="`${base}128x4x32/128x4x0`" class="welcome__code">
				<shima-checkbox v-model="grid" class="welcome__cb">Show grid</shima-checkbox>
			</shima-code>
		</wrapper>
	</div>
</template>
<script>
import { getEnvConfig } from "src/util";
import Wrapper          from "src/component/Wrapper.vue";
import ShimaCheckbox    from "src/component/Checkbox.vue";
import ShimaCode        from "src/component/GridManager/Code.vue";
import Icon             from "src/component/Icon.vue";


export default {

	components: {
		Wrapper,
		ShimaCheckbox,
		ShimaCode,
		Icon
	},

	data() {
		return {
			base: getEnvConfig("urlGrid"),
			grid: false
		};
	}

};
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.welcome
	margin-bottom $base-spacing--lg

	&__content
		position relative
		overflow hidden
		color: $base-color["api"]
		background white
		border-bottom 5px solid #EAEAEA

		p
			padding 0 $base-spacing
			margin-bottom $base-spacing

	&__title
		padding-top $base-spacing
		margin-left $base-spacing

	&__cb.shima-checkbox
		padding 0
		padding-top $base-spacing
		margin 0
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
				font-style normal
				text-decoration underline

</style>
