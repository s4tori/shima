/* global module */
import Vue    from "vue";
import pCb    from "./plugin/clipboard";
import App    from "./App.vue";
import { ee } from "./util";


ee.teaser();
Vue.use(pCb);
new Vue({ ...App, el: "#app" });

if (module.hot) {
	module.hot.accept("./util", () => ee.teaser());
}
