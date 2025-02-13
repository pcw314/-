package api

import (
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) ListAuditJob(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	enterprise, err := h.svc.GetEnterpriseByUserID(ctx, utils.GetUserID(ctx))
	if err != nil {
		return
	}
	if list, total, err := h.svc.ListJob(ctx, &req, enterprise.ID); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
			Size:   req.Size,
			Page:   req.Page,
			List:   list,
			Total:  total,
			Search: req.Search,
		}))
	}
	return
}

func (h *handler) CreateJob(ctx *gin.Context) {
	var req position.Job
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	if utils.GetUserRole(ctx) != 3 && utils.GetUserRole(ctx) != 2 {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	_, err = h.svc.CreateJob(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建成功", ""))
	return
}
func (h *handler) UpdateJob(ctx *gin.Context) {
	var req position.Job
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	req.ID = cast.ToInt(ctx.Param("id"))
	if utils.GetUserRole(ctx) != 3 && utils.GetUserRole(ctx) != 2 {
		response.Error(ctx, result.DefaultError("没有权限"))
		return
	}
	err = h.svc.UpdateJob(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
	return
}

func (h *handler) DeleteJob(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteJob(ctx, id)
	if utils.GetUserRole(ctx) != 3 && utils.GetUserRole(ctx) != 2 {
		response.Error(ctx, result.DefaultError("没有权限"))
		return
	}
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}

func (h *handler) ListJobBySchoolID(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	schoolID := cast.ToInt(ctx.Param("school_id"))

	if list, total, err := h.svc.ListJobBySchoolID(ctx, &req, schoolID); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
			Size:   req.Size,
			Page:   req.Page,
			List:   list,
			Total:  total,
			Search: req.Search,
		}))
	}
	return
}
