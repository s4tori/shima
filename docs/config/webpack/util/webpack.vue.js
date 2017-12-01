const _                 = require("lodash");
const ExtractTextPlugin = require("extract-text-webpack-plugin");


module.exports = {

	getVueConfig(type) {
		return {
			vue: _.merge({}, this.global, this[type === "production" ? "prod" : "dev"])
		};
	},

	global: {
		postcss: [require("autoprefixer")({ browsers: ["last 3 versions"] })],
		cssModules: {
			localIdentName: "[local]-[hash:base64:5]",
			camelCase: true
		}
	},

	dev: {

	},

	prod: {
		loaders: {
			css   : ExtractTextPlugin.extract({ use: "css-loader"                          , fallback: "vue-style-loader" }),
			stylus: ExtractTextPlugin.extract({ use: "css-loader!stylus-loader?resolve url", fallback: "vue-style-loader" })
		}
	}

};
