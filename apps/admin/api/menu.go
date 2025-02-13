package api

import (
	"gitee.com/xygfm/authorization/apps/admin"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) ListMenu(ctx *gin.Context) {
	var ids []int
	if utils.GetUserRole(ctx) == 1 {
		ids = admin.StudentMenu

	} else if utils.GetUserRole(ctx) == 2 {
		ids = admin.Enterprise

	} else if utils.GetUserRole(ctx) == 3 {
		ids = admin.AdminMenu
	}
	menu, err := h.svc.ListMenu(ctx, ids)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取菜单成功", menu))
	return
}

func (h *handler) CreateMenu(ctx *gin.Context) {
	var req admin.Menu
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	_, err = h.svc.CreateMenu(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建成功", ""))
	return
}
func (h *handler) UpdateMenu(ctx *gin.Context) {
	var req admin.Menu
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	req.ID = cast.ToInt(ctx.Param("id"))
	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("没有权限"))
		return
	}
	err = h.svc.UpdateMenu(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
	return
}

func (h *handler) DeleteMenu(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteMenu(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}
