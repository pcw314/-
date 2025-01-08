import { baseRoutePath } from '/@/router/static/base'

/*
 * 语言包按需加载映射表
 * 使用固定字符串 ${lang} 指代当前语言
 * key 为页面 path，value 为语言包文件相对路径，访问时，按需自动加载映射表的语言包，同时加载 path 对应的语言包（若存在）
 */
export default {
    '/': ['./${lang}/index.ts'],
    [baseRoutePath + '/moduleStore']: ['./${lang}/module.ts'],
    [baseRoutePath + '/user/rule']: ['./${lang}/auth/rule.ts'],
    [baseRoutePath + '/user/scoreLog']: ['./${lang}/user/moneyLog.ts'],
    [baseRoutePath + '/crud/crud']: ['./${lang}/crud/log.ts', './${lang}/crud/state.ts'],
}
