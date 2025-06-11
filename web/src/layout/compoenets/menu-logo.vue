<template>
  <div
    class="flex bg-white h-14 flex-items-center flex-justify-center justify-center pos-absolute pos-top-none pos-right-none pos-left-none pos-top-none pos-z-99"
  >
    <i class="iconfont icon-mianbaoxie" style="color: #333"></i>
    <text v-if="!userStore.getIsCollopse" class="ml-4">NILSPACE</text>
  </div>
</template>

<script lang="ts" setup>
import {} from "vue";
import { useUserStore } from "@/store/module/user";
const userStore = useUserStore();
</script>

<style scoped>
.bg-white {
  background-color: var(--el-menu-bg-color) !important ;
}
</style>
