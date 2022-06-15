<template>
	<li
		:class="{ 'example-item--active': modelValue === type }" class="example-item"
		@click="setExample"
	>
		<a :href="'#example-' + type" class="example-item__link">
			<cube :type="cube" class="example-item__cube" /><slot />
		</a>
	</li>
</template>
<script setup>
import Cube from "src/component/Cube.vue";


const emit  = defineEmits(["update:modelValue"]);
const props = defineProps({
	modelValue: { type: String, default: "vr" },
	type:       { type: String, default: "vr" },
	cube:       { type: Number, default: 0 }
});


// ### [Methods] ########################################

function setExample() {
	emit("update:modelValue", props.type);
}
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.example-item
	position relative
	display flex
	align-items center
	padding 0 $base-spacing--xs 0 0
	margin-bottom $base-spacing--xs
	overflow hidden
	shima-font-size(-2)
	color #678
	cursor pointer
	transition all 300ms ease-in-out

	&:after
		shima-absolute top right bottom left
		z-index -1
		content ""
		background #DDD
		border-radius 5px
		opacity 0
		transition all 300ms ease-in-out
		transform translateX(-100%)

	&:hover:after, &--active:after
		opacity 1
		transform translateX(0)

	&__link
		display flex
		flex-grow 1
		align-items center
		color inherit
		text-decoration none

	&__cube
		margin-right $base-spacing--sm

</style>
