const path     = require("path");
const distRoot = path.resolve(__dirname, "../dist/");


const htmlManager = {
	html: "",
	setHtml(html) { this.html = html; return html; },
	getHtml()    { return this.html; },
	exportHtml() { return this.getHtml.bind(this); },
};

// Load dev/hot middlewares / retrieve HTML file
module.exports = (app) => {
	if (process.env.NODE_ENV !== "development") {
		app.use(require("koa-static")(distRoot));
		const html = require("fs").readFileSync(path.resolve(distRoot, "./index.html"));
		return Promise.resolve(() => html.toString());
	}

	const koaWebpack    = require("koa-webpack");
	const webpackConfig = require("../config/webpack/webpack.config.dev");
	const webpack       = require("webpack");
	const compiler      = webpack(webpackConfig);
	const middlewares   = koaWebpack({
		compiler,
		dev: {
			publicPath: "/" + webpackConfig.output.publicPath,
			noInfo: false,
			quiet: false,
			stats: {
				colors: true,
				modules: false
			}
		},
		hot: {
			reload: false
		}
	});
	app.use(middlewares);

	// Force page reload when a file processed by html-webpack-plugin changes
	compiler.hooks.compilation.tap("HmrHtml", (compilation) => {
		compilation.hooks.htmlWebpackPluginAfterEmit.tapAsync("HmrHtml", (data, cb) => {
			// html-webpack-plugin always emits HMR event for each change (cache issue)
			// See: https://github.com/jantimon/html-webpack-plugin/issues/680
			// See: https://github.com/webpack-contrib/webpack-hot-client#communicating-with-client-websockets
			if (htmlManager.html !== data.html.source()) {
				htmlManager.setHtml(data.html.source());
				// eslint-disable-next-line no-console
				console.log("info", "[build   ]", "HTML page updated, call window-reload socket event...");
				middlewares.client.wss.broadcast(JSON.stringify({ type: "window-reload" }));
			}

			cb(null, data);
		});
	});

	// Get HTML file
	return new Promise((resolve, reject) => {
		middlewares.dev.waitUntilValid((stats) => {
			if (stats.hasErrors() || stats.hasWarnings()) {
				return reject(stats.toJson()[stats.toJson().errors.length >= 1 ? "errors" : "warnings"]);
			}

			// eslint-disable-next-line no-console
			console.log("info", "[build   ]", "First webpack build done, get index.html content...");
			htmlManager.setHtml(stats.compilation.assets["index.html"].source());
			resolve(htmlManager.exportHtml());
		});
	});

};
