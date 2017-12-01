const path              = require("path");
const webpack           = require("webpack");
const WebpackConfig     = require("webpack-config").default;
const HtmlWebpackPlugin = require("html-webpack-plugin");
const define            = require("./util/webpack.define");
const vueOptions        = require("./util/webpack.vue").getVueConfig("dev");
const webConf           = require("../config")("development");
const dir               = require("../dir");


module.exports = new WebpackConfig().extend("./config/webpack/webpack.config.base.js").merge({

	entry: [
		"./config/webpack/util/whm-custom-client?noInfo=false"
	],

	plugins: [
		// HMR
		new webpack.WatchIgnorePlugin([path.resolve(__dirname, "./../../node_modules/")]),
		new webpack.HotModuleReplacementPlugin(),
		new webpack.NoEmitOnErrorsPlugin(),

		new webpack.LoaderOptionsPlugin(vueOptions),
		new webpack.LoaderOptionsPlugin({
			options: {
				stylus: {
					define: {
						__DEV__: true,
						__URL__: webConf.client.urlGrid
					}
				}
			}
		}),

		new HtmlWebpackPlugin({
			template: path.resolve(dir.app, "./index.ejs"),
			filename: "index.html",
			custom: {
				base: webConf.client.base
			}
		}),

		new webpack.DefinePlugin(define({
			env: "development"
		}))

	],

	devtool: "#cheap-module-source-map"

});
