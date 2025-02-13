package api

import (
	"gitee.com/xygfm/authorization/apps/message"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ioc.ObjectImpl
	svc message.Service
}

func init() {
	err := ioc.RegisterHttp(&handler{})
	if err != nil {
		panic(err)
	}
}

func (h *handler) Init() error {
	h.svc = ioc.GetControllerObject(message.AppName).(message.Service)
	return nil
}

func (h *handler) Name() string {
	return message.AppName
}

func (h *handler) RegisterRoute(r gin.IRouter) error {
	// 登录不进行任何中间件校验
	r.Use(service.Security())
	//获取商户招聘信息
	r.GET("/ws", h.HandleWebSocket)
	r.GET("/ws/message/:to_user_id/list/:conv_id", h.ListMessages)
	r.GET("/ws/conversation/list", h.ListConversation)
	return nil
}
