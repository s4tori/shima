<template>
	<div>
		<div v-for="(v, k) in examples" :key="k" :id="'example-' + v" />

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
				<link-example :cube="1" v-model="example" type="vr">Vertical Rhythm</link-example>
				<link-example :cube="2" v-model="example" type="layout">Grid System</link-example>
				<link-example :cube="7" v-model="example" type="vrgrid">VR + Grid</link-example>
			</ul>

			<ul class="legend__list">
				<link-example :cube="23" v-model="example" type="bp">Blue print</link-example>
				<link-example :cube="0"  v-model="example" type="icon">Icon</link-example>
				<link-example :cube="10" v-model="example" type="iso">Isometric grid</link-example>
			</ul>

			<ul class="legend__list">
				<link-example :cube="20" v-model="example" type="oblique">Oblique grid</link-example>
				<link-example :cube="18" v-model="example" type="stripe-r">Stripe (right)</link-example>
				<link-example :cube="19" v-model="example" type="stripe-l">Stripe (left)</link-example>
			</ul>
		</div>

		<grid-example :type="example" />
	</div>
</template>
<style lang="stylus">
@import "~style/main"


.legend
	display flex
	align-items center
	flex-direction row
	flex-wrap wrap
	margin-bottom shima-vr()

	&__list
		list-style-type none
		margin-right shima-gutter(2)

</style>
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
		const examples = ["vr", "layout", "vrgrid", "bp", "icon", "iso", "oblique", "stripe-r", "stripe-l"];
		const hash     = _.trimStart(window.location.hash, "#example-");
		const example  = _.includes(examples, hash) ? hash : examples[0];

		return {
			examples,
			example
		};
	}

};
</script>
