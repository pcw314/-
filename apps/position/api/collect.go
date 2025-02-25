package api

import (
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
)

func (h *handler) ListCollect(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	if list, total, err := h.svc.ListCollect(ctx, &req, utils.GetUserID(ctx)); err != nil {
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

func (h *handler) CreateCollect(ctx *gin.Context) {
	var req position.Collect
	req.JobID = cast.ToInt(ctx.Param("id"))
	//err := ctx.ShouldBindJSON(&req)
	//if err != nil {
	//	response.Error(ctx, result.DefaultError(err.Error()))
	//	return
	//}
	req.UserID = utils.GetUserID(ctx)
	req.CreatedAt = time.Now().UnixMilli()
	_, err := h.svc.CreateCollect(ctx, req)
	if err != nil {
		if err.Error() == "collect already exists" {
			response.Success(ctx, result.NewCorrect("取消收藏成功", ""))
			return
		}
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("收藏成功", ""))
	return
}

func (h *handler) DeleteCollect(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteCollect(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}
