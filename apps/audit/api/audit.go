package api

import (
	"gitee.com/xygfm/authorization/apps/audit"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) ListAudit(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListAudit(ctx, &req, utils.GetUserID(ctx)); err != nil {
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

func (h *handler) ListAuditUser(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListAuditUser(ctx, &req); err != nil {
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

func (h *handler) ListAuditJob(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListAuditJob(ctx, &req); err != nil {
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

func (h *handler) CreateAuditJob(ctx *gin.Context) {
	JobID := cast.ToInt(ctx.Param("id"))
	var req audit.Audit
	err := ctx.ShouldBindJSON(&req)
	req.JobID = JobID
	req.Informer = utils.GetUserID(ctx)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	_, err = h.svc.CreateAudit(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("举报成功", ""))
	return
}
func (h *handler) CreateAuditUser(ctx *gin.Context) {
	userID := cast.ToInt(ctx.Param("id"))
	var req audit.Audit
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	req.UserID = userID
	req.Informer = utils.GetUserID(ctx)
	_, err = h.svc.CreateAudit(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("举报成功", ""))
	return
}
func (h *handler) UpdateAudit(ctx *gin.Context) {
	var req audit.Audit
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
	err = h.svc.UpdateAudit(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
	return
}

func (h *handler) DeleteAudit(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteAudit(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}
