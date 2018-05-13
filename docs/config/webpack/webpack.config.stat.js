const BundleAnalyzerPlugin = require("webpack-bundle-analyzer").BundleAnalyzerPlugin;
const WebpackConfig        = require("webpack-config").default;


module.exports = new WebpackConfig().extend("./config/webpack/webpack.config.dev.js").merge({

	plugins: [
		new BundleAnalyzerPlugin()
	]

});
