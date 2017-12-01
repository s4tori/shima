<template>
	<div class="gex-blueprint">
		<shima-grid-manager
			:xGrid="128" :xDiv="8" :xGutter="0"
			bg="#0F57A3"
			xGridColor="255-255-255-0.3"
			xDivColor="255-255-255-0.1"
		>
		</shima-grid-manager>
	</div>
</template>
<script>
import ShimaGridManager from "src/component/GridManager/GridManager.vue";


export default {

	components: {
		ShimaGridManager
	}

};
</script>
