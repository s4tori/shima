import Clipboard  from "./clipboard";


export default {
	install: Clipboard
};

// check if HMR is enabled
if (module.hot) {
	// accept itself
	module.hot.accept();
}
