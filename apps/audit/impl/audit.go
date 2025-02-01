package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/audit"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListAudit(ctx *gin.Context, req *response.Paging, role int) ([]*audit.Audit, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var po []*audit.Audit
	db := i.mdb.Model(&audit.Audit{})
	if req.Search != "" {
		db = db.Where("reason LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&po).Error
	if err != nil {
		return nil, 0, err
	}

	return po, total, nil
}
func (i *impl) UpdateAudit(ctx *gin.Context, req *audit.Audit) error {
	return nil
}
func (i *impl) CreateAudit(ctx *gin.Context, req *audit.Audit) (*audit.Audit, error) {
	return nil, nil
}
func (i *impl) DeleteAudit(ctx *gin.Context, req *audit.Audit) error {
	return nil
}
