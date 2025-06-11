import service from "./service";
import { CONTENT_TYPE } from "@/constants";
import { useUserStore } from "@/store/module/user";

const request = (option: AxiosConfig) => {
  const { url, method, params, data, headers, responseType } = option;
  const userStore = useUserStore();
  return service.request({
    url: url,
    method,
    params,
    data,
    responseType: responseType,
    headers: {
      "Content-Type": CONTENT_TYPE,
      Authorization:  userStore.getToken || "",
      ...headers,
    },
  });
};

export default {
  get: <T = any>(option: AxiosConfig) => {
    return request({ method: "get", ...option }) as Promise<any>;
  },
  post: <T = any>(option: AxiosConfig) => {
    return request({ method: "post", ...option }) as Promise<any>;
  },
  delete: <T = any>(option: AxiosConfig) => {
    return request({ method: "delete", ...option }) as Promise<any>;
  },
  put: <T = any>(option: AxiosConfig) => {
    return request({ method: "put", ...option }) as Promise<any>;
  },
  cancelRequest: (url: string | string[]) => {
    return service.cancelRequest(url);
  },
  cancelAllRequest: () => {
    return service.cancelAllRequest();
  },
};
