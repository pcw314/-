import { RawAxiosRequestHeaders } from "axios";
declare global {
  declare interface Fn<T = any> {
    (...arg: T[]): T;
  }
  declare type TimeoutHandle = ReturnType<typeof setTimeout>;
  declare type IntervalHandle = ReturnType<typeof setInterval>;
  declare type AxiosContentType =
    | "application/json"
    | "application/x-www-form-urlencoded"
    | "multipart/form-data"
    | "text/plain";

  declare type AxiosMethod = "get" | "post" | "delete" | "put";

  declare type AxiosResponseType =
    | "arraybuffer"
    | "blob"
    | "document"
    | "json"
    | "text"
    | "stream";

  declare interface AxiosConfig {
    params?: any;
    data?: any;
    url?: string;
    method?: AxiosMethod;
    headers?: RawAxiosRequestHeaders;
    responseType?: AxiosResponseType;
  }

  declare interface ErrorResponse {
    code: number;
    message:string,
    metadata:string,
    reason:string
  }
  declare type Recordable<T = any, K = string> = Record<
    K extends null | undefined ? string : K,
    T
  >;
}
