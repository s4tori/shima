const path = require("path");

module.exports = {
	root: path.resolve(__dirname, "./../"),
	app : path.resolve(__dirname, "./../", "src"),
	dist: path.resolve(__dirname, "./../", "dist"),
	css : path.resolve(__dirname, "./../", "src/style"),
	font: path.resolve(__dirname, "./../", "src/style/font"),
	stat: path.resolve(__dirname, "./../", "config/static")
};
