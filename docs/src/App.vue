<template>
	<!-- The root element needs to have the same ID than the mounted DOM element -->
	<!-- See: https://github.com/chrisvfritz/prerender-spa-plugin#caveats -->
	<div id="app">
		<main-view />
	</div>
</template>
<script>
import MainView from "src/view/Main.vue";


export default {

	components: {
		MainView
	}

};
</script>
