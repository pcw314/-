package api

import (
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc position.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(position.AppName).(position.Service)
	return nil
}

func (h *handler) Name() string {
	return position.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	// 登录不进行任何中间件校验
	r.Use(service.Security())
	//获取商户招聘信息
	r.GET("/list", h.ListAuditJob)
	r.GET("/list/:school_id", h.ListJobBySchoolID)
	r.GET("/:id", h.GetJobByID)
	r.PUT("/:id", h.UpdateJob)
	r.PUT("/putaway/:id", h.SetJobState)
	r.DELETE("/:id", h.DeleteJob)
	r.POST("", h.CreateJob)
	collect := r.Group("/collect")
	collect.GET("/list", h.ListCollect)
	collect.POST("/:id", h.CreateCollect)
	collect.DELETE("/:id", h.DeleteCollect)
	return nil
}
