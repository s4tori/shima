const _          = require("lodash");
const configDev  = require("./config-dev");
const configProd = require("./config-prod");


module.exports = function getConfig(env) {
	return _.merge({ }, configProd, env === "production" ? { } : configDev);
};
