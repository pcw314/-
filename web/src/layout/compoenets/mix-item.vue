<template>
  <div
    class="h-13 pl-2 flex flex-items-center text-size no-child"
    :class="{ active: item.path === activeRoute }"
    v-if="
      (!item.children ||
        item.children.length === 0 ||
        item.children[0].type === 3) &&
      !item.is_hide
    "
    @click="skipPatn(item.path)"
  >
    {{ item.title }}
  </div>
  <div
    v-else-if="
      !item.is_hide && item.children.length && item.children[0].type !== 3
    "
  >
    <el-collapse accordion v-model="defaultExpandCollopse">
      <el-collapse-item :name="item.path">
        <template #title>
          <span
            class="pl-2"
            :class="{ 'active-text': activeRoute.startsWith(item.path) }"
          >
            {{ item.title }}</span
          >
        </template>
        <div class="">
          <mix-item
            v-for="child in item.children"
            :key="item.id"
            :item="child"
            :activeRoute="activeRoute"
          ></mix-item>
        </div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script lang="ts" setup>
import { watch, ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
const defaultExpandCollopse = ref<string>("");
const { activeRoute, item } = defineProps({
  item: {
    type: Object,
    default: () => {},
  },
  activeRoute: {
    type: String,
    default: "",
  },
});
const skipPatn = (path: string) => {
  router.push({ path });
};

const expandDefaultExpandCollopse = (path: string) => {
  defaultExpandCollopse.value = path;
  console.log(path, "path");
};
watch(
  () => {
    return activeRoute;
  },
  (activeRoute) => {
    if (
      activeRoute.startsWith(item.path) &&
      item.children &&
      item.children.length &&
      !item.is_hide
    ) {
      expandDefaultExpandCollopse(item.path);
    }
  },
  { immediate: true }
);
</script>

<style scoped lang="scss">
:deep(.el-collapse-item__header) {
  font-size: 0.875rem;
  color: #000000;
  border: none;

  &:hover {
    background-color: var(--el-menu-hover-bg-color);
  }
}

:deep(.el-collapse) {
  border: 0 !important;
  font-size: 14px;
}

:deep(.el-collapse-item__content) {
  padding-bottom: 0;
}

.text-size {
  font-size: 14px;
}

.active-text {
  color: var(--el-color-primary);
}

.no-child {
  &:hover {
    background-color: var(--el-menu-hover-bg-color);
  }
}

:deep(.el-collapse-item__wrap) {
  border: none !important;
}

.active {
  color: var(--el-color-primary);
  background-color: var(--el-menu-hover-bg-color);

  &:hover {
    background-color: var(--el-menu-hover-bg-color);
  }
}
</style>
