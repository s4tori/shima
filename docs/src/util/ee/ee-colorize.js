/* eslint-disable no-console */
import _ from "src/util/lodash";


export default {

	browser: [
		{ id: "chrome" , "userAgent": /Chrome/        , "vendor": /Google Inc/    , color: true  },
		{ id: "safari" , "userAgent": /Safari/        , "vendor": /Apple Computer/, color: true  },
		{ id: "opera"  , "userAgent": /OPR/           , "vendor": /Opera/         , color: true  },
		{ id: "quantom", "userAgent": /Firefox\/57/   , "vendor": null            , color: false },
		{ id: "firefox", "userAgent": /Firefox/       , "vendor": null            , color: true  },
		{ id: "ie"     , "userAgent": /(MSIE|Trident)/, "vendor": null            , color: false }
	],

	colors: {
		red   : "#EC5f67",
		orange: "#F99157",
		yellow: "#FAC863",
		green : "#99C794",
		cyan  : "#5FB3B3",
		purple: "#C594C5",
		blue  : "#6699CC",
		brown : "#AB7967",
		gray  : "#999999",
		black : "#000000",
		def   : "inherit"
	},

	getBrowser() {
		const currentBrowser = _.find(this.browser, (b) => {
			return b.vendor ?
				b.userAgent.test(navigator.userAgent) && b.vendor.test(navigator.vendor) :
				b.userAgent.test(navigator.userAgent)
			;
		});

		this.getBrowser = function () {
			return currentBrowser || { color: false };
		};

		return this.getBrowser();
	},

	checkColorSupport() {
		return this.getBrowser().color;
	},

	isFirefoxQuatom() {
		return this.getBrowser().id === "quantom";
	},

	isSafari() {
		return this.getBrowser().id === "safari";
	},

	isChrome() {
		return this.getBrowser().id === "chrome";
	},

	colorize(text, clrs = []) {
		if (!this.checkColorSupport()) {
			console && console.log(text.replace(/\[[^\]]*\]|%c/g, ""));
			return;
		}

		// On Chrome, we can't break lines when we use colors (%c) with console.log
		// We need to split each line. 😒
		let iColor = 0;
		this.isChrome()
			? _.forEach(text.split("\n"), line => {
				const iLineColor = (line.match(/%c/g) || []).length;
				iColor = iColor + iLineColor;
				this._log(line, _.slice(clrs, iColor - iLineColor, iColor));
			})
			: this._log(text, clrs)
		;
	},

	_log(text, clrs = []) {
		let aColor = [];
		const cText = text.replace(new RegExp("\\[[\\w]+\\]", "g"), (color) => {
			aColor.push("color: " + this.colors[color.substring(1, color.length - 1)]);
			return "%c";
		});

		aColor = clrs.concat(aColor);
		aColor.unshift(cText);
		console && console.log.apply(console, aColor);
	}

};
