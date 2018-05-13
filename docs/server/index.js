/* eslint-disable no-console */
const Koa     = require("koa");
const conf    = require("../config/config")(process.env.NODE_ENV);
const devUtil = require("./index.dev");


const app = new Koa();

devUtil(app).then((getHtml) => {

	app.use(async ctx => {
		ctx.body = getHtml();
	});

	const { port, host } = conf.server;
	app.listen(port, host,  () => {
		console.log("  [########################################################]");
		console.log("  ♥                   (^.^)/ shima \\(^.^)                  ♥");
		console.log(" ");
		console.log(" ", `environment: '${process.env.NODE_ENV}'`);
		console.log(" ", `dirname    : ${__dirname}`);
		console.log(" ", `server     : ${host}:${port}`);
		console.log(" ");
		console.log(" ", `☕️ Server started!`);
		console.log(" ", "🚀 Waits for incoming requests...");
		console.log(" ");

	});

});
