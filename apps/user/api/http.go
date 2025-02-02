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
	r.POST("/register", h.Register)
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
	//删除员工
	r.DELETE("/staff/:id", h.DeleteStaff)
	//根据id获取员工信息
	r.GET("/staff/:id", h.GetStaff)
	//获取商户列表
	r.GET("/enterprise/list", h.ListEnterprise)
	//根据id获取商户信息
	r.GET("/enterprise/:id", h.GetEnterprise)
	//修改商户信息
	r.PUT("/enterprise/:id", h.UpdateEnterprise)
	//删除商户
	r.DELETE("/enterprise/:id", h.DeleteEnterprise)
	//根据id获取学生信息
	r.GET("/student/:id", h.GetStudent)
	//获取学生列表
	r.GET("/student/list", h.ListStudent)
	//修改学生信息
	r.PUT("/student/:id", h.UpdateStudent)
	//删除学生信息
	r.DELETE("/student/:id", h.DeleteStudent)
	return nil
}
