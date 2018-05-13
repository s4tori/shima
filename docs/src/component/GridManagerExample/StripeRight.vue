<template>
	<div>
		<shima-grid-manager
			:x-grid="30" :x-div="0" :x-gutter="0"
			:y-grid="0"  :y-div="0" :y-gutter="0"
			t="oblique"
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
