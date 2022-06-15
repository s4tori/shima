<template>
	<div class="gex-vrgrid">
		<shima-grid-manager
			:x-grid="128" :x-div="4" :x-gutter="-32"
			:y-grid="128" :y-div="4" :y-gutter="0"
		>
			<div class="gex-vrgrid__content">
				<h1>VR + Grid</h1>
				<div class="gex-vrgrid__row">
					<div class="gex-vrgrid__img">image</div>
					<div class="gex-vrgrid__text">
						<h2 class="gex-vrgrid__title">🍰 Dessert title</h2>
						<p>
							<a href="http://www.cupcakeipsum.com/" target="_blank">Cupcake ipsum dolor sit amet</a>. Fruitcake powder ice cream.
							Tart macaroon marzipan chupa chups macaroon sweet roll cookie.
							Apple pie gummi bears chupa chups halvah pie cookie jelly.
							Muffin cake danish caramels jujubes. Candy I love biscuit.
						</p>
					</div>
				</div>
			</div>
		</shima-grid-manager>
	</div>
</template>
<script setup>
import ShimaGridManager from "src/component/GridManager/GridManager.vue";
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.gex-vrgrid

	&__content
		margin 0 shima-vr()

		h1
			overflow hidden
			text-overflow ellipsis
			white-space nowrap

		a
			$redColor = #EC173A
			position relative
			z-index 0
			display inline-block
			color inherit
			text-decoration none
			box-shadow 0 5px 0 alpha($redColor, 0.2)
			transition all 0.3s ease-in-out

			&:hover
				color darken($redColor, 20)

	&__img
		display flex
		align-items center
		justify-content center
		height shima-vr(8)
		margin-bottom shima-vr()
		color black
		background alpha(#EAEAEA, 0.8)

	&__text
		flex 1 0 shima-span(2)
		margin-right $base-spacing
		text-align justify

	// To override ".main h2"
	& h2&__title
		color inherit
		background none

		&::after
			display none

	&__img
		flex 0 0 shima-span(2)
		width shima-span(2)
		margin-right shima-gutter()
		color white
		background alpha(#0074D9, 0.9)

		$c1  = alpha(#ECEDDC, 0.2)
		$c2  = alpha(#EC173A, 1)
		$l1 = linear-gradient(135deg, $c1 25%, transparent 25%) -50px 0
		$l2 = linear-gradient(225deg, $c1 25%, transparent 25%) -50px 0
		$l3 = linear-gradient(315deg, $c1 25%, transparent 25%)
		$l4 = linear-gradient( 45deg, $c1 25%, transparent 25%)
		background $l1, $l2, $l3, $l4
		background-color $c2
		background-size 100px 100px
		border-radius 10px
		box-shadow 0 5px 0 alpha(darken($c2, 20), 0.8)

	&__row
		display flex
		flex-wrap wrap

	+shima-mq("md-and-dn")

		&__row
			display block

		&__img
			width 100%
			margin-right 0

		&__title
			text-align center

		&__text
			margin-right 0

</style>
