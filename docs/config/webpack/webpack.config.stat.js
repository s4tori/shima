const { BundleAnalyzerPlugin } = require("webpack-bundle-analyzer");
const { merge }                = require("webpack-merge");


module.exports = merge(require("./webpack.config.prod"), {

	plugins: [
		new BundleAnalyzerPlugin()
	]

});
