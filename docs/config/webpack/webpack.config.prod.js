const path                 = require("path");
const chalk                = require("chalk");
const webpack              = require("webpack");
const WebpackConfig        = require("webpack-config").default;
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const OptCSSAssetsPlugin   = require("optimize-css-assets-webpack-plugin");
const ProgressBarPlugin    = require("progress-bar-webpack-plugin");
const HtmlWebpackPlugin    = require("html-webpack-plugin");
const PrerenderSpaPlugin   = require("prerender-spa-plugin");
const define               = require("./util/webpack.define");
const loaderOptions        = require("./util/webpack.loaders").getloaders("production");
const webConf              = require("../config")("production");
const dir                  = require("../dir");


module.exports = new WebpackConfig().extend("./config/webpack/webpack.config.base.js").merge({

	mode: "production",

	output: {
		filename: "[name]-[chunkhash:6].js"
	},

	module: {
		rules: [
			{ test: /\.styl/, include: [dir.app], use: [
				MiniCssExtractPlugin.loader,
				loaderOptions.css({ modules: false, importLoaders: 2 }),
				loaderOptions.postCss(),
				loaderOptions.stylus(),
			] }
		]
	},

	plugins: [
		new webpack.DefinePlugin(define({
			env: "production"
		})),

		// CSS
		new MiniCssExtractPlugin({ filename: "style-[contenthash:6].css" }),
		new OptCSSAssetsPlugin({ }),

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
	],

	// Performance options
	// https://webpack.js.org/configuration/performance/#performance-maxentrypointsize
	performance: {
		maxEntrypointSize: 500000
	},

	// "splitChunks" doesn't work very well with vue-loader yet
	// https://github.com/webpack-contrib/mini-css-extract-plugin/issues/113
	// https://github.com/webpack-contrib/mini-css-extract-plugin/issues/85
	optimization: {
		splitChunks: {
			cacheGroups: {
				default: false,
				styles: {
					name: "styles",
					test: /\.(css|styl)$/,
					chunks: "all",
					//enforce: true
				},
				commons: {
					test: /[\\/]node_modules[\\/]/,
					name: "vendor",
					priority: -10,
					chunks: "initial"
				}
			}
		},
		minimize: true
	}

});
