<template>
	<div id="grid-example" class="grid-example">
		<component :is="component" />
	</div>
</template>
<script setup>
import { computed }   from "vue";
import GridManager    from "src/component/GridManager/GridManager.vue";
import VerticalRhythm from "src/component/GridManagerExample/VerticalRhythm.vue";
import Layout         from "src/component/GridManagerExample/Layout.vue";
import VRGrid         from "src/component/GridManagerExample/VRGrid.vue";
import Blueprint      from "src/component/GridManagerExample/BluePrint.vue";
import Icon           from "src/component/GridManagerExample/Icon.vue";
import Iso            from "src/component/GridManagerExample/Iso.vue";
import Oblique        from "src/component/GridManagerExample/Oblique.vue";
import StripeRight    from "src/component/GridManagerExample/StripeRight.vue";
import StripeLeft     from "src/component/GridManagerExample/StripeLeft.vue";


const props = defineProps({
	type: { type: String, default: "vr" }
});

const component = computed(() => {
	switch (props.type) {
		case "vr"      : return VerticalRhythm;
		case "layout"  : return Layout;
		case "vrgrid"  : return VRGrid;
		case "bp"      : return Blueprint;
		case "icon"    : return Icon;
		case "iso"     : return Iso;
		case "oblique" : return Oblique;
		case "stripe-r": return StripeRight;
		case "stripe-l": return StripeLeft;
		default        : return GridManager;
	}
});
</script>
