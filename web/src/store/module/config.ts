import { defineStore } from 'pinia'
interface UserState {
 isScrollMenue:Boolean
}

export const useConfigStore = defineStore("config", {
  state: ():UserState => {
    return {
        isScrollMenue:true
     
    }
  },
  getters: {
    getIsScrollMenue():Boolean{
      return this.isScrollMenue
    },
 
    
  },
  actions: {
   setIsScrollMenue(data:boolean){
    this.isScrollMenue = data
   },
   
  },
  // pinia数据持久化
  persist: true
})

  