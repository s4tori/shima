<template>
	<div class="grid-url">
		<div class="grid-url__value"  target="_blank">{{ url }}</div>
		<button
			v-clipboard:copy="url"
			:class="clsButton" class="grid-url__button"
			@copied="copied">
			{{ text }}
		</button>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.grid-url
	display flex
	align-items center
	border-bottom 1px solid #F2F2F2
	overflow hidden
	shima-font-size(-2, shima-vr() + $base-spacing--sm - 1)

	&__value
		shima-font-size(-2, shima-vr() + $base-spacing--sm - 1)
		margin-left shima-span()
		text-decoration none
		color black
		cursor pointer
		overflow hidden
		white-space nowrap
		text-overflow ellipsis

	&__button
		display inline-block
		width 110px
		margin-left auto
		margin-right $base-spacing
		padding 0 $base-spacing--sm
		border 1px solid #EEEEEE
		border-radius 3px
		background white
		cursor pointer
		shima-font-size(-2, shima-vr())
		outline none
		transition all .3s ease-in-out

		@keyframes pulse
			0%
				box-shadow: 0 0 0 0 $base-color["cb-active"]
				background-color white

		&.button--active
			animation: pulse 1s;
			box-shadow: 0 0 0 20px rgba(255, 255, 255, 0)
			background-color: $base-color["cb-active"];
			color white
			pointer-events none

		&:hover
			border-color: $base-color["cb-active"]

	+shima-mq("md-and-dn")
		justify-content center

		&__value
			margin-left $base-spacing--sm

		&__button
			display none

</style>
<script>
export default {

	props: {
		url: { type: String, default: "" }
	},

	data() {
		return {
			text: "Copy URL",
			clsButton: ""
		};
	},


	// ### [Methods] ########################################

	methods: {

		copied() {
			this.text      = "Copied!";
			this.clsButton = "button--active";

			setTimeout(() => {
				this.text      = "Copy URL";
				this.clsButton = "";
			}, 1500);
		}

	}

};
</script>
