<template>
	<div>
		<shima-grid-manager
			:xGrid="30" :xDiv="0" :xGutter="0"
			:yGrid="0"  :yDiv="0" :yGutter="0"
			t="oblique"
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
