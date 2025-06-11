<template>
  <el-sub-menu :index="item.path" v-if="
    item.children &&
    item.children.length > 0 &&
    !item.is_hide &&
    item.children[0].type !== 3
  ">
    <template #title>
      <span>{{ item.title }} </span></template>
    <menu-item :key="child.id" :item="child" v-for="child in item.children" />
  </el-sub-menu>
  <el-menu-item v-else-if="
    !item.is_hide &&
    (!item.children ||
      item.children.length === 0 ||
      item.children[0].type === 3)
  " @click="goToPath(item.path, item)" :index="item.path">{{ item.title }}</el-menu-item>
</template>

<script lang="ts" setup>
import { useUserStore } from "@/store/module/user";
import { useRouter } from "vue-router";
let userStore = useUserStore();

let router = useRouter();
const props = defineProps({
  item: { type: Object, required: true },
});
const goToPath = (item, routeItem) => {
  console.log(routeItem)
  if (!userStore.visitedRoutes.some((route) => {
    return route.path === item
  })) {
    userStore.visitedRoutes.push(routeItem)

  }
  router.push({ path: item });
};
</script>

<style scoped></style>
