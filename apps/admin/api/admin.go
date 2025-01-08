package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
)

func (h *handler) Login(ctx *gin.Context) {
	fmt.Println("我是登录Api层")
	token, err := h.svc.Login(ctx, &admin.LoginRequest{Username: "", Password: ""})
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
	}
	response.Success(ctx, result.NewCorrect("登录成功", token))
}
