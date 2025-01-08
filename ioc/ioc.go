package ioc

import (
	"fmt"
	"gitee.com/xygfm/authorization/logger"
	"github.com/gin-gonic/gin"
	"path"
)

var (
	controllerObjects = map[string]Object{}
	httpObjects       = map[string]Object{}
	grpcObjects       = map[string]Object{}
)

// RegisterController 将传进来的对象注册到全局
func RegisterController(obj Object) error {
	controllerObjects[obj.Name()] = obj
	return nil
}
func GetControllerObjects() map[string]Object {
	return controllerObjects
}

func GetControllerObject(name string) Object {
	return controllerObjects[name]
}

func RegisterHttp(obj Object) error {
	httpObjects[obj.Name()] = obj
	return nil
}

func GetHttpObjects() map[string]Object {
	return httpObjects
}

func GetHttpObject(name string) Object {
	return httpObjects[name]
}

func RegisterGrpc(obj Object) error {
	grpcObjects[obj.Name()] = obj
	return nil
}

func GetGrpcObjects() map[string]Object {
	return grpcObjects
}

func GetGrpcObject(name string) Object {
	return grpcObjects[name]
}

func InitControllerObjects() error {
	for name, obj := range controllerObjects {
		err := obj.Init()
		if err != nil {
			return err
		}
		logger.Info(fmt.Sprintf("初始化%s controller成功", name))
	}
	return nil
}

func InitHttpObjects() error {
	for name, obj := range httpObjects {
		err := obj.Init()
		if err != nil {
			return err
		}
		logger.Info(fmt.Sprintf("初始化%s Http成功", name))
	}
	return nil
}

func InitGrpcObjects() error {
	for name, obj := range grpcObjects {
		err := obj.Init()
		if err != nil {
			return err
		}
		logger.Info(fmt.Sprintf("初始化%s grpc成功", name))
	}

	return nil
}

func LoadHttpObjects(pathPrefix string, r gin.IRouter) error {
	for name, obj := range httpObjects {
		err := obj.RegisterRoute(r.Group(path.Join(pathPrefix, name)))
		if err != nil {
			return err
		}
		logger.Info(fmt.Sprintf("加载%s http成功", name))
	}
	return nil
}

type ObjectImpl struct {
}

func (i *ObjectImpl) Init() error {
	return nil
}

func (i *ObjectImpl) Name() string {
	return ""
}

func (i *ObjectImpl) RegisterRoute(r gin.IRouter) error {
	return nil
}
