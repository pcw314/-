import { createRouter, createWebHashHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import type { App } from "vue";
import Layout from "@/layout/Layout.vue";

// 静态路由
export const constantRouterMap: AppRouteRecordRaw[] = [
  {
    name: "root",
    path: "/",
    redirect: "/home/data",
    component: Layout,
    meta: {
      title: "root",
      constant: true,
    },
  },
  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
    name: "Login",
    meta: {
      hidden: true,
      title: "登录",
      noTagsView: true,
    },
  },
  {
    path: "/userInit",
    component: () => import("@/views/userInit/index.vue"),
    name: "userInit",
    meta: {
      hidden: true,
      title: "222",
      noTagsView: true,
    },
  },
  {
    path: "/center/index",
    component: () => import("@/views/center/index.vue"),
    name: "center",
    meta: {
      hidden: true,
      title: "初始信息",
      noTagsView: true,
    },
  },
  {
    component: () => import("@/views/error/index.vue"),
    name: "404",
    path: "/404",
    meta: {
      hidden: true,
      title: "404",
      noTagsView: true,
    },
  },
];
const router = createRouter({
  history: createWebHashHistory(),
  strict: true,
  routes: constantRouterMap as RouteRecordRaw[],
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

export const resetRouter = (): void => {
  const resetWhiteNameList = ["Redirect", "Login", "NoFind", "Root"];
  router.getRoutes().forEach((route) => {
    const { name } = route;
    if (name && !resetWhiteNameList.includes(name as string)) {
      router.hasRoute(name) && router.removeRoute(name);
    }
  });
};

export const setupRouter = (app: App<Element>) => {
  app.use(router);
};

export default router;
