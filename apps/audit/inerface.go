package audit

import (
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	ListAudit(ctx *gin.Context, req *response.Paging, role int) ([]*Audit, int64, error)
	ListAuditUser(ctx *gin.Context, req *response.Paging) ([]*Audit, int64, error)
	ListAuditJob(ctx *gin.Context, req *response.Paging) ([]*Audit, int64, error)
	ListAuditJobInfo(ctx *gin.Context, req *response.Paging) ([]*position.Job, int64, error)
	UpdateAudit(ctx *gin.Context, req *Audit) error
	CreateAudit(ctx *gin.Context, req *Audit) (*Audit, error)
	DeleteAudit(ctx *gin.Context, id int) error
	CreateAuditStudent(ctx *gin.Context, req *Audit) (*Audit, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) ListAudit(ctx *gin.Context, req *response.Paging, role int) ([]*Audit, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) UpdateAudit(ctx *gin.Context, req *Audit) error {
	return nil
}
func (s *ServiceImpl) CreateAudit(ctx *gin.Context, req *Audit) (*Audit, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteAudit(ctx *gin.Context, id int) error {
	return nil
}
func (s *ServiceImpl) CreateAuditStudent(ctx *gin.Context, req *Audit) (*Audit, error) {
	return nil, nil
}
func (s *ServiceImpl) ListAuditJob(ctx *gin.Context, req *response.Paging) ([]*Audit, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) ListAuditUser(ctx *gin.Context, req *response.Paging) ([]*Audit, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) ListAuditJobInfo(ctx *gin.Context, req *response.Paging) ([]*position.Job, int64, error) {
	return nil, 0, nil
}
