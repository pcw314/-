package api

import (
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/ioc"
	service "gitee.com/xygfm/authorization/middleware"
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
	r.Use(service.LoggerMiddleware())
	// 登录不进行任何中间件校验
	r.POST("/register", h.Register)
	//修改密码
	r.PUT("/set_password/:id", h.UpdatePassword)
	//进行检验
	r.Use(service.Security())
	//学生创建
	r.POST("/student", h.CreateStudent)
	//商户创建
	r.POST("/enterprise", h.CreateEnterprise)
	//员工创建
	r.POST("/staff", h.CreateStaff)
	//初始化密码
	r.PUT("/:id/set_password", h.InitializePassword)
	//修改状态
	r.PUT("/:id/set_state", h.UpdateUserState)
	//获取员工列表
	r.GET("/staff/list", h.ListStaff)

	//删除员工
	r.DELETE("/staff/:id", h.DeleteStaff)
	//根据id获取员工信息
	r.GET("/staff/:id", h.GetStaff)
	//根据
	r.PUT("/staff/:id", h.UpdateStaff)
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
