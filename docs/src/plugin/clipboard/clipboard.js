import Clipboard from "clipboard";
import { _ }     from "src/util";


// Example :
// <code>
// <button v-clipboard="myDynamicValue">Text</button>
//</code>
export default function (Vue) {

	Vue.directive("clipboard", {
		bind: function (el, binding, vnode) {
			el._cbText = binding.value;
			el._cbElem = new Clipboard(el, {
				text: () => el._cbText
			});

			const handler = _.get(vnode, "data.on.copied");

			el._cbElem.on("success", () => {
				// vnode.context.$emit("copied")
				handler && handler.fns();
			});
		},

		componentUpdated: function (el, binding) {
			el._cbText = binding.value;
		},

		unbind: function (el) {
			el._cbElem && el._cbElem.destroy();
		}
	});

}
