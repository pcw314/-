<template>
  <el-menu
    :default-active="activeRoute"
    :unique-opened="false"
    class="el-menu-vertical "
    :collapse="userStore.getIsCollopse"
    :collapse-transition="true"
  >
    <menu-item
      v-for="route in routes as AppRouteRecordRaw[]"
      :key="route.id"
      :item="route"
    ></menu-item>
  </el-menu>
</template>

<script lang="ts" setup>
import { ref, watch, computed } from "vue";
import menuItem from "./menu-item.vue";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/module/user";

const { currentRoute } = useRouter();
const userStore = useUserStore();
const activeRoute = ref<any>(currentRoute.value.path);
const routes = computed(() => {
  return userStore.getRoleRouters;
});
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
</script>
<style scoped lang="scss">
.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
}
</style>
