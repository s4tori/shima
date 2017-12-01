<template>
	<div class="shima-code">
		<code class="shima-code__content">
			<div v-if="type === 'css' && !top">
				<span class="token selector">.myAwesomeDiv</span> {<br />
				<span class="token property">background-image</span>:
				<span class="token value">url(<span class="token url">"{{url}}"</span>)</span>;<br />
				}
			</div>
			<div v-else-if="type === 'css' && top">
				<span class="token selector">.myAwesomeDiv</span> {<br />
				<span class="token property">position</span>: <span class="token value">relative</span>;<br />
				}<br />
				<span class="token selector">.myAwesomeDiv::after</span> {<br />
				<template v-for="(item, index) in cssTop">
					<span :key="index">
						<span class="token property">{{ item.key }}</span>: <span class="token" :class="item.type">{{ item.value }}</span>;
						<br />
					</span>
				</template>
				<span class="token property">background-image</span>:
				<span class="token value">url(<span class="token url">"{{url}}"</span>)</span>;<br />
				}
			</div>
			<div v-else-if="type === 'stylus'">
				<span class="token selector">.foo</span><br />
				<span class="token property">background-image</span>
				<span class="token value">url(<span class="token url">"{{url}}"</span>)</span><br />
			</div>
		</code>
		<slot />
	</div>
</template>
<style lang="stylus">
@import "~style/functions"


.shima-code
	background: $base-color["bg-code"]
	padding $base-spacing

	&__content
		display block
		font-family Consolas, Monaco, "Andale Mono", "Ubuntu Mono", monospace
		color #a4b1d8

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
		color #e1773f

	.token.url, .token.string
		color #0AA395

</style>
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
