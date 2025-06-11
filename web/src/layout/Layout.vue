<template>
  <el-container class="h-screen flex" @click="menuVisible = false">
    <div
      :style="'min-width:' + configStore.getIsScrollMenue ? '200px;' : '254px;'"
    >
      <el-scrollbar class="shadow h-screen pos-relative">
        <div>
          <!-- <menuLogo></menuLogo> -->
          <mainMenue></mainMenue>
        </div>
        <!-- <mixMenu v-else></mixMenu> -->
      </el-scrollbar>
    </div>

      <el-main class="p-10px flex-col flex">
        <mainHeader />
      <div class="tab p-2 pt-4 pb-0">
        <el-tabs
          v-model="activeRoute"
          type="card"
          class="position-relative"
          @tab-remove="removeTab"
          closable
          @tab-click="clickTab"
        >
          <el-tab-pane
            v-for="(item, index) in userStore.visitedRoutes"
            :key="item.path"
            :name="item.path"
          >
            <template #label>
              <div @contextmenu.prevent="showContextMenu(index, $event)">
                {{ item.meta?item.meta.title:item.title }}
              </div>
            </template>
          </el-tab-pane>
        </el-tabs>
        <div
          v-if="menuVisible"
          :style="{ top: `${menuTop - 60}px`, left: `${menuLeft + 20}px` }"
          class="context-menu"
        >
          <div
            v-for="item in menuItems"
            :key="item.label"
            class="menu-item"
            @click="handleMenuClick(item.callback)"
          >
            {{ item.label }}
          </div>
        </div>
      </div>
          <mainBody class="flex-1 h-full" />
        </el-main>

  </el-container>
</template>

<script lang="ts" setup>
import mixMenu from "./compoenets/mix-menu.vue";
import mainBody from "./compoenets/main.vue";
import mainHeader from "./compoenets/header.vue";
import mainMenue from "./compoenets/menue.vue";
import menuLogo from "./compoenets/menu-logo.vue";
import { useConfigStore } from "@/store/module/config";

import { ref, watch } from "vue";
import { useUserStore } from "@/store/module/user";
import router from "@/router";
import { useRouter } from "vue-router";
import { TabsPaneContext } from "element-plus";
import { onMounted } from "vue";
const boxHeight = ref<string>("0");

const footerData = localStorage.getItem("copyright");
const configStore = useConfigStore();
const isDark = ref<Boolean>(false);
const menuChose = ref<number | null>(null);
const { currentRoute } = useRouter();
const toRouter = useRouter();
const menuVisible = ref(false);
const menuTop = ref(0);
const menuLeft = ref(0);
const goToOverview = () => {
  toRouter.push("/home/overview");
};
const menuItems = [
  // { label: "关闭当前", callback: "current" },
  { label: "关闭所有", callback: "all" },
  { label: "关闭右侧", callback: "right" },
  { label: "关闭左侧", callback: "left", disabled: true }, // 示例禁用项
];
const activeRoute = ref<string>("");
let userStore = useUserStore();
const removeTab = (e: any) => {
  let name = e;
  e = getRouteIndex(e);
  console.log(e, "eeee");
  let routeLength = userStore.visitedRoutes.length;
  if (e < routeLength && e > 0 && routeLength > 1) {
    if (name === activeRoute.value) {
      let toName = userStore.visitedRoutes[e - 1].path;
      router.push({ path: toName });
    }
    userStore.visitedRoutes.splice(e, 1);
  } else if (e === 0 && routeLength > 1) {
    if (name === activeRoute.value) {
      let toName = userStore.visitedRoutes[e + 1].path;
      router.push({ path: toName });
    }
    userStore.visitedRoutes.splice(e, 1);
  } else if (e === 0 && routeLength === 1) {
    router.push({ path: localStorage.getItem("homePage") as string });
    userStore.visitedRoutes.splice(e, 1);
  }
};
const getRouteIndex = (name: string) => {
  let index = -1;
  for (let i = 0; i < userStore.visitedRoutes.length; i++) {
    let item = userStore.visitedRoutes[i];
    if (item.path === name) {
      index = i;
      break;
    }
  }
  return index;
};
const clickTab = (pane: TabsPaneContext, ev: Event) => {
  let routeName = pane.props.name;

  if (routeName === activeRoute.value) {
    return;
  } else {
    router.push({ path: routeName as string });
  }
};
const showContextMenu = (index, event) => {
  menuChose.value = index;
  menuTop.value = event.clientY;
  menuLeft.value = event.clientX;
  menuVisible.value = true;
};
// 处理菜单项点击
const handleMenuClick = (command) => {
  let activeIndex = getRouteIndex(activeRoute.value);
  let routeLength = userStore.visitedRoutes.length;
  let activeName = userStore.visitedRoutes[menuChose.value as number].name;

  switch (command) {
    case "current":
      removeTab(activeName);
      break;
    case "right":
      if (
        (menuChose.value as number) < routeLength - 1 &&
        (menuChose.value as number) > activeIndex
      ) {
        let deleteCount = routeLength - ((menuChose.value as number) + 1);
        userStore.visitedRoutes.splice(
          (menuChose.value as number) + 1,
          deleteCount
        );
      } else if (
        (menuChose.value as number) < routeLength - 1 &&
        (menuChose.value as number) <= activeIndex
      ) {
        let deleteCount = routeLength - ((menuChose.value as number) + 1);
        userStore.visitedRoutes.splice(
          (menuChose.value as number) + 1,
          deleteCount
        );
        router.push({ name: activeName });
      }
      break;
    case "left":
      if (
        routeLength > 1 &&
        (menuChose.value as number) >= activeIndex &&
        (menuChose.value as number) > 0
      ) {
        userStore.visitedRoutes.splice(0, menuChose.value as number);
        router.push({ name: activeName });
      } else if (
        routeLength > 1 &&
        (menuChose.value as number) < activeIndex &&
        (menuChose.value as number) > 0
      ) {
        userStore.visitedRoutes.splice(0, menuChose.value as number);
      }
      break;
    case "all":
      userStore.visitedRoutes.length = 0;
      router.push({ path: "/", replace: true });
  }

  menuVisible.value = false; // 点击后隐藏菜单
};
const openClose = () => {
  userStore.setIsCollopse(!userStore.getIsCollopse);
};
const changeScrollMenue = () => {
  configStore.setIsScrollMenue(!configStore.getIsScrollMenue);
};
watch(
  () => currentRoute.value,
  (route: any) => {
    if (route.path.startsWith("/redirect/")) {
      return;
    }
    activeRoute.value = route.path;
  },
  {
    immediate: true,
  }
);
// const changeTheme = () => {
//   let htmlDom = document.documentElement;
//   if (!isDark.value) {
//     htmlDom.setAttribute("class", "dark");
//   } else {
//     htmlDom.setAttribute("class", "");
//   }
//   isDark.value = !isDark.value;
// };
const loginOut = () => {
  userStore.loginOut();
};
</script>

<style scoped lang="scss">
:deep(.el-menu) {
  border: none;
}

.main {
  height: 100vh;

  &-header {
    min-width: 650px;
  }
}
.menu-item {
  padding: 8px 12px;
  cursor: pointer;
  font-size: 14px;

  &:hover {
    background-color: #f3f3f5;
    color: #409eff;
  }
}

.menu-item.disabled {
  color: #ccc;
  cursor: not-allowed;
}
.context-menu {
  position: absolute;
  background: white;
  border: 1px solid #ccc;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  z-index: 999;
}
</style>
