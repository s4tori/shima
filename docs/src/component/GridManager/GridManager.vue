<template>
	<div :id="gid" class="grid">
		<shima-top-bar :url="url" />
		<form class="grid__form">
			<div class="grid__row grid__row--top">
				<div class="grid__cell grid__cell--axis">
					<p>&nbsp;</p>
					<p>X axis</p>
					<p>Y axis</p>
				</div>
				<div class="grid__cell">
					<p>Spacing</p>
					<input v-model.number="x.grid" :placeholder="y.grid" class="grid__input grid__input--1" type="number" min="0" />
					<input v-model.number="y.grid" :placeholder="x.grid" class="grid__input grid__input--1" type="number" min="0" />
				</div>
				<div class="grid__cell">
					<p>Divisions</p>
					<input v-model.number="x.div" :placeholder="y.div" class="grid__input grid__input--2" type="number" min="0" />
					<input v-model.number="y.div" :placeholder="x.div" class="grid__input grid__input--2" type="number" min="0" />
				</div>
				<div class="grid__cell">
					<p>Gutter</p>
					<input v-model.number="x.gutter" :placeholder="y.gutter" class="grid__input grid__input--3" type="number" />
					<input v-model.number="y.gutter" :placeholder="x.gutter" class="grid__input grid__input--3" type="number" />
				</div>
				<div class="grid__cell--break" />
				<div class="grid__cell grid__cell--gutter grid__cell--first-gutter">
					<p>Grid color</p>
					<color-picker v-model="x.gridColor" />
					<color-picker v-model="y.gridColor" />
				</div>
				<div class="grid__cell grid__cell--gutter">
					<p>Div color</p>
					<color-picker v-model="x.divColor" />
					<color-picker v-model="y.divColor" />
				</div>
				<div class="grid__cell grid__cell--gutter">
					<p>Grid type</p>
					<shima-grid-picker v-model="type" />
				</div>
			</div>
		</form>
		<div class="grid__main">
			<div :style="style" class="grid__content">
				<slot />
			</div>
		</div>
		<div class="grid__code">
			<shima-code :url="url" :top="gridOnTop" type="css">
				<shima-checkbox v-model="gridOnTop" class="grid__cb">Grid on top</shima-checkbox>
			</shima-code>
		</div>
	</div>
</template>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.grid
	margin 0
	margin-bottom shima-vr(3)
	background white
	position relative
	border-bottom 5px solid #EAEAEA

	&__main
		border-top 1px solid #EEEEEE
		background-color #FAFAFA
		padding $base-spacing
		//box-shadow inset 0px 0px 30px #ECECEC

	&__content
		position relative
		box-shadow inset 0px 0px 30px #ECECEC
		box-shadow inset 0px 0px 1px #EEEEEE
		min-height shima-vr(10) - 1
		overflow hidden

		&:before, &:after
			content ""
			position absolute
			top 0
			right 0
			bottom 0
			left 0
			pointer-events none

		&:before
			border 1px dashed #ECECEC

		&:after
			display none

	&__form
		margin $base-spacing--sm shima-gutter()
		color #667788

		p
			white-space nowrap
			overflow hidden
			text-overflow ellipsis
			shima-font-size(-2)

	&__row
		display flex
		flex-wrap wrap

		// Alignment per row
		&--top
			align-items flex-start
		&--bottom
			align-items flex-end
		&--center
			align-items center

	// https://philipwalton.github.io/solved-by-flexbox/demos/grids/
	// https://github.com/suitcss/suit/blob/master/doc/naming-conventions.md#components
	&__cell
		flex 1

		&--1of2
			flex: 0 0 50%
		&--1of3
			flex: 0 0 33.3333%
		&--1of4
			flex: 0 0 25%

		&--gutter
			margin-left: shima-gutter()

		&--axis
			w = $grid-cell * 3
			flex  0 0 w
			width w

			p
				margin-bottom $base-spacing--sm

				&:first-child
					margin-bottom 0

	&__input
		display block
		width 100%
		-moz-appearance textfield
		border 1px solid #EEEEEE
		height: shima-vr()
		shima-font-size(-2, shima-vr())
		//top -1px
		margin 0 0 $base-spacing--sm 0
		padding 0 $base-spacing--xs
		outline 0
		$input-hover-active = $base-color["ip-active"]

		&::placeholder
			color #BCD

		&:focus
			border-color: $input-hover-active
			border-image: linear-gradient(90deg, $input-hover-active, rgba($input-hover-active, 0)) 1;

		&::-webkit-outer-spin-button,
		&::-webkit-inner-spin-button
			-webkit-appearance none
			margin 0

		&--1
			border-right none

		&--2
			border-left none
			border-right none
			&:focus
				border-image: linear-gradient(90deg, rgba($input-hover-active, 0), $input-hover-active, rgba($input-hover-active, 0)) 1;

		&--3
			border-left none
			&:focus
				border-image: linear-gradient(-90deg, $input-hover-active, rgba($input-hover-active, 0)) 1;

	&__code
		min-height 100px
		overflow hidden

	&__cb
		padding-top $base-spacing
		shima-font-size(-2, shima-vr())
		vertical-align: middle;

	+shima-mq("sm")
		&__cell
			&--axis
				display none

			&--break::before, &--break::after
				display: block;
				content: ""
				width 100%
				order: 1;

			&--break
				width: 100%

			&--first-gutter
				margin-left 0


