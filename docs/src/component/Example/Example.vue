<template>
	<div>
		<div :key="k" :id="'example-' + v" v-for="(v, k) in examples"></div>

		<h2 id="examples">
			<a href="#examples">▨ Examples <icon i="anchor" /></a>
		</h2>

		<div v-if="false">
			<cube :key="i" :type="i-1"    v-for="i in 10" /><br />
			<cube :key="i" :type="i-1+10" v-for="i in 10" /><br />
			<cube :key="i" :type="i-1+20" v-for="i in 10" /><br />
		</div>

		<div class="legend">
			<ul class="legend__list">
				<link-example :cube="1" type="vr"        v-model="example">Vertical Rhythm</link-example>
				<link-example :cube="2" type="layout"    v-model="example">Grid System</link-example>
				<link-example :cube="7" type="vrgrid"    v-model="example">VR + Grid</link-example>
			</ul>

			<ul class="legend__list">
				<link-example :cube="23" type="bp"       v-model="example">Blue print</link-example>
				<link-example :cube="0"  type="icon"     v-model="example">Icon</link-example>
				<link-example :cube="10" type="iso"      v-model="example">Isometric grid</link-example>
			</ul>

			<ul class="legend__list">
				<link-example :cube="20" type="oblique"  v-model="example">Oblique grid</link-example>
				<link-example :cube="18" type="stripe-r" v-model="example">Stripe (right)</link-example>
				<link-example :cube="19" type="stripe-l" v-model="example">Stripe (left)</link-example>
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

	data() {
		const examples = ["vr", "layout", "vrgrid", "bp", "icon", "iso", "oblique", "stripe-r", "stripe-l"];
		const hash     = _.trimStart(window.location.hash, "#example-");
		const example  = _.includes(examples, hash) ? hash : examples[0];

		return {
			examples,
			example
		};
	},

	components: {
		LinkExample,
		GridExample,
		Cube,
		Icon
	}

};
</script>
