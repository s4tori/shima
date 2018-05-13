<template>
	<div class="gex-bp">
		<shima-grid-manager
			:x-grid="128" :x-div="4" :x-gutter="-32"
		>
			<div class="gex-bp__icons">
				<icon i="beaker" class="gex-bp__icon" />
				<icon i="rocket" class="gex-bp__icon" />
				<icon i="arrow"  class="gex-bp__icon" />
				<icon i="kebab"  class="gex-bp__icon" />
			</div>
		</shima-grid-manager>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"

.gex-bp
	&__icons
		font-size 0

	&__icon
		display inline-block
		margin-left shima-vr()
		margin-top shima-vr()
		width shima-span()
		height shima-span()

</style>
<script>
import ShimaGridManager from "src/component/GridManager/GridManager.vue";
import Icon             from "src/component/Icon.vue";


export default {

	components: {
		ShimaGridManager,
		Icon
	}

};
</script>
