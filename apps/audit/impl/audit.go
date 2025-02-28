package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/audit"
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"time"
)

func (i *impl) ListAudit(ctx *gin.Context, req *response.Paging, role int) ([]*audit.Audit, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var po []*audit.Audit
	db := i.mdb.Model(&audit.Audit{})
	if req.Search != "" {
		db = db.Where("reason LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	if role == 1 {
		db = db.Where("enterprise_id != 0")
	}
	if role == 2 {
		db = db.Where("job_id != 0")
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
	if role == 1 {
		var enterpriseIDs []int
		for _, item := range po {
			enterpriseIDs = append(enterpriseIDs, item.EnterpriseID)
		}
		var enterprise []*user.Enterprise
		err = i.mdb.Model(&user.Enterprise{}).
			Where("id in (?)", enterpriseIDs).
			Find(&enterprise).Error
		if err != nil {
			return nil, 0, err
		}
		enterpriseMap := make(map[int]*user.Enterprise)
		for _, item := range enterprise {
			enterpriseMap[item.ID] = item
		}
		for j, item := range po {
			po[j].EnterpriseInfo = enterpriseMap[item.EnterpriseID]
		}
		return po, total, nil
	} else if role == 2 {
		var job []*audit.Job
		var jobIDs []int
		for _, item := range po {
			jobIDs = append(jobIDs, item.EnterpriseID)
		}
		err = i.mdb.Model(&audit.Job{}).
			Where("id in (?)", jobIDs).
			Find(&job).Error
		if err != nil {
			return nil, 0, err
		}
		jobMap := make(map[int]*audit.Job)
		for _, item := range job {
			jobMap[item.ID] = item
		}
		for j, item := range po {
			po[j].JobInfo = jobMap[item.JobID]
		}
		return po, total, nil
	}

	return po, total, nil
}
func (i *impl) UpdateAudit(ctx *gin.Context, req *audit.Audit) error {
	// 开启事务
	tx := i.mdb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 第一个更新操作
	err := tx.Model(&audit.Audit{}).Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"reply":      req.Reply,
			"state":      req.State,
			"updated_at": time.Now().UnixMilli(),
			"auditor":    utils.GetUserID(ctx),
		}).Error
	if err != nil {
		tx.Rollback() // 如果出错，回滚事务
		return err
	}
	// 第二个更新操作
	err = tx.Model(&position.Job{}).Where("id = ?", req.JobID).Updates(map[string]interface{}{
		"reply": req.Reply,
		"state": -2,
	}).Error
	if err != nil {
		tx.Rollback() // 如果出错，回滚事务
		return err
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
func (i *impl) CreateAudit(ctx *gin.Context, req *audit.Audit) (*audit.Audit, error) {
	req.CreatedAt = time.Now().UnixMilli()
	err := i.mdb.Model(&audit.Audit{}).Create(&req).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *impl) DeleteAudit(ctx *gin.Context, id int) error {
	err := i.mdb.Delete(&audit.Audit{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
