package api

import (
	"gitee.com/xygfm/authorization/apps/oss"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc oss.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(oss.AppName).(oss.Service)
	return nil
}

func (h *handler) Name() string {
	return oss.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	// 登录不进行任何中间件校验
	r.Static("/uploads", "./storage/uploads")
	r.Use(service.Security())
	r.GET("/list", h.List)
	r.POST("/upload", h.UploadFiles)
	return nil
}
