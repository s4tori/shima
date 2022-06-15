const path              = require("path");
const webpack           = require("webpack");
const { merge }         = require("webpack-merge");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const define            = require("./util/webpack.define");
const loaderOptions     = require("./util/webpack.loaders").getloaders("development");
const webConf           = require("../config")("development");
const dir               = require("../dir");


module.exports = merge(require("./webpack.config.base"), {

	mode: "development",
	devtool: "eval-source-map",

	devServer: {
		compress: true,
		port: webConf.server.port,
		open: webConf.devServer.open,
		historyApiFallback: true,
		hot: true,
		// parent index.html
		// see: https://github.com/webpack/webpack-dev-server/issues/4231
		devMiddleware: {
			publicPath: "auto",
			writeToDisk: false
		},
		host: process.env.HOST || "localhost",
		client: {
			logging: "none",
			overlay: true
		}
	},

	module: {
		rules: [
			{ test: /\.styl/, include: [dir.app], use: [
				"vue-style-loader",
				loaderOptions.css({ modules: false, importLoaders: 2 }),
				loaderOptions.postCss(),
				loaderOptions.stylus(),
			] }
		]
	},

	plugins: [
		new webpack.DefinePlugin(define({
			env: "development"
		})),

		// HMR
		new webpack.WatchIgnorePlugin({
			paths: [path.resolve(__dirname, "./../../node_modules/")]
		}),

		new HtmlWebpackPlugin({
			template: path.resolve(dir.app, "./index.ejs"),
			filename: "index.html",
			custom: {
				base: webConf.client.base
			}
		})
	]

});
