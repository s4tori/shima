import _ from "src/util/lodash";


export default function getEnvConfig(key) {
	return key ? _.get(__CONFIG__, key) : __CONFIG__;
}
