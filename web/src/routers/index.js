import {createRouter, createWebHistory} from "vue-router";
import routes from "./basic_routes"
import NProgress from '../utils/nprogress';

//创建路由
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
});



//全局前置守卫
router.beforeEach(async (to, from) => {
    NProgress.start();
    return true
})

//全局后置钩子
router.afterEach(() => {
    NProgress.done();

});

router.onError(error => {
    NProgress.done();
    console.warn("路由错误", error.message);
});

// router.onBeforeRouteEnter(to, form, next)
// {
//     console.log("跳转")
//     console.log(to)
//     console.log(form)
//     console.log(next)
// }
export default router;