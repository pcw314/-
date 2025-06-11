
import request from "@/axios";
export const allowOrNOComplain = (id:Number,data : any): Promise<any> => {
  return request.put({ url: "/audit/"+id,data });
};
export const deleteComplain = (id:Number,): Promise<any> => {
    return request.delete({ url: "/audit/"+id, });
  };
  export const getJobList = (params:any): Promise<any> => {
    return request.get({ url: "/position/audit/list",params });
  };
  export const checkJob = (id,data:any): Promise<any> => {
    return request.put({ url: "/position/audit/"+id,data });
  };
    export const getUserComplain = (params): Promise<any> => {
    return request.get({ url: "/audit/user/list" ,params});
  };