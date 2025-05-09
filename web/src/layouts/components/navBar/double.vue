<template>
    <div class="layouts-menu-horizontal-double">
        <el-scrollbar ref="layoutMenuScrollbarRef" class="double-menus-scrollbar">
            <el-menu ref="layoutMenuRef" class="menu-horizontal" mode="horizontal" :default-active="state.defaultActive">
                <MenuTree :extends="{ position: 'horizontal', level: 1 }" :menus="navTabs.state.tabsViewRoutes" />
            </el-menu>
        </el-scrollbar>
        <NavMenus />
    </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, reactive } from 'vue'
import { useRoute, onBeforeRouteUpdate, type RouteLocationNormalizedLoaded } from 'vue-router'
import { currentRouteTopActivity } from '/@/layouts/components/menus/helper'
import MenuTree from '/@/layouts/components/menus/menuTree.vue'
import { layoutMenuRef, layoutMenuScrollbarRef } from '/@/stores/refs'
import NavMenus from '/@/layouts/components/navMenus.vue'
import { useNavTabs } from '/@/stores/navTabs'
import { useConfig } from '/@/stores/config'

const config = useConfig()
const navTabs = useNavTabs()
const route = useRoute()

const state = reactive({
    defaultActive: '',
})

// 激活当前路由的菜单
const currentRouteActive = (currentRoute: RouteLocationNormalizedLoaded) => {
    let routeChildren = currentRouteTopActivity(currentRoute.path, navTabs.state.tabsViewRoutes)
    if (routeChildren) state.defaultActive = currentRoute.path
}

// 滚动条滚动到激活菜单所在位置
const verticalMenusScroll = () => {
    nextTick(() => {
        let activeMenu: HTMLElement | null = document.querySelector('.el-menu.menu-horizontal li.is-active')
        if (!activeMenu) return false
        layoutMenuScrollbarRef.value?.setScrollTop(activeMenu.offsetTop)
    })
}

onMounted(() => {
    currentRouteActive(route)
    verticalMenusScroll()
})

onBeforeRouteUpdate((to) => {
    currentRouteActive(to)
})
</script>

<style scoped lang="scss">
.layouts-menu-horizontal-double {
    display: flex;
    align-items: center;
    height: var(--el-header-height);
    background-color: var(--ba-bg-color-overlay);
    border-bottom: 1px solid var(--el-color-info-light-8);
}
.double-menus-scrollbar {
    width: 70vw;
    height: var(--el-header-height);
}
.menu-horizontal {
    border: none;
    --el-menu-bg-color: v-bind('config.getColorVal("menuBackground")');
    --el-menu-text-color: v-bind('config.getColorVal("menuColor")');
    --el-menu-active-color: v-bind('config.getColorVal("menuActiveColor")');
}

.el-sub-menu .icon,
.el-menu-item .icon {
    vertical-align: middle;
    margin-right: 5px;
    width: 24px;
    text-align: center;
    flex-shrink: 0;
}
.is-active .icon {
    color: var(--el-menu-active-color) !important;
}
.el-menu-item.is-active {
    background-color: v-bind('config.getColorVal("menuActiveBackground")');
}
</style>
