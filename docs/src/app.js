import { createSSRApp } from "vue";
import pCb              from "./plugin/clipboard";
import App              from "./App.vue";


export default () => createSSRApp(App).use(pCb);
