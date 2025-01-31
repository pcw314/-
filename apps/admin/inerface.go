package admin

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	// Login 登录
	GetUser(ctx *gin.Context, req *LoginRequest) (*User, error)
	// Register 注册
	Register(ctx *gin.Context, req *LoginRequest) (string, error)
	// Logout 退出登录
	Logout(ctx *gin.Context) error
	// ListMenu 获取菜单
	ListMenu(ctx *gin.Context, req *[]int) ([]*MenuRequest, error)
	// ListArea 获取全部地区
	ListArea(ctx *gin.Context, id string) ([]*Areas, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) GetUser(ctx *gin.Context, req *LoginRequest) (*User, error) {
	return nil, nil
}

func (s *ServiceImpl) Register(ctx *gin.Context, req *LoginRequest) (string, error) {
	return "", nil
}

func (s *ServiceImpl) ListMenu(ctx *gin.Context, req *[]int) ([]*MenuRequest, error) {
	return nil, nil
}

func (s *ServiceImpl) Logout(ctx *gin.Context) error {
	return nil
}

func (s *ServiceImpl) ListArea(ctx *gin.Context, id string) ([]*Areas, error) {
	return nil, nil
}
