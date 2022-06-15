<template>
	<div class="gex-bp">
		<shima-grid-manager
			:x-grid="128" :x-div="4" :x-gutter="-32"
		>
			<div class="gex-bp__icons">
				<icon i="terminal"    class="gex-bp__icon" />
				<icon i="arrow-heart" class="gex-bp__icon" />
				<icon i="arrow"       class="gex-bp__icon" />
				<icon i="layout"      class="gex-bp__icon" />
				<icon i="translate"   class="gex-bp__icon" />
				<icon i="kiss"        class="gex-bp__icon" />
			</div>
		</shima-grid-manager>
	</div>
</template>
<script setup>
import ShimaGridManager from "src/component/GridManager/GridManager.vue";
import Icon             from "src/component/Icon.vue";
</script>
<style lang="stylus">
@import "~style/functions"

.gex-bp
	&__icons
		font-size 0

	&__icon
		display inline-block
		width shima-span()
		height shima-span()
		margin-top shima-vr()
		margin-left shima-vr()

</style>
