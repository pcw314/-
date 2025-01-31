package place

import (
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	// ListArea 树形结构
	ListArea(ctx *gin.Context, id int) ([]*Area, error)
	UpdateArea(ctx *gin.Context, req Area) error
	CreateArea(ctx *gin.Context, req Area) (*Area, error)
	DeleteArea(ctx *gin.Context, id int) error
	ListSchool(ctx *gin.Context, req *response.Paging, area int) ([]*School, int64, error)
	UpdateSchool(ctx *gin.Context, req School) error
	CreateSchool(ctx *gin.Context, req School) (*School, error)
	DeleteSchool(ctx *gin.Context, id int) error
}

type ServiceImpl struct{}

func (s *ServiceImpl) ListArea(ctx *gin.Context, id int) ([]*Area, error) {
	return nil, nil
}

func (s *ServiceImpl) UpdateArea(ctx *gin.Context, req Area) error {
	return nil
}
func (s *ServiceImpl) CreateArea(ctx *gin.Context, req Area) (*Area, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteArea(ctx *gin.Context, id int) error {
	return nil
}
func (s *ServiceImpl) ListSchool(ctx *gin.Context, req *response.Paging, area int) ([]*School, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) UpdateSchool(ctx *gin.Context, req School) error {
	return nil
}
func (s *ServiceImpl) CreateSchool(ctx *gin.Context, req School) (*School, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteSchool(ctx *gin.Context, id int) error {
	return nil
}
