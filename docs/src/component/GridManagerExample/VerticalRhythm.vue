<template>
	<div class="grid-example">
		<shima-grid-manager
			:x-grid="0"  :x-div="0" :x-gutter="0"
			:y-grid="32" :y-div="0" :y-gutter="0"
		>
			<div class="grid-example__content">
				<h1>Vertical Rhythm</h1>
				<p>
					<a href="http://www.cupcakeipsum.com/" target="_blank">Cupcake ipsum dolor sit amet</a>. Fruitcake powder ice cream.
					Tart macaroon marzipan chupa chups macaroon sweet roll cookie.
					Apple pie gummi bears chupa chups halvah pie cookie jelly.
					Muffin cake danish caramels jujubes. Candy I love biscuit.
				</p>
				<p>
					Croissant gummi bears bonbon I love cookie muffin. Jujubes chocolate jelly fruitcake cheesecake jelly-o I love.
					Gingerbread fruitcake marshmallow biscuit I love. Tiramisu candy liquorice.
					Brownie carrot cake bear claw chupa chups powder.
				</p>
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


.grid-example

	&__content

		h1
			overflow hidden
			text-overflow ellipsis
			white-space nowrap

		a
			position relative
			z-index 0
			display inline-block
			color inherit
			text-decoration none
			transition all 0.3s ease-in-out
			$color = #FFF87E

			&::after
				$aColor = alpha(#FFF87E, 0.6)
				shima-absolute top left -7px right -7px bottom 0
				z-index -1
				height shima-vr()
				content ""
				background-image linear-gradient(to top, $aColor 0%, $aColor 100%)
				background-repeat no-repeat
				background-position 0 100%
				background-size 100% 5px
				border-radius 10px

			&:hover
				color darken($color, 50)

		p
			margin-bottom shima-vr()
			color: $base-color["api"]

</style>
