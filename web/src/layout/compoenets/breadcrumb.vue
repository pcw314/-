<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item v-for="item in breadcrumbLsit" :key="item.id">
      <el-dropdown @command="handleCommand">
        <span class="f16" style="outline: none;color: #333;">
          {{ item.title }}
        </span>
        <template #dropdown v-if="item.children && item.children.length">
          <el-dropdown-menu>
            <el-dropdown-item
              v-for="child in item.children"
              :key="child.id"
              :command="child"
              >{{ child.title }}</el-dropdown-item
            >
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/module/user";
import { findParents } from "@/utils/route";
import { ref, watch } from "vue";
import router from "@/router";
const { currentRoute } = useRouter();
const userStore = useUserStore();
const breadcrumbLsit = ref<AppRouteRecordRaw[]>([]);
const handleCommand = (command: AppRouteRecordRaw) => {
  if (!command.children) {
    router.push({ path: command.path as string });
  } else {
    findLowestChildRoutte(command.children);
  }
};
function findLowestChildRoutte(list) {
  let route = list[0];
  if (!route.children) {
    router.push({ path: route.path });
  } else {
    findLowestChildRoutte(route.children);
  }
}

const getBreadcrumb = () => {
  const currentPath = currentRoute.value.matched.slice(-1)[0].path;
  let result = findParents(userStore.getRoleRouters as AppRouteRecordRaw[], currentPath);
  breadcrumbLsit.value = result;

};
watch(
  () => currentRoute.value,
  (route: any) => {
    if (route.path.startsWith("/redirect/")) {
      return;
    }
    getBreadcrumb();
  },
  {
    immediate: true,
  }
);
</script>

<style scoped lang="scss">
:deep(.el-dropdown) {
  outline: none !important;
}
</style>
