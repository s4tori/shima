<template>
	<div>
		<div v-for="(v, k) in examples" :id="'example-' + v" :key="k" />

		<h2 id="examples">
			<a href="#examples">▨ Examples <icon i="anchor" /></a>
		</h2>

		<div v-if="false">
			<cube v-for="i in 10" :key="i" :type="i-1"    /><br />
			<cube v-for="i in 10" :key="i" :type="i-1+10" /><br />
			<cube v-for="i in 10" :key="i" :type="i-1+20" /><br />
		</div>

		<div class="legend">
			<ul class="legend__list">
				<link-example v-model="example" :cube="1" type="vr">Vertical Rhythm</link-example>
				<link-example v-model="example" :cube="2" type="layout">Grid System</link-example>
				<link-example v-model="example" :cube="7" type="vrgrid">VR + Grid</link-example>
			</ul>

			<ul class="legend__list">
				<link-example v-model="example" :cube="23" type="bp">Blue print</link-example>
				<link-example v-model="example" :cube="0"  type="icon">Icon</link-example>
				<link-example v-model="example" :cube="10" type="iso">Isometric grid</link-example>
			</ul>

			<ul class="legend__list">
				<link-example v-model="example" :cube="20" type="oblique">Oblique grid</link-example>
				<link-example v-model="example" :cube="18" type="stripe-r">Stripe (right)</link-example>
				<link-example v-model="example" :cube="19" type="stripe-l">Stripe (left)</link-example>
			</ul>
		</div>

		<grid-example :type="example" />
	</div>
</template>
<script>
import { _ }       from "src/util";
import LinkExample from "src/component/Example/LinkExample.vue";
import GridExample from "src/component/Example/GridExample.vue";
import Cube        from "src/component/Cube.vue";
import Icon        from "src/component/Icon.vue";


export default {

	components: {
		LinkExample,
		GridExample,
		Cube,
		Icon
	},

	data() {
		return {
			examples: ["vr", "layout", "vrgrid", "bp", "icon", "iso", "oblique", "stripe-r", "stripe-l"],
			example: "vr"
		};
	},

	// location.hash isn't available on server side
	// We check the hash with mounted event to let Vue hydrates existing markup correctly first
	mounted() {
		const hash   = _.replace(global.location.hash, "#example-", "");
		this.example = _.includes(this.examples, hash) ? hash : this.examples[0];
	}

};
</script>
<style lang="stylus">
@import "~style/main"


.legend
	display flex
	flex-direction row
	flex-wrap wrap
	align-items center
	margin-bottom shima-vr()

	&__list
		margin-right shima-gutter(2)
		list-style-type none

</style>
