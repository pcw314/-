import createAxios from '/@/utils/axios'

export const url = '/version'

export function getVersionList(params: anyObj) {
    return createAxios({
        url: url,
        method: 'get',
        params: params,
    })
}
