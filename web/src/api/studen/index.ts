import request from "@/axios";

export const getUserList = (params:{page:Number,size:Number,search:String}): Promise<any> => {
  return request.get({ url: "/user/student/list",params  });
};