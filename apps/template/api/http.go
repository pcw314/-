package api

import (
	"gitee.com/xygfm/authorization/apps/template"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc template.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(template.AppName).(template.Service)
	return nil
}

func (h *handler) Name() string {
	return template.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	// 登录不进行任何中间件校验
	r.Use(service.Security())
	r.GET("/list", h.List)
	return nil
}
