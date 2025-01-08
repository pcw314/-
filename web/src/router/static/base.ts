import type { RouteRecordRaw } from 'vue-router'

/**
 * 后台基础路由路径
 * 您可以随时于后台->系统配置中修改此值，程序可自动完成代码修改，同时建立对应的API入口和禁止admin应用访问
 */
export const baseRoutePath = '/'

/*
 * 后台基础静态路由
 */
const baseRoute: RouteRecordRaw = {
    path: baseRoutePath,
    name: 'main',
    component: () => import('/@/layouts/index.vue'),
    // 直接重定向到 loading 路由
    redirect: '/loading',
    meta: {
        title: `pagesTitle.main`,
    },
    children: [
        {
            path: 'loading/:to?',
            name: 'mainLoading',
            component: () => import('/@/layouts/common/components/loading.vue'),
            meta: {
                title: `pagesTitle.loading`,
            },
        },
    ],
}

export default baseRoute
