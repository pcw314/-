package user

import (
	"context"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	// Login 登陆
	Login(ctx context.Context, req *LoginRequest) (string, error)
	// Register 注册，创建
	Register(ctx *gin.Context, req *User) (*User, error)
	// CreateStudent 创建学生信息
	CreateStudent(ctx *gin.Context, req *Student) (*Student, error)
	// CreateEnterprise 创建商户信息
	CreateEnterprise(ctx *gin.Context, req *Enterprise) (*Enterprise, error)
	// ListMenu 根据权限id数字，获取菜单目录
	ListMenu(ctx context.Context, req *[]int) ([]*MenuRequest, error)
	// ListStaff 获取员工列表
	ListStaff(ctx *gin.Context, req *response.Paging) ([]*User, int64, error)
	// ListStudent 获取学生列表
	ListStudent(ctx *gin.Context, req *response.Paging) ([]*Student, int64, error)
	// ListEnterprise 获取商户列表
	ListEnterprise(ctx *gin.Context, req *response.Paging) ([]*Enterprise, int64, error)
	// GetStaffByID 根据id获取员工信息
	GetStaffByID(ctx *gin.Context, id int) (*User, error)
	// GetStudentByID 根据id获取学生信息
	GetStudentByID(ctx *gin.Context, id int) (*Student, error)
	// GetEnterpriseByID 根据id获取商户信息
	GetEnterpriseByID(ctx *gin.Context, id int) (*Enterprise, error)
	UpdateStudent(ctx *gin.Context, req *Student) (*Student, error)
	UpdateEnterprise(ctx *gin.Context, req *Enterprise) (*Enterprise, error)
	UpdatePassword(ctx *gin.Context, id int, nowPassword, oldPassword string) error
	DeleteStudent(ctx *gin.Context, id int) error
	DeleteEnterprise(ctx *gin.Context, id int) error
	DeleteStaff(ctx *gin.Context, id int) error
}

type ServiceImpl struct{}

func (s *ServiceImpl) Login(ctx context.Context, req *LoginRequest) (string, error) {
	return "", nil
}

func (s *ServiceImpl) Register(ctx *gin.Context, req *User) (*User, error) {
	return nil, nil
}
func (s *ServiceImpl) CreateStudent(ctx *gin.Context, req *Student) (*Student, error) {
	return nil, nil
}

func (s *ServiceImpl) CreateEnterprise(ctx *gin.Context, req *Enterprise) (*Enterprise, error) {
	return nil, nil
}

func (s *ServiceImpl) ListMenu(ctx context.Context, req *[]int) ([]*MenuRequest, error) {
	return nil, nil
}

func (s *ServiceImpl) ListStaff(ctx *gin.Context, req *[]int) ([]*User, int64, error) {
	return nil, 0, nil
}

func (s *ServiceImpl) ListStudent(ctx *gin.Context, req *response.Paging) ([]*Student, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) ListEnterprise(ctx *gin.Context, req *response.Paging) ([]*Enterprise, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) GetStaffByID(ctx *gin.Context, id int) (*User, error) {
	return nil, nil
}
func (s *ServiceImpl) GetStudentByID(ctx *gin.Context, id int) (*Student, error) {
	return nil, nil
}
func (s *ServiceImpl) GetEnterpriseByID(ctx *gin.Context, id int) (*Enterprise, error) {
	return nil, nil
}
func (s *ServiceImpl) UpdateStudent(ctx *gin.Context, req *Student) (*Student, error) {
	return nil, nil
}
func (s *ServiceImpl) UpdateEnterprise(ctx *gin.Context, req *Enterprise) (*Enterprise, error) {
	return nil, nil
}
func (s *ServiceImpl) UpdatePassword(ctx *gin.Context, id int, nowPassword, oldPassword string) error {
	return nil
}
func (s *ServiceImpl) DeleteStudent(ctx *gin.Context, id int) error {
	return nil
}

func (s *ServiceImpl) DeleteEnterprise(ctx *gin.Context, id int) error {
	return nil
}
func (s *ServiceImpl) DeleteStaff(ctx *gin.Context, id int) error {
	return nil
}
