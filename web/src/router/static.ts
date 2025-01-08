import type { RouteRecordRaw } from 'vue-router'
import { baseRoutePath } from './static/base'

const pageTitle = (name: string): string => {
    return `pagesTitle.${name}`
}

/*
 * 静态路由
 * 自动加载 ./static 目录的所有文件，并 push 到以下数组
 */
const staticRoutes: Array<RouteRecordRaw> = [
    {
        // 管理员登录页 - 不放在 baseRoute.children 因为登录页不需要使用后台的布局
        path: baseRoutePath + 'login',
        name: 'login',
        component: () => import('/@/views/login.vue'),
        meta: {
            title: pageTitle('login'),
        },
    },
    // {
    //     path: '/:path(.*)*',
    //     redirect: '/404',
    // },
    {
        // 404
        path: '/404',
        name: 'notFound',
        component: () => import('/@/views/common/error/404.vue'),
        meta: {
            title: pageTitle('notFound'), // 页面不存在
        },
    },
    {
        // 后台找不到页面了-可能是路由未加载上
        path: baseRoutePath + ':path(.*)*',
        redirect: (to) => {
            return {
                name: 'mainLoading',
                params: {
                    to: JSON.stringify({
                        path: to.path,
                        query: to.query,
                    }),
                },
            }
        },
    },
    {
        // 无权限访问
        path: '/401',
        name: 'noPower',
        component: () => import('/@/views/common/error/401.vue'),
        meta: {
            title: pageTitle('noPower'),
        },
    },
]

const staticFiles: Record<string, Record<string, RouteRecordRaw>> = import.meta.glob('./static/*.ts', { eager: true })
for (const key in staticFiles) {
    if (staticFiles[key].default) staticRoutes.push(staticFiles[key].default)
}

export default staticRoutes
