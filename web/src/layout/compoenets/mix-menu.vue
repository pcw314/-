<template>
  <div @mouseleave="closeShow">
    <div class="flex">
      <div
        class="h-screen w-20"
        style="background-color: #1d2c63; color: #949eee"
      >
        <div class="flex h-16 flex-items-center flex-justify-center">
          <!-- <img src="../../assets/image/logo.png" alt="" class="h-16" /> -->
          <i class="iconfont icon-lujing font-size-13"></i>
        </div>
        <div
          @click="findActiveIndex((item as AppRouteRecordRaw).path as string)"
          class="flex flex-col h-16 py-2 flex-items-center flex-justify-center first-menu cursor-pointer"
          :class="{ active: activeIndex === index }"
          v-for="(item, index) in userStore.getRoleRouters"
          :key="index"
        >
          <i
            v-if="(item as AppRouteRecordRaw).icon"
            :class="`iconfont icon-${(item as AppRouteRecordRaw).icon} font-size-5`"
          ></i>
          <i v-else class="iconfont icon-moon font-size-5"></i>

          <div class="font-size-3.5 mt-1">
            {{ (item as AppRouteRecordRaw).title }}
          </div>
        </div>
      </div>
      <div>
        <transition name="el-zoom-in-center">
          <div class="h-screen w-40" v-if="isShow">
            <div
              class="h-16 flex flex-justify-center flex-items-center font-medium"
            >
              Super 管理系统
            </div>
            <mixItem
              class="cursor-pointer"
              v-for="(item, index) in currentRouteItem"
              :key="item.path"
              :item="item"
              :activeRoute="activeRoute"
            ></mixItem>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from "vue";
import { useUserStore } from "@/store/module/user";
import { useRouter } from "vue-router";
import mixItem from "./mix-item.vue";
const { currentRoute } = useRouter();
const userStore = useUserStore();
const activeRoute = ref<any>(currentRoute.value.path);
const activeIndex = ref<number>(0);
const isShow = ref<Boolean>(false);
const currentRouteItem = ref<any[]>([]);
const findActiveIndex = (path) => {
  userStore.getRoleRouters!.forEach((item, index) => {
    if (path.startsWith(item.path)) {
      activeIndex.value = index;
      isShow.value = true;
      currentRouteItem.value = item.children;
    }
  });
};
const closeShow = () => {
  isShow.value = true;
};
watch(
  () => currentRoute.value,
  (route: any) => {
    if (route.path.startsWith("/redirect/")) {
      return;
    }
    activeRoute.value = route.path;
    findActiveIndex(route.path);
  },
  {
    immediate: true,
  }
);
</script>

<style scoped lang="scss">
.first-menu {
  &:hover {
    background-color: var(--el-menu-hover-bg-color);
  }

  &.active {
    color: var(--el-color-primary);
    background-color: var(--el-menu-hover-bg-color);
  }
}
</style>
