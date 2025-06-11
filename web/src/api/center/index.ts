import request from "@/axios";
export const postTask = (data): Promise<any> => {
  return request.post({ url: "/position", data });
};
export const getTaskList = (id,params): Promise<any> => {
  return request.get({ url: "/position/list/" +id, params });
};

export const getTaskById = (id): Promise<any> => {
  return request.get({ url: "/position/" +id,  });
};

export const editTaskById = (id,data): Promise<any> => {
  return request.put({ url: "/position/" +id,data  });
};
export const deleteTaskById = (id,data)=>{
  return request.delete({ url: "/position/" +id,data  });
  
}
export const getSelfPost = ()=>{
  return request.get({ url: "/position/list" });
  
}

export const getOffTask = (id)=>{
  return request.put({ url: "/position/putaway/"+id });
  
}
export const complainJob = (id,data)=>{
  return request.post({ url: "/audit/job/"+id,data });
  
}

export const createConversation = (params)=>{
  return request.get({ url: "/message/conv_id",params });
  
}
export const  getChatList=  (params)=>{
  return request.get({ url: "/message/conversation/list",params });
  
}
export const getChatRecord = (userId,chatId,params)=>{
  return request.get({ url:`/message/${userId}/list/${chatId}`, params});
}
export const sendMessage  = (data)=>{
  return request.post({ url:`/message/send`, data});
}
export const comfirmTop  = (id :number)=>{
  return request.put({ url:`/message/conversation/top/`+id});
}
export const cancelTop  = (id :number)=>{
  return request.delete({ url:`/message/conversation/top/`+id});
}
export const complainUser =  (id :number,data)=>{
  return request.post({ url:`/audit/user/`+id,data});
}
export const deletChat =  (id :number,)=>{
  return request.delete({ url:`/message/conversation/`+id,});
}

export const getUserComplainList = ()=>{
  return request.get({ url:`/audit/private/list`,});
}
export const deletComplainUser = (id)=>{
  return request.delete({ url:`/audit/`+id,});
}
export const editStudentInfo = (id,data)=>{
  return request.put({ url:`/user/student/`+id,data});
}
export const editCompanyInfo = (id,data)=>{
  return request.put({ url:`/user/enterprise/`+id,data});
}
