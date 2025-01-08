package ioc

import "github.com/gin-gonic/gin"

type Object interface {
	Init() error
	Name() string
	RegisterRoute(r gin.IRouter) error
}
