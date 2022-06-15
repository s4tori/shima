<template>
	<div class="shima-code">
		<code class="shima-code__content">
			<div v-if="type === 'css' && !top">
				<span class="token selector">.myAwesomeDiv</span> {<br />
				<span class="token property">background-image</span>:
				<span class="token value">url(<span class="token url">"{{ url }}"</span>)</span>;<br />
				}
			</div>
			<div v-else-if="type === 'css' && top">
				<span class="token selector">.myAwesomeDiv</span> {<br />
				<span class="token property">position</span>: <span class="token value">relative</span>;<br />
				}<br />
				<span class="token selector">.myAwesomeDiv::after</span> {<br />
				<template v-for="(item, index) in cssTop" :key="index">
					<span>
						<span class="token property">{{ item.key }}</span>: <span :class="item.type" class="token">{{ item.value }}</span>;
						<br />
					</span>
				</template>
				<span class="token property">background-image</span>:
				<span class="token value">url(<span class="token url">"{{ url }}"</span>)</span>;<br />
				}
			</div>
			<div v-else-if="type === 'stylus'">
				<span class="token selector">.foo</span><br />
				<span class="token property">background-image</span>
				<span class="token value">url(<span class="token url">"{{ url }}"</span>)</span><br />
			</div>
		</code>
		<slot />
	</div>
</template>
<script>
import { _ } from "src/util";


export default {

	props: {
		url : { type: String , default: "http://" },
		top : { type: Boolean, default: false     },
		type: { type: String , default: "css", validator: (value) => _.indexOf(["css", "stylus"], value) !== -1 }
	},

	data: function () {
		return {
			cssTop: [
				{ key: "content"       , value: '""'      , type: "value"  },
				{ key: "position"      , value: "absolute", type: "value"  },
				{ key: "top"           , value: "0"       , type: "number" },
				{ key: "right"         , value: "0"       , type: "number" },
				{ key: "bottom"        , value: "0"       , type: "number" },
				{ key: "left"          , value: "0"       , type: "number" },
				{ key: "pointer-events", value: "none"    , type: "value"  },
			]
		};
	}

};
</script>
<style lang="stylus">
@import "~style/functions"


.shima-code
	padding $base-spacing
	background: $base-color["bg-code"]

	&__content
		display block
		font-family Consolas, Monaco, "Andale Mono", "Ubuntu Mono", monospace
		color #A4B1D8

	div
		shima-font-size(-2)

	.token.property
		margin-left $base-spacing
		color #5C87FF

	.token.selector
		color #E15555

	.token.value
		color #667788

	.token.number
		color #E1773F

	.token.url, .token.string
		color #0AA395
</style>
