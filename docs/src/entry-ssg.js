import createApp          from "./app";
import { renderToString } from "vue/server-renderer";


global.__SSR__ = true;

renderToString(createApp())
	.then(global._ssg.resolve)
	.catch(global._ssg.reject)
;
