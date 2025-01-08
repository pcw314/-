import createAxios from '/@/utils/axios'

export const url = '/admin/'

export function login(data: anyObj) {
    return createAxios({
        url: url + 'login',
        data: data,
        method: 'POST',
    })
}

export function logout() {
    return createAxios({
        url: url + 'logout',
        method: 'POST',
    })
}

export function getMenu() {
    return createAxios({
        url: url + 'menu',
        method: 'get',
    })
}
