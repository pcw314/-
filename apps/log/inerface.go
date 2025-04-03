package log

import (
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	ListLog(ctx *gin.Context, req *response.Paging) ([]*Log, int64, error)
	DeleteLog(ctx *gin.Context, id int) error
}

type ServiceImpl struct{}

func (s *ServiceImpl) ListLog(ctx *gin.Context, req *response.Paging) ([]*Log, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) DeleteLog(ctx *gin.Context, id int) error {
	return nil
}
