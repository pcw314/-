import createAxios from '/@/utils/axios'

export const url = '/soft'

export function getSoftList(params: anyObj) {
    return createAxios({
        url: url,
        method: 'get',
        params: params,
    })
}
