<template>
	<block id="api-general" title="General Configuration" class="api-general">
		<block-url url>
			<em>x axis config</em> / <em>&lt;y axis config&gt;</em> / <em>&lt;type&gt;</em>
		</block-url>

		<block-option title="x axis config">
			<p>
				Configuration for the X axis.<br />
				Mandatory.<br />
				<a href="#api-axis" class="api__def">See definition<icon i="arrow" /></a>
			</p>
		</block-option>

		<block-option title="y axis config">
			<p>
				Configuration for the Y axis.<br />
				Optional : If not provided, the <em>x-axis</em> and <em>y-axis</em> will have the same configuration.<br />
				<a href="#api-axis" class="api__def">See definition<icon i="arrow" /></a>
			</p>
		</block-option>

		<block-option title="type">
			<p>
				Grid type.<br />
				Optional : If not provided, the grid type will be equal to <em>standard</em>.<br />
				Values : <em>standard</em> / <em>oblique</em> / <em>isometric</em> / <em>isometric-2-1</em> / <em>trimetric-left</em> / <em>trimetric-right</em>.
			</p>
		</block-option>
	</block>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.api-general

	&__type
		position relative
		text-decoration underline
		color: $base-color["api"]

		svg
			margin-left $base-spacing--xs

		&:after
			content ""
			absolute: top -5px right -12px bottom -5px left -12px
			border-radius 5px
			opacity 0
			transition all .3s ease-in-out

		&:hover:after
			opacity 1
			background alpha($base-color["api"], .1)

</style>
<script>
import Block       from "./Block.vue";
import BlockUrl    from "./BlockUrl.vue";
import BlockOption from "./BlockOption.vue";
import Icon        from "src/component/Icon.vue";


export default {

	components: {
		Block,
		BlockUrl,
		BlockOption,
		Icon
	}

};
</script>
