const vm                = require("node:vm");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const _                 = require("lodash");


/**
 * HtmlWebpackPlugin plugin to handle Static Site Generation
 * Extract SSR html generated by Vue into <div id="app" /> HTML tag
 *
 * This plugin depends on "entry-ssg.js" file (see webpack configuration):
 * - add new entry to get SSR content: entry.ssg: "./src/entry-ssg.js"
 * - do not include entry-ssg.js in index.html: HtmlWebpackPlugin.excludeChunks: ["ssg"]
 * - do not split entry-ssg.js: splitChunks.chunks: chunk.name !== "ssg"
 * @see src/entry-ssg.js (renderToString function)
 * @see https://vuejs.org/guide/scaling-up/ssr.html#ssr-vs-ssg
 */
class HtmlSsgPlugin {

	constructor(options) {
		this.options = _.merge({
			chunkName: "ssg",
			tplKey: "options.custom.ssg",
			tagId: "app"
		}, options || { });
	}

	apply(compiler) {
		const { options } = this;

		compiler.hooks.compilation.tap("HtmlSsgPlugin", (compilation) => {

			HtmlWebpackPlugin.getHooks(compilation).beforeEmit.tapAsync(
				"HtmlSsgPlugin",
				(data, cb) => {
					// Extract code source (from "./src/entry-ssg.js")
					const fileReg = new RegExp(`^${options.chunkName}-.*\\.js$`);
					const fileKey = Object.keys(compilation.assets).find(file => fileReg.test(file));
					const asset   = compilation.getAsset(fileKey);
					const source  = asset.source.source();

					// Prevent webpack to write SSG asset into output folder
					compilation.deleteAsset(fileKey);

					// Execute code source (get result from "renderToString" vue function)
					// The result will be written in <div id="app" /> tag (see index.ejs)
					const ctx = vm.createContext({ _ssg: { } });
					new Promise((res, rej) => {
						ctx._ssg.resolve = res;
						ctx._ssg.reject  = rej;
						vm.runInContext(source, ctx);
					}).then(html => {
						data.html = data.html.replace(`<div id="${options.tagId}">`,`<div id="${options.tagId}">${html}`);
						cb(null, data);
					}).catch(err => (
						cb(new Error(`[webpack] Error during Static Site Generation!\n${err}`), data)
					));
				}
			);

		});
	}
}

module.exports = HtmlSsgPlugin;
