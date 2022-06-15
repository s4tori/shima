<template>
	<div class="gex-layout">
		<shima-grid-manager
			:x-grid="128" :x-div="4" :x-gutter="32"
			:y-grid="0"   :y-div="0" :y-gutter="0"
		>
			<div class="gex-layout__content">
				<div class="gex-layout__row">
					<div class="gex-layout__sidebar">sidebar</div>
					<div class="gex-layout__main">content</div>
				</div>

				<div class="gex-layout__row">
					<div class="gex-layout__col gex-layout__col--narrow">narrow</div>
					<div class="gex-layout__col gex-layout__col--flex">flex</div>
				</div>

				<div class="gex-layout__row">
					<div class="gex-layout__col gex-layout__col--flex">flex</div>
					<div class="gex-layout__col gex-layout__col--wide">wide</div>
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

.gex-layout

	&__content p
		margin-bottom shima-vr()

	&__row
		display flex

	&__sidebar, &__main
		display flex
		align-items center
		justify-content center
		height shima-vr(8)
		margin-bottom shima-vr()
		color black
		background alpha(#EAEAEA, 0.8)

	&__sidebar
		flex 0 0 shima-span()
		width shima-span()
		margin-right shima-gutter()
		color white
		background alpha(#5C87FF, 0.9)

	&__main
		flex 1
		color white
		background alpha(#00D99E, 0.9)

	&__col
		display flex
		align-items center
		justify-content center
		margin-bottom shima-vr()
		background alpha(#000, 0.1)
		background alpha(#EAEAEA, 0.5)
		box-shadow 0 3px 0 rgba(0, 0, 0, 0.1)

		&--narrow
			width shima-span()
			margin-right shima-gutter()

		&--wide
			width shima-span(2, "wide")
			margin-left shima-gutter()

		&--flex
			flex 1

</style>
