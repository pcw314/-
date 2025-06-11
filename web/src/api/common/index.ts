import request from "@/axios";

export const getUserInfoById = (data: Number): Promise<any> => {
  return request.get({ url: "/user/student/" + data });
};
export const getCompanyInfo = (data: Number): Promise<any> => {
  return request.get({ url: "/user/enterprise/" + data });
};
export const uploadFile = (data: Object): Promise<any> => {
  return request.post({
    url: "/oss/upload",
    data,
    headers: { "Content-Type": "mulmultipart/form-data" },
  });
};
export const getProvinceList = (data: Number): Promise<any> => {
  return request.get({ url: "/place/area/" + data });
};
export const getSchoolList = (data: Number,params): Promise<any> => {
  return request.get({ url: "/place/school/" + data ,params});
};

export const editCompanyInfo = (id: Number, data: any): Promise<any> => {
  return request.put({ url: "/user/enterprise/" + id, data });
};

export const getHotSchool = (data): Promise<any> => {
  return request.get({ url: "/place/school/top", data });
};
export const editStudentInfo = (id: Number, data: any): Promise<any> => {
  return request.put({ url: "/user/student/" + id, data });
};
export const collectTaskById = (id: Number): Promise<any> => {
  return request.post({ url: "/position/collect/" + id });
};
export const cancelCollectTaskById = (id: Number): Promise<any> => {
  return request.delete({ url: "/position/collect/" + id });
};

export const getCollectList = (): Promise<any> => {
  return request.get({ url: "/position/collect/list" });
};

export const deleteStudentByID = (id: Number): Promise<any> => {
  return request.delete({ url: "/user/student/" + id });
};

export const adminCreateStudent = (data): Promise<any> => {
  return request.post({ url: "/user/student", data });
};
export const banUserById = (id): Promise<any> => {
  return request.put({ url: `/user/${id}/set_state` });
};
export const initUserPassword = (id): Promise<any> => {
  return request.put({ url: `/user/${id}/set_password` });
};

export const getCompanyList = (params): Promise<any> => {
  return request.get({ url: `/user/enterprise/list`,params },);
};

export const adminCreateCompany = (data): Promise<any> => {
  return request.post({ url: `/user/enterprise`, data });
};

export const deleteCompany = (id): Promise<any> => {
  return request.delete({ url: `/user/enterprise/` + id });
};
export const getComapnyStaffList = (params :any): Promise<any> => {
  return request.get({ url: `/user/staff/list`,params  },);
};
export const deleteStaffById =  (id): Promise<any> => {
  return request.delete({ url: `/user/staff/`+id  });
};
export const editStaffById =  (id,data): Promise<any> => {
  return request.put({ url: `/user/staff/`+id,data  });
};
export const  adminCreateStaff = (data): Promise<any> => {
  return request.post({ url: `/user/staff`,data  });
};
export const  getStaffInfo = (id=0): Promise<any> => {
  return request.get({ url: `/user/staff/`+id  });
};
export const jobComplainList = (params)=>{
  return request.get({ url: '/audit/job/list',params });
}
export const getDashBoard = ()=>{
  return request.get({ url: '/admin/statistics', });
}
