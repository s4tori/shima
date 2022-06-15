import Clipboard from "clipboard";
import { _ }     from "src/util";


// Example :
// <code>
// <button v-clipboard="myDynamicValue" @copied="copied">Text</button>
//</code>
export default function (app) {

	app.directive("clipboard", {
		beforeMount(el, binding, vnode) {
			el._cbText = binding.value;
			el._cbElem = new Clipboard(el, {
				text: () => el._cbText
			});

			const handler = _.get(vnode, "props.onCopied");
			el._cbElem.on("success", () => {
				handler && handler();
			});
		},

		updated(el, binding) {
			el._cbText = binding.value;
		},

		unmounted(el) {
			el._cbElem && el._cbElem.destroy();
		}
	});

}
