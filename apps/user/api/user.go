package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
)

func (h *handler) Login(ctx *gin.Context) {
	fmt.Println("我是登录Api层")
	token, err := h.svc.Login(ctx, &user.LoginRequest{Username: "", Password: ""})
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
	}
	response.Success(ctx, result.NewCorrect("登录成功", token))
}

func (h *handler) Register(ctx *gin.Context) {
	fmt.Println("我是Register Api层")
	token, err := h.svc.Register(ctx, &user.LoginRequest{Username: "", Password: ""})
	if err != nil {
		fmt.Println("adjnjksjndksndkjas")
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("登录成功", token))
	return
}

func (h *handler) ListMenu(ctx *gin.Context) {
	fmt.Println("进入获取菜单的api层")
	var ids = []int{1, 2, 3}
	menu, err := h.svc.ListMenu(ctx, &ids)
	if err != nil {
		fmt.Println("adjnjksjndksndkjas")
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取菜单成功", menu))
	return
}
