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
<style lang="stylus">
@import "~style/functions"

.gex-layout

	&__content p
		margin-bottom shima-vr()

	&__row
		display flex

	&__sidebar, &__main
		display flex
		justify-content center
		align-items center
		height shima-vr(8)
		margin-bottom: shima-vr()
		color black
		background alpha(#EAEAEA, .8)

	&__sidebar
		width shima-span()
		flex 0 0 shima-span()
		margin-right shima-gutter()
		background alpha(#0074D9, .9)
		color white

	&__main
		flex 1
		background alpha(#3D9970, .9)
		color white

	&__col
		display flex
		justify-content center
		margin-bottom: shima-vr()
		align-items center
		background alpha(#000, .1)
		background alpha(#EAEAEA, .5)
		box-shadow 0 3px 0px rgba(0,0,0,.1)

		&--narrow
			width shima-span()
			margin-right shima-gutter()

		&--wide
			width shima-span(2, "wide")
			margin-left shima-gutter()

		&--flex
			flex 1

</style>
<script>
import ShimaGridManager from "src/component/GridManager/GridManager.vue";


export default {

	components: {
		ShimaGridManager
	}

};
</script>
