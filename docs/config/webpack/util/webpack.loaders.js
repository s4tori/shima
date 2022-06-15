const postcssPresetEnv = require("postcss-preset-env");


class Loaders {

	constructor(devMode = false) {
		this.devMode = devMode;
		this.webConf = require("../../config")(devMode ? "development" : "production");
	}

	css({ modules = true, importLoaders = 1 } = { }) {
		return {
			loader: "css-loader",
			options: {
				sourceMap: this.devMode,
				// modules,
				importLoaders,
				// https://github.com/webpack-contrib/css-loader#modules
				modules: modules ? {
					modules: {
						localIdentName: "[local]_[hash:base64:5]",
						exportLocalsConvention: "camelCase"
					}
				} : false
			}
		};
	}

	stylus() {
		return {
			loader: "stylus-loader",
			options: {
				sourceMap: this.devMode,
				//define: {
				//	__DEV__: this.devMode,
				//	__URL__: this.webConf.client.urlGrid
				//},
				// https://github.com/webpack-contrib/stylus-loader#object
				stylusOptions: {
					define: [
						["__DEV__", this.devMode],
						["__URL__", this.webConf.client.urlGrid]
					]
				}
			}
		};
	}

	postCss() {
		return {
			loader: "postcss-loader",
			options: {
				sourceMap: this.devMode,
				postcssOptions: {
					plugins: [postcssPresetEnv({
						browsers: "last 2 versions",
					})]
				},
			},
			// options: { sourceMap: this.devMode, plugins: () => [autoprefixer("last 3 versions")] }
		};
	}

}

module.exports = {

	getloaders(type) {
		return new Loaders(type !== "production");
	}

};
