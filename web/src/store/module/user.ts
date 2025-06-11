import router from '@/router'
import { defineStore } from 'pinia'
import { filterAsyncRouter,getHomePage } from '@/utils/route'
// import { menue } from '@/utils/test'
import Layout from "@/layout/Layout.vue";
interface UserState {
  userInfo?: any,
  token: string
  roleRouters?: string[] | AppRouteRecordRaw[],
  isAddDynamicRouter:Boolean,
  filterRoutes:string[] | AppRouteRecordRaw[],
  visitedRoutes:any[]|AppRouteRecordRaw[],
  isCollopse:Boolean
}

export const useUserStore = defineStore("user", {
  state: ():UserState => {
    return {
      token: '',
      userInfo: undefined,
      roleRouters:[],
      isAddDynamicRouter:false,
      filterRoutes: [],
      visitedRoutes:[],
      isCollopse:false
     
    }
  },
  getters: {
    getIsCollopse():Boolean{
      return this.isCollopse
    },
    getFilterRoutes(): string[] | AppRouteRecordRaw[] {
      return this.filterRoutes
    },
    getToken(): string {
      return this.token
    },
    getUserInfo(): Object | undefined {
      return this.userInfo
    },
    getRoleRouters(): string[] | AppRouteRecordRaw[] | undefined {
      return this.roleRouters
    },
    getIsAddDynamicRouter(): Boolean {
      return  this.isAddDynamicRouter
    },
    
  },
  actions: {
    setIsCollopse(data:boolean){
      this.isCollopse = data
     },
      setVisitedRoutes(list:any[]){
        this.visitedRoutes = []
      },
      setFilterRoutes(data: string[] | AppRouteRecordRaw[]) {
         this.filterRoutes = data
      },
      setToken(tokenStr:string){
        this.token = tokenStr
      },
      setUserInfo(userInfo:Object) {
        this.userInfo = userInfo
      },
      setRoleRouters(roleRouters: string[] | AppRouteRecordRaw[]) {
        this.roleRouters = roleRouters
        let homePage = getHomePage(roleRouters[0])
        localStorage.setItem('homePage',homePage)
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
        filterAsyncRouter(roleRouters)
        console.log(router.getRoutes())
        router.push({path:"/"})
      },
      setIsAddDynamicRouter(isAdd:boolean) {
        this.isAddDynamicRouter = isAdd
      },
    loginOut(){
      this.setFilterRoutes([])
      this.setToken("")
      this.setUserInfo("")
      localStorage.setItem('activeTab','')
      localStorage.setItem('taskDetail','')
      // this.setRoleRouters([])
      this.roleRouters= []
      this.setIsAddDynamicRouter(false)
      this.setVisitedRoutes([])
      router.replace({path:"/login"})
    }
  },
  // pinia数据持久化
  persist: true
})

  