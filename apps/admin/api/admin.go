package api

import (
	"errors"
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	service "gitee.com/xygfm/authorization/middleware"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *handler) Login(ctx *gin.Context) {
	fmt.Println("我是登录Api层")
	var req admin.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	fmt.Println("req", req)
	user, err := h.svc.GetUser(ctx, &admin.LoginRequest{Username: req.Username, Role: req.Role})
	if err != nil {
		response.Error(ctx, result.DefaultError("用户名或密码错误"))
		return
	}
	fmt.Println("user", user)
	if user.Password == "" || !utils.BcryptCheck(req.Password, user.Password) {
		response.Error(ctx, result.DefaultError("用户名或密码错误"))
		return
	}
	token, err := service.MakeToken(user.ID, user.Username, user.Role)
	response.Success(ctx, result.NewCorrect("登录成功", token))
}

func (h *handler) Register(ctx *gin.Context) {
	fmt.Println("我是Register Api层")
	var req admin.LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if req.Password != "" {
		req.Password = utils.BcryptHash(req.Password)
	} else {
		response.Error(ctx, result.DefaultError("请输入密码"))
	}

	_, err = h.svc.GetUser(ctx, &admin.LoginRequest{Username: req.Username, Role: req.Role})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			token, err := h.svc.Register(ctx, &req)
			if err != nil {
				response.Error(ctx, result.DefaultError(err.Error()))
				return
			}
			response.Success(ctx, result.NewCorrect("注册成功", token))
			return
		}
		response.Error(ctx, result.DefaultError(err.Error()))
	}
	response.Error(ctx, result.DefaultError("用户名已存在"))
	return

}

func (h *handler) Logout(ctx *gin.Context) {
	response.Success(ctx, result.NewCorrect("退出成功", ""))
	return
}

func (h *handler) ListArea(ctx *gin.Context) {
	id := ctx.Param("id")
	area, err := h.svc.ListArea(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取区域列表", area))
	return
}
