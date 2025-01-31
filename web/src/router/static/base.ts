import type { RouteRecordRaw } from 'vue-router'

/**
 * 后台基础路由路径
 * 您可以随时于后台->系统配置中修改此值，程序可自动完成代码修改，同时建立对应的API入口和禁止admin应用访问
 */
export const baseRoutePath = '/'

/**
 * 后台基础静态路由配置
 * 这里定义了整个应用的基础路由结构
 */
const baseRoute: RouteRecordRaw = {
    // 路由路径，使用上面定义的baseRoutePath
    path: baseRoutePath,
    // 路由名称，用于编程式导航
    name: 'main',
    // 异步加载主布局组件
    component: () => import('/@/layouts/index.vue'),
    // 初始重定向到loading路由
    redirect: '/loading',
    // 路由元信息
    meta: {
        title: `pagesTitle.main`, // 页面标题，支持国际化
    },
    // 子路由配置
    children: [
        {
            // 加载页路由，:to? 是可选参数，用于指定加载完成后的跳转目标
            path: 'loading/:to?',
            // 路由名称
            name: 'mainLoading',
            // 异步加载加载页组件
            component: () => import('/@/layouts/common/components/loading.vue'),
            // 路由元信息
            meta: {
                title: `pagesTitle.loading`, // 加载页标题，支持国际化
            },
        },
    ],
}

export default baseRoute