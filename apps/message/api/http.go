package api

import (
	"gitee.com/xygfm/authorization/apps/message"
	"gitee.com/xygfm/authorization/ioc"
	service "gitee.com/xygfm/authorization/middleware"
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
	r.Use(service.LoggerMiddleware())
	ws := r.Group("/ws")
	ws.Use(service.Ws())
	ws.GET("", h.HandleWebSocket)
	//ws.GET("/message/:to_user_id/list/:conv_id", h.ListMessages)
	ws.GET("/conversation/list", h.ListConversationWebSocket)
	r.Use(service.Security())
	r.POST("/send", h.SendMessage)
	r.GET("/conv_id", h.GetConversationID)
	r.GET("/:to_user_id/list/:conv_id", h.ListMessages)
	r.GET("/conversation/list", h.ListConversation)
	r.DELETE("/conversation/:conv_id", h.DeleteConversation)
	r.PUT("/conversation/top/:conv_id", h.ConversationTop)
	r.DELETE("/conversation/top/:conv_id", h.DeleteConversationTop)
	return nil
}
