<template>
	<div class="gex-blueprint">
		<shima-grid-manager
			:x-grid="128" :x-div="8" :x-gutter="0"
			bg="#0F57A3"
			x-grid-color="255-255-255-0.3"
			x-div-color="255-255-255-0.1"
		/>
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
