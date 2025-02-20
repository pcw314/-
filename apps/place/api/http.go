package api

import (
	"gitee.com/xygfm/authorization/apps/place"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc place.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(place.AppName).(place.Service)
	return nil
}

func (h *handler) Name() string {
	return place.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	r.Use(service.Security())
	area := r.Group("/area")
	area.PUT("/:id", h.UpdateArea)
	area.POST("", h.CreateArea)
	area.DELETE("/:id", h.DeleteArea)
	r.GET("/area/:id", h.ListArea)
	school := r.Group("/school")
	school.GET("/:area", h.ListSchool)
	school.PUT("/:id", h.UpdateSchool)
	school.POST("", h.CreateSchool)
	school.DELETE("/:id", h.DeleteSchool)
	school.GET("/top", h.ListTopSchool)
	return nil
}
