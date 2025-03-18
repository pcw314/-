package admin

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetUser(ctx *gin.Context, req *LoginRequest) (*User, error)
	// Register 注册
	Register(ctx *gin.Context, req *LoginRequest) (string, error)
	// Logout 退出登录
	Logout(ctx *gin.Context) error
	// ListMenu 获取菜单
	ListMenu(ctx *gin.Context, req []int) ([]*Menu, error)
	// UpdateMenu 修改菜单
	UpdateMenu(ctx *gin.Context, req *Menu) error
	// CreateMenu 创建菜单
	CreateMenu(ctx *gin.Context, req *Menu) (*Menu, error)
	// DeleteMenu 删除菜单
	DeleteMenu(ctx *gin.Context, id int) error
	// ListArea 获取全部地区
	ListArea(ctx *gin.Context, id string) ([]*Areas, error)
	GetStatistics(ctx *gin.Context) (*Statistics, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) GetUser(ctx *gin.Context, req *LoginRequest) (*User, error) {
	return nil, nil
}

func (s *ServiceImpl) Register(ctx *gin.Context, req *LoginRequest) (string, error) {
	return "", nil
}

func (s *ServiceImpl) Logout(ctx *gin.Context) error {
	return nil
}

func (s *ServiceImpl) ListMenu(ctx *gin.Context, req []int) ([]*Menu, error) {
	return nil, nil
}

func (s *ServiceImpl) UpdateMenu(ctx *gin.Context, req *Menu) error {
	return nil
}
func (s *ServiceImpl) CreateMenu(ctx *gin.Context, req *Menu) (*Menu, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteMenu(ctx *gin.Context, id int) error {
	return nil
}

func (s *ServiceImpl) ListArea(ctx *gin.Context, id string) ([]*Areas, error) {
	return nil, nil
}

func (s *ServiceImpl) GetStatistics(ctx *gin.Context) (*Statistics, error) {
	return nil, nil
}
