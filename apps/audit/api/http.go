package api

import (
	"gitee.com/xygfm/authorization/apps/audit"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc audit.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(audit.AppName).(audit.Service)
	return nil
}

func (h *handler) Name() string {
	return audit.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	// 登录不进行任何中间件校验
	r.Use(service.Security())
	r.GET("/enterprise/list", h.ListAuditEnterprise)
	r.GET("/job/list", h.ListAuditJob)
	r.PUT("/:id", h.UpdateAudit)
	r.DELETE("/:id", h.DeleteAudit)
	r.POST("", h.CreateAudit)
	return nil
}
