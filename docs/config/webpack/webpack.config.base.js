const path                   = require("path");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const CopyWebpackPlugin      = require("copy-webpack-plugin");
const { VueLoaderPlugin }    = require("vue-loader");
const dir                    = require("../dir");


module.exports = {

	name: "client",
	context: dir.root,

	entry: {
		main: "./src/entry-client.js"
	},

	output: {
		filename: "client.js",
		path: dir.dist + "/assets/",
		publicPath: "auto"
	},

	resolve: {
		alias: { src: dir.app, style: dir.css, font: dir.font }
	},

	module: {
		rules: [
			{ test: /\.vue$/, use: ["vue-loader"]  , include: [dir.app] },
			{ test: /\.js$/ , use: ["babel-loader"], include: [dir.app] }
		]
	},

	plugins: [
		new VueLoaderPlugin(),

		new CleanWebpackPlugin({
			cleanOnceBeforeBuildPatterns: [
				"**/*",
				path.resolve(dir.dist, "index.html"),
				path.resolve(dir.dist, "humans.txt")
			],
			dry: false,
			verbose: false
		}),

		new CopyWebpackPlugin({
			patterns: [
				{
					from: path.resolve(dir.stat, "humans.txt"),
					to : dir.dist
				},
				{
					from: path.resolve(dir.stat, "og-image.png"),
					to : path.resolve(dir.dist, "assets/image/[name][ext]") // Flatten copy
				},
				{
					from: path.resolve(dir.css , "img/favicon/*.png").replace(/\\/g, "/"),
					to : path.resolve(dir.dist, "assets/image/favicon/[name][ext]")
				}
			]
		})
	]

};
