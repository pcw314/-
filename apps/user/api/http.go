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
	//学生注册
	r.POST("/student/create", h.CreateStudent)
	//企业注册
	r.POST("/enterprise/create", h.CreateEnterprise)
	//进行检验
	r.Use(service.Security())
	//创建员工
	r.POST("/staff/create", h.CreateStaff)
	r.GET("/staff/list", h.ListStaff)
	//修改密码
	r.PUT("/staff/:id", h.UpdatePassword)
	r.DELETE("/staff/:id", h.DeleteStaff)
	r.GET("/staff/:id", h.GetStaff)
	r.GET("/enterprise/list", h.ListEnterprise)
	r.GET("/enterprise/:id", h.GetEnterprise)
	r.PUT("/enterprise/:id", h.UpdateEnterprise)
	r.DELETE("/enterprise/:id", h.DeleteEnterprise)
	r.GET("/student/:id", h.GetStudent)
	r.GET("/student/list", h.ListStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)
	return nil
}
