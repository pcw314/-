import request from "@/axios";


export const deleteLog = (id:Number,): Promise<any> => {
    return request.delete({ url: "/log/"+id, });
  };

  export const getLogList = (params:any): Promise<any> => {
    return request.get({ url: "/log/list",params });
  };

