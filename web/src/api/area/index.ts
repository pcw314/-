
import request from "@/axios";
export const areaRegionList = (id=0): Promise<any> => {
  return request.get({ url: "/place/area/"+id, });
};
export const createAreaRegionList = (data): Promise<any> => {
    return request.post({ url: "/place/area",data });
  };
  export const deleteAreaRegionList = (id): Promise<any> => {
    return request.delete({ url: "/place/area/" +id, });
  };
  export const editAreaRegionList = (id,data): Promise<any> => {
    return request.put({ url: "/place/area/" +id,data });
  };
  export const getSchoolList = (id=0,params:any): Promise<any> => {
    return request.put({ url: "/place/school/" +id, params});
  };
  export const adminAddSchool = (data:any): Promise<any> => {
    return request.post({ url: "/place/school" , data});
  };
  export const editAddSchoolApi = (id,data:any): Promise<any> => {
    return request.put({ url: "/place/school/"+id , data});
  };
  export const deleteSchoolApi = (id,): Promise<any> => {
    return request.delete({ url: "/place/school/"+id , });
  };
