import axios, { AxiosError } from "axios";
import {
  defaultRequestInterceptors,
  defaultResponseInterceptors,
} from "./config";

import {
  AxiosInstance,
  InternalAxiosRequestConfig,
  RequestConfig,
  AxiosResponse,
} from "./types";
import { ElMessage } from "element-plus";
import { REQUEST_TIMEOUT } from "@/constants";

export const PATH_URL = import.meta.env.VITE_APP_BASE_API;
// export const PATH_URL = import.meta.env.VITE_SERVER;


const abortControllerMap: Map<string, AbortController> = new Map();

const axiosInstance: AxiosInstance = axios.create({
  timeout: REQUEST_TIMEOUT,
  baseURL: PATH_URL,
});

axiosInstance.interceptors.request.use((res: InternalAxiosRequestConfig) => {
  const controller = new AbortController();
  const url = res.url || "";
  res.signal = controller.signal;
  abortControllerMap.set(url, controller);
  return res;
});

axiosInstance.interceptors.response.use(
  (res: AxiosResponse) => {
    const url = res.config.url || "";
    abortControllerMap.delete(url);
    // 这里不能做任何处理，否则后面的 interceptors 拿不到完整的上下文了
    return res;
  },
  (err: AxiosError) => {
   let data = err.response!.data as ErrorResponse
    ElMessage.error(data.code +":" +data.message);
    return Promise.reject(err);
  }
);

axiosInstance.interceptors.request.use(defaultRequestInterceptors);
axiosInstance.interceptors.response.use(defaultResponseInterceptors);

const service = {
  request: (config: RequestConfig) => {
    return new Promise((resolve, reject) => {
      if (config.interceptors?.requestInterceptors) {
        config = config.interceptors.requestInterceptors(config as any);
      }

      axiosInstance
        .request(config)
        .then((res) => {
          resolve(res);
        })
        .catch((err: any) => {
 
          reject(err);
        });
    });
  },
  cancelRequest: (url: string | string[]) => {
    const urlList = Array.isArray(url) ? url : [url];
    for (const _url of urlList) {
      abortControllerMap.get(_url)?.abort();
      abortControllerMap.delete(_url);
    }
  },
  cancelAllRequest() {
    for (const [_, controller] of abortControllerMap) {
      controller.abort();
    }
    abortControllerMap.clear();
  },
};

export default service;
