const getConfig = require("../../config");


module.exports = ({ env }) => {
	const prodEnv = (env === "production");
	const webConf = getConfig(env);

	return {
		__PRODUCTION__:  prodEnv,
		__DEVELOPMENT__: !prodEnv,
		__CONFIG__: JSON.stringify(webConf.client),
		__VUE_OPTIONS_API__: true,
		__VUE_PROD_DEVTOOLS__: false,
		"process.env"  : {
			NODE_ENV   : JSON.stringify(env),
			VUE_ENV    : "'client'"
		}
	};
};
