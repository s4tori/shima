<template>
	<li
		class="example-item" :class="{ 'example-item--active': curr === type }"
		@click="setExample"
	>
		<a class="example-item__link" :href="'#example-' + type">
			<cube class="example-item__cube" :type="cube" /><slot></slot>
		</a>
	</li>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.example-item
	position relative
	display flex
	align-items center
	margin-bottom $base-spacing--xs
	padding 0 $base-spacing--xs 0 0
	shima-font-size(-2)
	color #678
	cursor pointer
	transition all 300ms ease-in-out
	overflow hidden

	&:after
		content ""
		absolute: top right bottom left
		background #DDD
		border-radius 5px
		z-index -1
		transition all 300ms ease-in-out
		transform translateX(-100%)
		opacity 0

	&:hover:after, &--active:after
		transform translateX(0)
		opacity 1

	&__link
		display flex
		align-items center
		color inherit
		text-decoration none

	&__cube
		margin-right $base-spacing--sm

</style>
<script>
import Cube from "src/component/Cube.vue";


export default {

	model: {
		prop: "curr",
		event: "input"
	},

	props: {
		curr: { type: String, default: "vr" },
		type: { type: String, default: "vr" },
		cube: { type: Number, default: 0    }
	},

	components: {
		Cube
	},

	methods: {
		setExample() {
			this.$emit("input", this.type);
		}
	}

};
</script>
