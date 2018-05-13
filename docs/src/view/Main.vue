<template>
	<div>
		<shima-header />
		<section class="main">
			<wrapper page gutter>
				<welcome-section />

				<h2 id="grid">
					<a href="#grid">▨ Create your grid <icon i="anchor" /></a>
				</h2>
				<shima-grid-manager
					:x-grid="128" :x-div="4" :x-gutter="-32"
					:y-grid="128" :y-div="4" :y-gutter="0"
				/>
				<example-section />
				<api-section />
			</wrapper>
		</section>
		<shima-footer />
	</div>
</template>
<style lang="stylus">
@import "~style/main"
</style>
<script>
import Wrapper          from "src/component/Wrapper.vue";
import ShimaGridManager from "src/component/GridManager/GridManager.vue";
import WelcomeSection   from "src/component/Welcome.vue";
import ExampleSection   from "src/component/Example/Example.vue";
import ApiSection       from "src/component/Api/Api.vue";
import ShimaHeader      from "src/component/Layout/Header.vue";
import ShimaFooter      from "src/component/Layout/Footer.vue";
import Icon             from "src/component/Icon.vue";


export default {

	components: {
		Wrapper,
		ShimaHeader,
		ShimaGridManager,
		WelcomeSection,
		ExampleSection,
		ApiSection,
		ShimaFooter,
		Icon
	}

};
</script>
