<template>
	<footer class="footer">
		<logo :toggle-ee="ee" class="footer__logo" />
		<p class="footer__about">
			<span>
				<span>&lt;/&gt;</span>
				with <a rel="nofollow" href="https://web.archive.org/web/20210417162557/http://hashboxed.com/loveyourcode/" class="footer__heart">♡</a>
				by
			</span>
			<span class="footer__ee" @click="ee = !ee">547021</span> |
			<a :href="urlSource"><icon i="github" /></a>
		</p>
	</footer>
</template>
<script>
import Logo             from "./Logo.vue";
import Icon             from "src/component/Icon.vue";
import { getEnvConfig } from "src/util";


export default {

	components: {
		Logo,
		Icon
	},

	data: function () {
		return {
			pid: "pattern-footer-",
			w: 20,
			h: 20,
			ee: false,
			urlSource: getEnvConfig("urlSource")
		};
	}

};
</script>
<style lang="stylus">
@import "~style/functions"
@import "~style/mixins"


.footer
	background #111

	&__logo
		width shima-span(4)
		max-width 100%
		padding ($base-spacing--xxl * 2) $base-spacing--sm
		margin 0 auto

		+shima-mq("sm")
			padding $base-spacing--xxl $base-spacing--sm

	&__about
		shima-font-size(-2)
		padding-bottom $base-spacing
		color: $base-color["blue"]
		text-align center
		user-select none

		@keyframes shima-heart
			  0% { transform: scale3d(2  , 2  , 1); opacity:1;  }
			 10% { transform: scale3d(3  , 3  , 1); opacity:1;  }
			 30% { transform: scale3d(2.5, 2.5, 1); opacity:.8; }
			 40% { transform: scale3d(3  , 3  , 1); opacity:1;  }
			 80% { transform: scale3d(2  , 2  , 1); opacity:1;  }
			100% { transform: scale3d(2  , 2  , 1); opacity:1;  }


		span, a
			reset-link()
			color #666
			transition all 300ms ease-in-out

	&__about > span:hover &__heart
		padding 0 shima-vr()
		color red
		animation shima-heart 1.5s infinite

	&__heart
		display inline-block
		cursor pointer

	&__ee
		cursor pointer

</style>
