/* eslint-disable no-undef, no-console */
// __resourceQuery = "?noInfo=true&reload=true<...>"
const hotClient = require(`webpack-hot-middleware/client${__resourceQuery}`);

// We create a custom action to process a complete reload
// See server/index.dev.js
hotClient.subscribe(function (event) {
	if (event.action === "reload") {
		console.log("[HMR] [Custom] Reload the page");
		window.location.reload();
	}
});
