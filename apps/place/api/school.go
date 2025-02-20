package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/place"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) ListSchool(ctx *gin.Context) {
	fmt.Println("我是Api层")
	area := cast.ToInt(ctx.Param("area"))
	fmt.Println("area", area)
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListSchool(ctx, &req, area); err != nil {
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

func (h *handler) CreateSchool(ctx *gin.Context) {
	var req place.School
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	_, err = h.svc.CreateSchool(ctx, req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建成功", ""))
	return
}
func (h *handler) UpdateSchool(ctx *gin.Context) {
	var req place.School
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
	err = h.svc.UpdateSchool(ctx, req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
	return
}

func (h *handler) DeleteSchool(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteSchool(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}

func (h *handler) ListTopSchool(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, err := h.svc.ListHotSchool(ctx, &req); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
			Size: req.Size,
			Page: req.Page,
			List: list,
		}))
	}
	return
}
