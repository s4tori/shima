const autoprefixer = require("autoprefixer");


class Loaders {

	constructor(devMode = false) {
		this.devMode = devMode;
		this.webConf = require("../../config")(devMode ? "development" : "production");
	}

	css({ modules = true, importLoaders = 1 } = { }) {
		return {
			loader: "css-loader",
			options: { sourceMap: this.devMode, modules, localIdentName: "[local]_[hash:base64:5]", camelCase: true, importLoaders }
		};
	}

	stylus() {
		return {
			loader: "stylus-loader",
			options: {
				sourceMap: this.devMode,
				define: {
					__DEV__: this.devMode,
					__URL__: this.webConf.client.urlGrid
				}
			}
		};
	}

	postCss() {
		return {
			loader: "postcss-loader",
			options: { sourceMap: this.devMode, plugins: () => [autoprefixer("last 3 versions")] }
		};
	}

}

module.exports = {

	getloaders(type) {
		return new Loaders(type !== "production");
	}

};
