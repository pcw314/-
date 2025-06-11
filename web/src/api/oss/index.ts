import request from "@/axios";


export const deleteOss = (id:Number,): Promise<any> => {
    return request.delete({ url: "/oss/"+id, });
  };

  export const getOssList = (params:any): Promise<any> => {
    return request.get({ url: "/oss/list",params });
  };