</style>
<script>
import { _ }            from "src/util";
import { getEnvConfig } from "src/util";
import ShimaCheckbox    from "src/component/Checkbox.vue";
import ShimaTopBar      from "./TopBar.vue";
import ColorPicker      from "./ColorPicker.vue";
import ShimaGridPicker  from "./GridPicker.vue";
import ShimaCode        from "./Code.vue";


export default {

	components: {
		ColorPicker,
		ShimaTopBar,
		ShimaGridPicker,
		ShimaCheckbox,
		ShimaCode
	},

	props: {
		xGrid     : { type: Number, default: 128 }, yGrid     : { type: Number, default: null },
		xDiv      : { type: Number, default: 4   }, yDiv      : { type: Number, default: null },
		xGutter   : { type: Number, default: 32  }, yGutter   : { type: Number, default: null },
		xGridColor: { type: String, default: ""  }, yGridColor: { type: String, default: null },
		xDivColor : { type: String, default: ""  }, yDivColor : { type: String, default: null },
		t         : { type: String, default: "standard" },
		bg        : { type: String, default: ""  }
	},

	data() {
		return {
			gid: "grid-manager-" + this._uid,
			gUrl: "",
			base: getEnvConfig("urlGrid"),
			x: {
				grid: this.xGrid,
				div: this.xDiv,
				gutter: this.xGutter,
				gridColor: this.xGridColor, // FF00FF
				divColor: this.xDivColor
			},
			y: {
				grid: this.yGrid,
				div: this.yDiv,
				gutter: this.yGutter,
				gridColor: this.yGridColor || this.xGridColor, // FF00FF
				divColor: this.yDivColor || this.xDivColor
			},
			type: this.t,
			gridOnTop: false
		};
	},

	computed: {
		url() {
			const { x, y, base, type } = this;
			let nb = this.getNumber;
			let cl = this.getColor;

			const xGrid      = nb(x.grid  , y.grid);
			const xDiv       = nb(x.div   , y.div);
			const xGutter    = nb(x.gutter, y.gutter);
			const xDivColor  = cl(x.divColor);
			const xGridColor = cl(x.gridColor, xDivColor ? "CCC" : null);

			const yGrid      = nb(y.grid  , x.grid);
			const yDiv       = nb(y.div   , x.div);
			const yGutter    = nb(y.gutter, x.gutter);
			const yDivColor  = cl(y.divColor);
			const yGridColor = cl(y.gridColor, yDivColor ? "CCC" : null);

			let url = `${base}${xGrid}x${xDiv}x${xGutter}`;
			url += (xGridColor) ? `x${xGridColor}` : "";
			url += (xDivColor)  ? `x${xDivColor}`  : "";

			if (xGrid !== yGrid || xDiv !== yDiv || xGutter !== yGutter || xGridColor !== yGridColor) {
				url += `/${yGrid}x${yDiv}x${yGutter}`;
				url += yGridColor ? `x${yGridColor}` : "";
				url += yDivColor  ? `x${yDivColor}`  : "";
			}

			if (type !== "standard") {
				url += `/${type}`;
			}

			this.setGridUrl(url);
			return url;
		},

		style() {
			if (this.gridOnTop && !this.bg) { return null; }
			const st = { };
			if (this.bg)         { st.backgroundColor = this.bg; }
			if (!this.gridOnTop) { st.backgroundImage = `url('${this.gUrl}')`; }
			return st;
		}
	},

	watch: {
		gridOnTop: function (newValue) {
			newValue ? this.addGridOnTop(this.gUrl) : this.removeGridOnTop();
		}
	},


	// ### [Lifecycle] ######################################

	mounted() {
		this.setGridUrl = _.debounce(this.setGridUrl, 500, { leading: false, trailing: true });
	},

	beforeDestroy() {
		this.removeGridOnTop();
	},


	// ### [Methods] ########################################

	methods: {
		setGridUrl(url) {
			this.addGridOnTop(url);
			return this.gUrl = url;
		},

		getNumber(val, def) {
			return _.isNumber(val) ? val : def;
		},

		getColor(val, def) {
			return val.length > 0 ? val : def;
		},

		toggleGridOnTop(url) {
			this.gridOnTop ? this.addGridOnTop() : this.removeGridOnTop(url);
		},

		// Dynamically add CSS rule to an ::after element (to display the grid on top)
		// See: http://pankajparashar.com/posts/modify-pseudo-elements-css/
		// See: https://davidwalsh.name/add-rules-stylesheets
		addGridOnTop(url) {
			this.removeGridOnTop();
			if (!this.gridOnTop) { return; }
			const el = document.createElement("style");
			el.appendChild(document.createTextNode("")); // WebKit hack
			el.setAttribute("type", "text/css");
			el.setAttribute("id"  , "style-" + this.gid);
			document.head.appendChild(el);
			el.sheet.insertRule(`#${this.gid} .grid__content::after { display: block; background-image: url("${url}"); }`);
			return el.sheet;
		},

		removeGridOnTop() {
			const el = document.getElementById("style-" + this.gid);
			el && el.parentNode.removeChild(el);
		}
	}

};
</script>
