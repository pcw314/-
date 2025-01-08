import { defineStore } from 'pinia'
import { ADMIN_INFO } from '/@/stores/constant/cacheKey'
import type { AdminInfo } from '/@/stores/interface'

export const useAdminInfo = defineStore('adminInfo', {
    state: (): AdminInfo => {
        return {
            type: 2,
            username: '',
            token: '',
        }
    },
    actions: {
        dataFill(state: AdminInfo) {
            this.$state = { ...this.$state, ...state }
        },
        removeToken() {
            this.token = ''
        },
        setToken(token: string) {
            this.token = token
        },
        getToken() {
            return this.token
        },
    },
    persist: {
        key: ADMIN_INFO,
    },
})
