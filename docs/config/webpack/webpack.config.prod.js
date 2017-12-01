const path               = require("path");
const chalk              = require("chalk");
const webpack            = require("webpack");
const WebpackConfig      = require("webpack-config").default;
const ExtractTextPlugin  = require("extract-text-webpack-plugin");
const ProgressBarPlugin  = require("progress-bar-webpack-plugin");
const HtmlWebpackPlugin  = require("html-webpack-plugin");
const PrerenderSpaPlugin = require("prerender-spa-plugin");
const define             = require("./util/webpack.define");
const vueOptions         = require("./util/webpack.vue").getVueConfig("production");
const webConf            = require("../config")("production");
const dir                = require("../dir");


module.exports = new WebpackConfig().extend("./config/webpack/webpack.config.base.js").merge({

	output: {
		filename: "bundle-[chunkhash:6].js"
	},

	plugins: [
		new ExtractTextPlugin("style-[contenthash:6].css"),
		new webpack.LoaderOptionsPlugin({ minimize: true }),
		new webpack.LoaderOptionsPlugin(vueOptions),
		new webpack.LoaderOptionsPlugin({
			options: {
				stylus: {
					define: {
						__DEV__: false,
						__URL__: webConf.client.urlGrid
					}
				}
			}
		}),

		new webpack.optimize.CommonsChunkPlugin({
			name: "vendor",
			filename: "[name]-[chunkhash:6].js",
			minChunks: (module) => {
				// See: https://webpack.js.org/plugins/commons-chunk-plugin/#passing-the-minchunks-property-a-function
				if (module.resource && (/^.*\.(css|scss)$/).test(module.resource)) {
					return false;
				}

				return module.resource && (module.resource.indexOf(dir.app) === -1);
			}
		}),

		new webpack.DefinePlugin(define({
			env: "production"
		})),

		new webpack.optimize.UglifyJsPlugin({ compress: { warnings: false } }),

		new HtmlWebpackPlugin({
			template: path.resolve(dir.app, "./index.ejs"),
			filename: "../index.html",
			minify: {
				removeComments: true,
				collapseWhitespace: true
			},
			custom: {
				base: webConf.client.base
			}
		}),

		new PrerenderSpaPlugin(
			dir.dist,
			["/"]
		),

		new ProgressBarPlugin({
			clear: true,
			summary: false,
			format: `${chalk.green.gray("[build]")} [:bar] ${chalk.green.bold(":percent")} | :msg`,
			customSummary: ((bTime) => console.log("Build completed in", bTime))  // eslint-disable-line no-console
		})
	]

});
