import getEnvConfig from "src/util/config";
import eeColorize   from "./ee-colorize";


export default (g) => {

	g.shima = {

		sourceCode() {
			window.location.href = getEnvConfig("urlSource");
			return "🚀 let's go!";
		},

		author() {
			eeColorize.colorize("lovingly [red]♥[black] brought to you by:");
			return "547021 [¬_¬]~";
		}
	};

};
