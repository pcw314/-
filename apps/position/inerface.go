package position

import (
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	ListJob(ctx *gin.Context, req *response.Paging, enterpriseID int) ([]*Job, int64, error)
	UpdateJob(ctx *gin.Context, req *Job) error
	CreateJob(ctx *gin.Context, req *Job) (*Job, error)
	DeleteJob(ctx *gin.Context, id int) error
	GetEnterpriseByUserID(ctx *gin.Context, userID int) (*user.Enterprise, error)
	ListCollect(ctx *gin.Context, id int) ([]*Collect, error)
	CreateCollect(ctx *gin.Context, collect Collect) (*Collect, error)
	DeleteCollect(ctx *gin.Context, id int) error
	ListJobBySchoolID(ctx *gin.Context, req *response.Paging, id int) ([]*Job, int64, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) ListJob(ctx *gin.Context, req *response.Paging, enterpriseID int) ([]*Job, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) UpdateJob(ctx *gin.Context, req *Job) error {
	return nil
}
func (s *ServiceImpl) CreateJob(ctx *gin.Context, req *Job) (*Job, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteJob(ctx *gin.Context, id int) error {
	return nil
}

func (s *ServiceImpl) GetEnterpriseByUserID(ctx *gin.Context, userID int) (*user.Enterprise, error) {
	return nil, nil
}
func (s *ServiceImpl) ListCollect(ctx *gin.Context, id int) ([]*Collect, error) {
	return nil, nil
}
func (s *ServiceImpl) CreateCollect(ctx *gin.Context, collect Collect) (*Collect, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteCollect(ctx *gin.Context, id int) error {
	return nil
}
func (s *ServiceImpl) ListJobBySchoolID(ctx *gin.Context, req *response.Paging, id int) ([]*Job, int64, error) {
	return nil, 0, nil
}
