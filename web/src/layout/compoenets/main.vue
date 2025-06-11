<template>
  <div >
    <router-view v-slot="{ Component, route }">
     
      <component :is="Component" :key="route.path" />
    </router-view>
  </div>
</template>

<script lang="ts" setup></script>

<style scoped lang="scss">
.slide-right-enter-active,
.slide-right-leave-active {
  will-change: transform;
  transition: all 1s;
}

.slide-right-enter {
  opacity: 0.2;
  transform: scale(0);
}

.slide-right-leave-active {
  opacity: 0;
  transform: scale(1);
}
</style>
