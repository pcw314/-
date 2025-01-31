package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
)

func (h *handler) List(ctx *gin.Context) {
	fmt.Println("我是Api层")
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListName(ctx, &req); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("登录成功", response.Paging{
			Size:   req.Size,
			Page:   req.Page,
			List:   list,
			Total:  total,
			Search: req.Search,
		}))
	}

}
