package api

import (
	"gitee.com/xygfm/authorization/apps/place"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) ListArea(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	area, err := h.svc.ListArea(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取区域列表", area))
	return
}

func (h *handler) CreateArea(ctx *gin.Context) {
	var req place.Area
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	_, err = h.svc.CreateArea(ctx, req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建成功", ""))
	return
}
func (h *handler) UpdateArea(ctx *gin.Context) {
	var req place.Area
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	req.ID = cast.ToInt(ctx.Param("id"))
	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("没有权限"))
		return
	}
	err = h.svc.UpdateArea(ctx, req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
	return
}

func (h *handler) DeleteArea(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteArea(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}
