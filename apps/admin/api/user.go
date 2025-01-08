package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
)

func (h *handler) Register(ctx *gin.Context) {
	fmt.Println("我是Register Api层")
	token, err := h.svc.Register(ctx, &admin.LoginRequest{Username: "", Password: ""})
	if err != nil {
		fmt.Println("adjnjksjndksndkjas")
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("登录成功", token))
	return
}
