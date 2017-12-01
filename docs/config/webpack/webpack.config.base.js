const path                 = require("path");
const WebpackCleanupPlugin = require("webpack-cleanup-plugin");
const CopyWebpackPlugin    = require("copy-webpack-plugin");
const dir                  = require("../dir");


module.exports = {

	name: "client",
	context: dir.root,

	entry: [
		"./src/client-entry.js"
	],

	output: {
		filename: "client.js",
		path: dir.dist + "/assets/",
		publicPath: "assets/"
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
		new WebpackCleanupPlugin(),

		new CopyWebpackPlugin([
			{ from: path.resolve(dir.stat, "humans.txt")       , to : dir.dist },
			{ from: path.resolve(dir.css , "img/favicon/*.ico"), to : dir.dist, flatten: true },
			{ from: path.resolve(dir.css , "img/favicon/*.png"), to : path.resolve(dir.dist, "assets/image/favicon"), flatten: true },
			{ from: path.resolve(dir.stat, "og-image.png")     , to : path.resolve(dir.dist, "assets/image"), flatten: true }
		])
	]

};
