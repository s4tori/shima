<template>
	<header class="header">
		<wrapper :grid="false" page gutter dark>
			<a class="header__content" href="">
				<div class="header__logo">
					<h1 class="header__title">Shima</h1>
					<logo :toggle-ee="ee" class="header__svg" />
				</div>
				<div class="header__text">
					<h2 class="header__jp" @click.prevent="ee = !ee">縞</h2>
					<p>▨ Shima is a quick, painless, javascript-free grid overlay. Helpful for vertical rhythm, grid systems, icon designing and more. ▨</p>
				</div>
			</a>
		</wrapper>
	</header>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.header
	padding-top $base-spacing--xl
	padding-bottom $base-spacing--xl
	margin-bottom $base-spacing--xl
	background #111

	&__content
		display flex
		text-decoration none
		color alpha(#FFF, .9)
		outline none

	&__text
		margin-left auto
		width shima-span(2, "wide")
		max-width 100%
		text-align justify

	&__title
		display block
		shima-font-size(2)
		padding-bottom shima-vr(2)
		font-weight 200
		opacity 0

	&__logo
		flex 1 1 500px
		flex 1 1 shima-span(3, "wider") + shima-gutter()
		shima-font-size(2)
		padding-bottom shima-vr(2)
		font-weight 200

	&__svg
		width 100%
		max-height shima-vr(4)
		padding-right shima-span()

		+shima-mq("md-and-dn")
			padding-right shima-vr()

	&__jp
		// 🌸 Set a font containing standard Kanji glyphs 🌸
		// http://www.ruanyifeng.com/blog/2014/07/chinese_fonts.html
		// https://stackoverflow.com/a/14573813
		// http://www.xwbetter.com/font-family/
		font-family "Microsoft Yahei", "微软雅黑", STXihei, "华文细黑", "MS PGothic", sans-serif, serif;
		//font-family: Tahoma, Helvetica, Arial, "Microsoft Yahei","微软雅黑", STXihei, "华文细黑", sans-serif;
		font-weight normal
		shima-font-size(1)
		margin-bottom $base-spacing--lg
		cursor pointer


	+shima-mq("sm")
		padding-top $base-spacing
		padding-bottom $base-spacing
		margin-bottom $base-spacing

		&__content
			display block

		&__logo
			padding-bottom shima-vr()
			margin 0 auto

		&__svg
			padding 0

		&__jp
			margin-bottom shima-vr()
			text-align center

		&__title
			display none

		&__text
			width 100%
			margin-left 0
			text-align center

</style>
<script>
import Wrapper from "src/component/Wrapper.vue";
import Logo    from "./Logo.vue";

export default {

	components: {
		Wrapper,
		Logo
	},

	data() {
		return {
			ee: false
		};
	}

};
</script>
