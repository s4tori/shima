<template>
	<footer class="footer">
		<logo class="footer__logo" :toggle-ee="ee" />
		<p class="footer__about">
			<span>
				<span>&lt;/&gt;</span>
				with <a rel="nofollow" href="http://hashboxed.com/loveyourcode/" class="footer__heart">♡</a>
				by
			</span>
			<span class="footer__ee" @click="ee = !ee">547021</span> |
			<a :href="urlSource"><icon i="github" /></a>
		</p>
	</footer>
</template>

<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"

.footer
	background #111

	&__logo
		width shima-span(4)
		max-width 100%
		margin 0 auto
		padding ($base-spacing--xxl * 2) $base-spacing--sm

		+shima-mq("sm")
			padding $base-spacing--xxl $base-spacing--sm

	&__about
		text-align center
		user-select none
		color: $base-color["blue"]
		shima-font-size(-2)
		padding-bottom $base-spacing

		@keyframes shima-heart {
			  0% { transform: scale3d(2  , 2  , 1); opacity:1;  }
			 10% { transform: scale3d(3  , 3  , 1); opacity:1;  }
			 30% { transform: scale3d(2.5, 2.5, 1); opacity:.8; }
			 40% { transform: scale3d(3  , 3  , 1); opacity:1;  }
			 80% { transform: scale3d(2  , 2  , 1); opacity:1;  }
			100% { transform: scale3d(2  , 2  , 1); opacity:1;  }
		}

		span, a
			reset-link()
			transition: all 300ms ease-in-out;
			color: #666;

	&__about > span:hover &__heart
		color: red;
		padding: 0 shima-vr();
		animation: shima-heart 1.5s infinite;

	&__heart
		display: inline-block;
		cursor: pointer;

	&__ee
		cursor pointer

</style>
<script>
import Logo             from "./Logo.vue";
import Icon             from "src/component/Icon.vue";
import { getEnvConfig } from "src/util";


export default {

	data: function () {
		return {
			pid: "pattern-footer-",
			w: 20,
			h: 20,
			ee: false,
			urlSource: getEnvConfig("urlSource")
		};
	},

	components: {
		Logo,
		Icon
	}

};
</script>
