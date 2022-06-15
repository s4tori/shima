<template>
	<div class="grid-url">
		<div class="grid-url__value">{{ url }}</div>
		<button
			v-clipboard:copy="url"
			:class="clsButton" class="grid-url__button"
			@copied="copied"
		>
			{{ text }}
		</button>
	</div>
</template>
<script setup>
import { onBeforeUnmount, ref } from "vue";


defineProps({
	url: { type: String, default: "" }
});

let text       = ref("Copy URL");
let clsButton  = ref("");
let timer      = null;


// ### [Lifecycle] ######################################

onBeforeUnmount(() => {
	timer && clearTimeout(timer);
});


// ### [Methods] ########################################

function copied() {
	text.value      = "Copied!";
	clsButton.value = "button--active";

	timer = setTimeout(() => {
		text.value      = "Copy URL";
		clsButton.value = "";
	}, 1500);
}
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.grid-url
	display flex
	align-items center
	overflow hidden
	border-bottom 1px solid #F2F2F2
	shima-font-size(-2, shima-vr() + $base-spacing--sm - 1)

	&__value
		margin-left shima-span()
		overflow hidden
		shima-font-size(-2, shima-vr() + $base-spacing--sm - 1)
		color black
		text-decoration none
		text-overflow ellipsis
		white-space nowrap
		cursor pointer

	&__button
		display inline-block
		width 110px
		padding 0 $base-spacing--sm
		margin-right $base-spacing
		margin-left auto
		shima-font-size(-2, shima-vr())
		cursor pointer
		background white
		border 1px solid #EEEEEE
		border-radius 3px
		outline none
		transition all 0.3s ease-in-out

		@keyframes pulse
			0%
				background-color white
				box-shadow: 0 0 0 0 $base-color["cb-active"]

		&.button--active
			pointer-events none
			box-shadow 0 0 0 20px rgba(255, 255, 255, 0)
			animation pulse 1s

		&:hover
			border-color: $base-color["cb-active"]

	+shima-mq("md-and-dn")
		justify-content center

		&__value
			margin-left $base-spacing--sm

		&__button
			display none

</style>
