package api

import (
	"gitee.com/xygfm/authorization/apps/admin"
	"gitee.com/xygfm/authorization/ioc"
	service "gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc admin.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(admin.AppName).(admin.Service)
	return nil
}

func (h *handler) Name() string {
	return admin.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	r.Use(service.LoggerMiddleware())
	// 登录不进行任何中间件校验
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.Use(service.Security())
	r.POST("/logout", h.Logout)
	r.GET("/statistics", h.GetStatistics)
	menu := r.Group("/menu")
	menu.GET("/list", h.ListMenu)
	menu.POST("", h.CreateMenu)
	menu.PUT("/:id", h.UpdateMenu)
	menu.DELETE("/:id", h.DeleteMenu)
	r.GET("/area/:id", h.ListArea)
	return nil
}
