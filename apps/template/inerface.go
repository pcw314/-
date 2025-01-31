package template

import (
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	ListName(ctx *gin.Context, req *response.Paging) ([]*Template, int64, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) ListName(ctx *gin.Context, req *response.Paging) ([]*Template, int64, error) {
	return nil, 0, nil
}
