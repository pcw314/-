package api

import (
	"gitee.com/xygfm/authorization/apps/log"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc log.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(log.AppName).(log.Service)
	return nil
}

func (h *handler) Name() string {
	return log.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	r.Use(service.Security())
	r.GET("/list", h.ListLog)
	r.DELETE("/:id", h.DeleteLog)
	return nil
}
