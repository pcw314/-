package api

import (
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc user.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(user.AppName).(user.Service)
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	// 登录不进行任何中间件校验
	r.POST("/login", h.Login)
	r.GET("/menu", h.ListMenu)
	r.Use(service.Security())
	r.POST("/register", h.Register)
	return nil
}
