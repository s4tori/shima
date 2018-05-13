const path              = require("path");
const webpack           = require("webpack");
const WebpackConfig     = require("webpack-config").default;
const HtmlWebpackPlugin = require("html-webpack-plugin");
const define            = require("./util/webpack.define");
const loaderOptions     = require("./util/webpack.loaders").getloaders("development");
const webConf           = require("../config")("development");
const dir               = require("../dir");


module.exports = new WebpackConfig().extend("./config/webpack/webpack.config.base.js").merge({

	mode: "development",

	module: {
		rules: [
			{ test: /\.styl/, include: [dir.app], use: ["vue-style-loader",
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
		new webpack.WatchIgnorePlugin([path.resolve(__dirname, "./../../node_modules/")]),
		new webpack.NoEmitOnErrorsPlugin(),

		new HtmlWebpackPlugin({
			template: path.resolve(dir.app, "./index.ejs"),
			filename: "index.html",
			custom: {
				base: webConf.client.base
			}
		})
	],

	// Performance options: disable warnings ([HMR] asset size limit)
	performance: { hints: false },

	devtool: "#cheap-module-source-map"

});
