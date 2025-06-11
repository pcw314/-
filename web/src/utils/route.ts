import router from "@/router/index";

export const Layout = () =>
  import("../layout/Layout.vue").then((res) => res.default);
export const Login = () =>
  import("@/views/Login/index.vue").then((res) => res.default);
export function findParents(routes: AppRouteRecordRaw[], targetPath: string) {
  let parents: AppRouteRecordRaw[] = [];
  let fullPath: string = "";
  function traverse(routes: AppRouteRecordRaw[], currentPath = "") {
    for (let route of routes) {
      fullPath = route.path as string;

      if (targetPath.startsWith(fullPath)) {
        parents.push(route);
        if (!route.children) {
          break;
        }
        traverse(route.children, fullPath);
      }
    }
  }
  traverse(routes);
  return parents;
}
 
let modules = import.meta.glob('@/views/**/*.vue')
// 过滤路由的方法

export function filterAsyncRouter(asyncRouterMap) {

  asyncRouterMap.forEach((routeItem) => {
    if (
      !routeItem.redirect &&
      (!routeItem.children ||
        (!routeItem.children.length ) || routeItem.children[0].type===3 ) 
    ) {
      let componentPath = `/src/views${routeItem.path}/index.vue`
      let componentFun = modules[componentPath]
       let route 
      if(componentFun){
      route =   {
        ...routeItem,
        component: componentFun,
        meta: { title: routeItem.title },
      };
      }else {
       route =   {
        ...routeItem,
        component: Login,
        meta: { title: routeItem.title },
      };
      }
      if(routeItem.children && routeItem.children.length){
         filterAsyncRouter(routeItem.children)
      }
     
      router.addRoute("root", route);
    } else if (routeItem.children && routeItem.children.length) {
      filterAsyncRouter(routeItem.children);
    }
  });
}
export const initRouter = async () => {
  let userInfo = JSON.parse(localStorage.getItem("user") || "{}");
  if (userInfo && userInfo.roleRouters && userInfo.roleRouters.length>0) {
   let homePage = localStorage.getItem('homePage')
  console.log(homePage)
    router.addRoute({
      path:"/",
      name: 'root',
      redirect:homePage,
      component:Layout,
      meta: {
        title: "root",
        constant: true,
      },
      children:[]
    })
    filterAsyncRouter(userInfo.roleRouters);
    router.addRoute({
      path: "/:pathMatch(.*)*",
      component: () => import("@/views/error/index.vue"),
      name: "NoFind",
      redirect: "/404",
      children: [],
      meta: {
        hidden: true,
        title: "404",
        noTagsView: true,
      },
    });
  }
};

export function getHomePage(routeItem){
if(routeItem.children && routeItem
  .children.length > 0 ){
 return  getHomePage(routeItem.children[0])
}else{
  return routeItem.path
}

}