const getConfig = require("../../config");


module.exports = ({ env }) => {
	const prodEnv = (env === "production");
	const webConf = getConfig(env);

	return {
		__PRODUCTION__ :  prodEnv,
		__DEVELOPMENT__: !prodEnv,
		__CONFIG__     : JSON.stringify(webConf.client),
		"process.env"  : {
			NODE_ENV   : JSON.stringify(env),
			VUE_ENV    : "'client'"
		}
	};
};
