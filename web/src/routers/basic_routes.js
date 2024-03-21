// 基础路由
export default [
    {
        path: '/',
        component: () => import('@/views/home/index.vue'),
        meta: {
            title: '前置页'
        },
        hidden: true,
    },
]