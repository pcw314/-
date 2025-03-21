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
	GetJobByID(ctx *gin.Context, userID int) (*Job, error)
	SetJobState(ctx *gin.Context, id int) (string, error)
	GetEnterpriseByUserID(ctx *gin.Context, userID int) (*user.Enterprise, error)
	ListCollect(ctx *gin.Context, req *response.Paging, userID int) ([]*Job, int64, error)
	CreateCollect(ctx *gin.Context, collect Collect) (*Collect, error)
	DeleteCollect(ctx *gin.Context, id int) error
	ListJobBySchoolID(ctx *gin.Context, req *response.Paging, id int) ([]*Job, int64, error)
	ListAuditJob(ctx *gin.Context, req *AuditJobRequest) ([]*Job, int64, error)
	AuditJob(ctx *gin.Context, req *UpdateAuditJob) error
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
func (s *ServiceImpl) GetJobByID(ctx *gin.Context, ID int) (*Job, error) {
	return nil, nil
}
func (s *ServiceImpl) SetJobState(ctx *gin.Context, id int) (string, error) {
	return "", nil
}
func (s *ServiceImpl) GetEnterpriseByUserID(ctx *gin.Context, userID int) (*user.Enterprise, error) {
	return nil, nil
}
func (s *ServiceImpl) ListCollect(ctx *gin.Context, req *response.Paging, id int) ([]*Job, int64, error) {
	return nil, 0, nil
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
func (s *ServiceImpl) ListAuditJob(ctx *gin.Context, req *AuditJobRequest) ([]*Job, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) AuditJob(ctx *gin.Context, req *UpdateAuditJob) error {
	return nil
}
