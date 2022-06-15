<template>
	<div>
		<shima-grid-manager
			:x-grid="0"  :x-div="0" :x-gutter="0"
			:y-grid="30" :y-div="0" :y-gutter="0"
			t="oblique"
		/>
	</div>
</template>
<script setup>
import ShimaGridManager from "src/component/GridManager/GridManager.vue";
</script>
