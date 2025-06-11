import request from '@/axios'
import type { loginData } from './types'

export const loginAdminApi = (data:loginData ): Promise<any> => {
    return request.post({ url: '/admin/login', data })
  }
export const menuApi = ():Promise<any>=>{
    return request.get({ url: '/admin/menu/list',  })
}
export const userRegister = (data:loginData ): Promise<any> => {
  return request.post({ url: '/user/register', data })
}
