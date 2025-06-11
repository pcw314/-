/**
 * 请求成功状态码
 */
export const SUCCESS_CODE = 200;

/**
 * 请求contentType
 */
export const CONTENT_TYPE: AxiosContentType = "application/json";

/**
 * 请求超时时间
 */
export const REQUEST_TIMEOUT = 60000;

/**
 * 不重定向白名单
 */
export const NO_REDIRECT_WHITE_LIST = ["/login", "/404"];

/**
 * 不重置路由白名单
 */
export const NO_RESET_WHITE_LIST = ["Redirect", "Login", "NoFind", "Root"];

/**
 * 是否根据headers->content-type自动转换数据格式
 */
export const TRANSFORM_REQUEST_DATA = true;

// 初始化网页标题名称
export const WEB_TITLE = "SuperAdmin";
