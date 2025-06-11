import router from "./router/index";
import { NProgressStart, NProgressDone } from "@/utils/nprogress";
import { useUserStore } from "@/store/module/user";
import { filterAsyncRouter } from "@/utils/route";
import { NO_REDIRECT_WHITE_LIST } from "@/constants";
import { useTitle } from "@/hoosk/web/useTitle";
import { ElLoading } from "element-plus";
let loading;
router.beforeEach(async (to, from, next) => {
  NProgressStart();
  loading = ElLoading.service({
    lock: true,
    text: "Loading",
    background: "rgba(0, 0, 0, 0.7)",
  });
  const userStore = useUserStore();
  if (userStore.getToken) {
    if (to.path === "/login") {
      if (userStore.userInfo?.role != 3) {
        next("/center/index");
      }
    } else {
      if (userStore.userInfo?.role !== 3) {
        next();
        return;
      }
      if (userStore.getIsAddDynamicRouter) {
        next();
        return;
      }

      router.addRoute({
        path: "/:pathMatch(.*)*",
        redirect: "/404",
        component: () => import("@/views/error/index.vue"),
        name: "NoFind",
        children: [],
        meta: {
          hidden: true,
          title: "404",
          noTagsView: true,
        },
      });

      const redirectPath = from.query.redirect || to.path;
      const redirect = decodeURIComponent(redirectPath as string);
      const nextData =
        to.path === redirect ? { ...to, replace: true } : { path: redirect };
      userStore.setIsAddDynamicRouter(true);

      next(nextData);
    }
  } else {
    if (NO_REDIRECT_WHITE_LIST.indexOf(to.path) !== -1) {
      next();
    } else {
      next(`/login?redirect=${to.path}`); // 否则全部重定向到登录页
    }
  }
});
router.afterEach((to) => {
  const userStore = useUserStore();
  let result = userStore.visitedRoutes.find((val) => val.name == to.name);
  if (!result && NO_REDIRECT_WHITE_LIST.indexOf(to.path) == -1) {
    console.log(to.path);
    userStore.visitedRoutes.push(to as AppRouteRecordRaw);
  }
  loading.close();
  useTitle(to?.meta?.title as string);
  NProgressDone();
});
