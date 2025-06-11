import request from "@/axios";
export const getMenuList = (): Promise<any> => {
  return request.get({ url: "/admin/menu/list" });
};
export const createMnue = (data): Promise<any> => {
  return request.post({ url: "/admin/menu", data });
};
export const editMenue = (id, data): Promise<any> => {
  return request.put({ url: "/admin/menu/" + id, data });
};
export const deletMenu = (id, ): Promise<any> => {
  return request.delete({ url: "/admin/menu/" + id, });
};

