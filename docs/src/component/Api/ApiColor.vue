<template>
	<block id="api-color" title="Color configuration">
		<block-url>
			<em>hexa</em> or <em>r-g-b</em> - <em>&lt;a&gt;</em>
		</block-url>

		<block-option title="hexa">
			<p>
				Hexadecimal color (without <em>#</em>).<br />
				Mandatory.
			</p>
		</block-option>

		<block-option title="r-g-b">
			<p>
				RGB color<br />
				Mandatory.
			</p>
		</block-option>

		<block-option title="a">
			<p>
				Alpha channel.<br />
				Optional (if not provided, <em>a</em> will be equal to <em>1</em>).
			</p>
		</block-option>

		<block-code>
			<ul>
				<block-option title="hex"><p><span class="api-color__ex">FF00FF</span>090210</p></block-option>
				<block-option title="rgb"><p><span class="api-color__ex">255-0-255</span>9-2-16</p></block-option>
				<block-option title="rgba"><p><span class="api-color__ex">255-0-255-1</span>9-2-16-.5</p></block-option>
			</ul>
		</block-code>
	</block>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.api-color

	&__ex
		display inline-block
		width shima-span(1, "wide")

</style>
<script>
import Block       from "./Block.vue";
import BlockUrl    from "./BlockUrl.vue";
import BlockOption from "./BlockOption.vue";
import BlockCode   from "./BlockCode.vue";


export default {

	components: {
		Block,
		BlockUrl,
		BlockOption,
		BlockCode
	}

};
</script>
