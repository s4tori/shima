<template>
	<block id="api-axis" title="Axis configuration">
		<block-url url>
			<em>grid</em> x <em>division</em> x <em>gutter</em> x <em>&lt;grid color&gt;</em> x <em>&lt;div color&gt;</em>
		</block-url>

		<block-option title="grid">
			<p>
				Set the spacing of the main grid.<br />
				Mandatory.
			</p>
		</block-option>

		<block-option title="division">
			<p>
				Number of divisions per grid.<br />
				Mandatory.
			</p>
		</block-option>

		<block-option title="gutter">
			<p>
				Spacing between each grid (you can set a negative value to start the grid by a gutter).<br />
				Mandatory.
			</p>
		</block-option>

		<block-option title="grid color">
			<p>
				Set the color of the grid lines.<br />
				Optional.<br />
				<a href="#api-color" class="api__def">See definition<icon i="arrow" /></a>
			</p>
		</block-option>

		<block-option title="div color">
			<p>
				Set the color of the subdivision lines.<br />
				Optional.<br />
				<a href="#api-color" class="api__def">See definition<icon i="arrow" /></a>
			</p>
		</block-option>
	</block>
</template>
<script setup>
import Block       from "./Block.vue";
import BlockUrl    from "./BlockUrl.vue";
import BlockOption from "./BlockOption.vue";
import Icon        from "src/component/Icon.vue";
</script>
