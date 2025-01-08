package protocol

import (
	"context"
	"fmt"
	"gitee.com/xygfm/authorization/conf"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Http struct {
	Address string
	server  *http.Server
	Router  *gin.Engine
}

func NewHttp() *Http {
	r := gin.Default()
	// cors中间件
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	return &Http{
		Address: conf.GetConfig().Http.Address(),
		server: &http.Server{
			Addr:         conf.GetConfig().Http.Address(),
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
		Router: r,
	}
}

func (h *Http) PathPrefix() string {
	return ""
}

func (h *Http) Start() error {
	// 初始化Http对象
	err := ioc.InitHttpObjects()
	if err != nil {
		panic(err)
	}
	// 加载Http
	err = ioc.LoadHttpObjects(h.PathPrefix(), h.Router)
	if err != nil {
		panic(err)
	}
	logger.Info(fmt.Sprintf("http server: %s", h.Address))
	// 监听并启动server
	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	// 优雅的关闭
	return h.server.Shutdown(ctx)
}
