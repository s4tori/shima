const path                   = require("path");
const CssMinimizerPlugin     = require("css-minimizer-webpack-plugin");
const HtmlWebpackPlugin      = require("html-webpack-plugin");
const MiniCssExtractPlugin   = require("mini-css-extract-plugin");
const { isMinimal }          = require("std-env");
const TerserPlugin           = require("terser-webpack-plugin");
const webpack                = require("webpack");
const { mergeWithCustomize } = require("webpack-merge");
const { customizeObject }    = require("webpack-merge");
const WebpackBar             = require("webpackbar");
const define                 = require("./util/webpack.define");
const loaderOptions          = require("./util/webpack.loaders").getloaders("production");
const HtmlSsgPlugin          = require("./util/webpack.ssg.plugin");
const webConf                = require("../config")("production");
const dir                    = require("../dir");


const mergeStrategy = mergeWithCustomize({
	customizeObject: customizeObject({ "entry": "prepend" }),
});

module.exports = mergeStrategy(require("./webpack.config.base"), {

	mode: "production",

	// HtmlSsgPlugin: server side rendering
	entry: {
		ssg: "./src/entry-ssg.js"
	},

	output: {
		filename: "[name]-[contenthash:6].js"
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

		new MiniCssExtractPlugin({ filename: "style-[contenthash:6].css" }),

		new HtmlWebpackPlugin({
			template: path.resolve(dir.app, "./index.ejs"),
			filename: "../index.html",
			excludeChunks: ["ssg"],
			minify: {
				removeComments: true,
				collapseWhitespace: true
			},
			custom: {
				base: webConf.client.base
			}
		}),

		// Small plugin to handle Static Site Generation
		new HtmlSsgPlugin(),

		new WebpackBar({
			reporters: isMinimal ? ["basic"] : ["fancy", "profile"],
			profile: true
		})
	],

	stats: {
		modules: false
	},

	// Performance options
	// https://webpack.js.org/configuration/performance/#performancemaxentrypointsize
	performance: {
		maxEntrypointSize: 500000
	},

	optimization: {
		minimizer: [
			new CssMinimizerPlugin(),
			new TerserPlugin({ extractComments: false })
		],

		splitChunks: {
			// chunks: "all",
			chunks(chunk) {
				return chunk.name !== "ssg";
			},
			// https://stackoverflow.com/q/66986664
			// Keep default configuration for "default" and "defaultVendors" cache groups
			name: (module, chunks, cacheGroupKey) => {
				const allChunksNames = chunks.map((chunk) => chunk.name).join("~");
				const prefix = cacheGroupKey === "defaultVendors" ? "vendors" : cacheGroupKey;
				return `${prefix}~${allChunksNames}`;
			}
		}
	}
});
