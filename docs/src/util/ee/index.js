/* eslint-disable no-console */
import _            from "src/util/lodash";
import getEnvConfig from "src/util/config";
import eeColorize   from "./ee-colorize";
import eeActions    from "./ee-actions";


export default {

	teaser() {
		if (!getEnvConfig("teaser")) { return; }
		eeActions(global);

		if (!eeColorize.checkColorSupport()) {
			let bwText = "\n"
				+ "   ▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣\n"
				+ "   ▤▤▤▤▹▹▥▹▹▥▹▹▨◃◃▧◃◃▧◃◃◃▦▦◃\n"
				+ "   ▤▹▹▹▹▹▥▹▹▥▹▹▨◃◃◃▧▧◃◃◃▦◃◃▦\n"
				+ "   ▹▤▤▹▹▹▹▥▥▹▹▹▨◃◃▧◃◃▧◃◃▦▦▦▦\n"
				+ "   ▹▹▹▤▹▹▥▹▹▥▹▹▨◃◃▧◃◃▧◃◃▦◃◃▦\n"
				+ "   ▤▤▤▤▹▹▥▹▹▥▹▹▨◃◃▧◃◃▧◃◃▦◃◃▦\n"
				+ "   ▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣\n\n"
				+ "  💽 source-code: shima.sourceCode(); ↵\n"
				+ "  ⌨️ author:      shima.author(); ↵\n\n"
			;

			// Firefox Quantom doesn't like funny console.log output anymore 😭
			if (eeColorize.isFirefoxQuatom()) {
				bwText = bwText.replace(/▹|◃/g, "▢");
			}

			// Safari doesn't use a monospace font in the console 😭
			if (eeColorize.isSafari()) {
				bwText = bwText.replace(/▹|◃/g, "◻");
			}

			console && console.log(bwText);
			return;
		}

		// We need to use ◻ (not ▢ !) for Safari 😅
		const shapes = ["▤", "▥", "▨", "▧", "▦", "▩"]; //, "◧", "◨", "◩", "◪"];
		let ee       = ("\n"
			+ "◻◻▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣\n"
			+ "◻◻▤▤▤▤◻▥◻◻▥◻▨◻▧◻◻▧◻▹▦▦▦\n"
			+ "◻◻▤◻◻◻◻▥◻◻▥◻▨◻▹▧▧▧▹◻▦◻◻▦\n"
			+ "◻◻▹▤▤▤◃▥▥▥▹◻▨◻▧▹▧▹▧◻▦▦▦▦\n"
			+ "◻◻◻◻◻▤◻▥◻◻▥◻▨◻▧◻◻▧◻▦◻◻▦\n"
			+ "◻◻▤▤▤▤◻▥◻◻▥◻▨◻▧◻◻▧◻▦◻◻▦\n"
			+ "◻◻▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣▣\n"
		);

		let aColor    = [];
		const fs      = 20;
		const randClr = _.values(eeColorize.colors);
		const clrs    = {
			"▤": ["#333333", "#65737E"], "▥": ["#000000", "#65737E"],
			"▨": ["#5C5C5C", "#65737E"], "▧": ["#000000", "#65737E"],
			"▦": ["#333333", "#65737E"], "◻": randClr[_.random(0, randClr.length - 1)]
		};

		ee = ee.replace(/\S/g, text => {
			if (_.includes(["▹", "◻", "◃"], text)) {
				aColor.push(`color: transparent; font-size: ${text === "▹" ? fs / 2 : fs}px; line-height: 0;`);
				return `%c${text === "◃" ? "◻◻" : "◻"}`;
			}

			const clr = clrs[text] || clrs["◻"];
			aColor.push(`color: ${typeof clr === "string" ? clr : clr[_.random(0, clr.length - 1)]}; font-size: ${fs}px; line-height: ${fs}px;`);
			return `%c${text === "▣" ? "▣" : shapes[_.random(0, shapes.length - 1)]}`;
		});

		eeColorize.colorize(""
			+ ee + "\n"
			+ "[def]\n💻 author:      [blue]shima[gray].[purple]author[gray]()[black]; [gray]↵"
			+ "[def]\n💽 source-code: [blue]shima[gray].[purple]sourceCode[gray]()[black]; [gray]↵\n\n"
		, aColor);
	}

};
