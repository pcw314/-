import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "virtual:uno.css";
import "./permission";
import { setupStore } from "@/store/idnex";
import { initRouter } from "@/utils/route";
import "element-plus/theme-chalk/dark/css-vars.css";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
// 路由
import { setupRouter } from "./router";



const setupAll = async () => {
  const app = createApp(App);
  for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
  }
  await initRouter();
  setTimeout(() => {
    setupStore(app);
    setupRouter(app);
    app.use(ElementPlus);
    app.mount("#app");
  }, 100);
};
setupAll();
