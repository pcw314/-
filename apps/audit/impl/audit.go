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

func (i *impl) ListAudit(ctx *gin.Context, req *response.Paging, userID int) ([]*audit.Audit, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*audit.Audit
	db := i.mdb.Model(&audit.Audit{}).Where("informer = ?", userID)
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
	err = db.Order("id desc").Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	return pos, total, nil
}
func (i *impl) ListAuditUser(ctx *gin.Context, req *response.Paging) ([]*audit.Audit, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*audit.Audit
	db := i.mdb.Model(&audit.Audit{}).Where("user_id != 0")
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
	err = db.Order("id desc").Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	var userIds []int
	for _, item := range pos {
		userIds = append(userIds, item.Informer)
		userIds = append(userIds, item.UserID)
	}
	var enterprise []*user.Enterprise
	err = i.mdb.Model(&user.Enterprise{}).
		Where("user_id in (?)", userIds).
		Find(&enterprise).Error
	if err != nil {
		return nil, 0, err
	}
	userInfoMap := make(map[int]*audit.UserInfo)
	for _, item := range enterprise {
		userInfoMap[item.UserID] = &audit.UserInfo{
			Name:   item.Name,
			Phone:  item.Phone,
			Avatar: item.Avatar,
		}
	}
	var student []*user.Student
	err = i.mdb.Model(&user.Student{}).
		Where("user_id in (?)", userIds).
		Find(&student).Error
	if err != nil {
		return nil, 0, err
	}
	for _, item := range student {
		userInfoMap[item.UserID] = &audit.UserInfo{
			Name:   item.Name,
			Phone:  item.Phone,
			Avatar: item.Avatar,
		}
	}
	if len(userInfoMap) == 0 {
		return pos, total, nil
	}
	for j, item := range pos {
		pos[j].InformerName = userInfoMap[item.Informer].Name
		pos[j].Name = userInfoMap[item.UserID].Name
		pos[j].Phone = userInfoMap[item.UserID].Phone
		pos[j].Avatar = userInfoMap[item.UserID].Avatar
	}
	return pos, total, nil
}
func (i *impl) ListAuditJob(ctx *gin.Context, req *response.Paging) ([]*audit.Audit, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*audit.Audit
	db := i.mdb.Model(&audit.Audit{}).Where("job_id != 0")
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
	err = db.Order("id desc").Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	//var userIds []int
	//for _, item := range po {
	//	userIds = append(userIds, item.UserID)
	//}
	//var ids []int //Enterprise 的id
	//err = i.mdb.Model(&user.Enterprise{}).Select("id").
	//	Where("user_id in (?)", userIds).
	//	Find(&ids).Error
	//if err != nil {
	//	return nil, 0, err
	//}

	var job []*audit.Job
	var jobIDs []int
	for _, item := range pos {
		jobIDs = append(jobIDs, item.JobID)
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
	var userIds []int
	for _, item := range pos {
		userIds = append(userIds, item.Informer)
	}
	var student []*user.Student
	err = i.mdb.Model(&user.Student{}).
		Where("user_id in (?)", userIds).
		Find(&student).Error
	if err != nil {
		return nil, 0, err
	}
	userInfoMap := make(map[int]*audit.UserInfo)
	for _, item := range student {
		userInfoMap[item.UserID] = &audit.UserInfo{
			Name:   item.Name,
			Phone:  item.Phone,
			Avatar: item.Avatar,
		}
	}

	for j, item := range pos {
		pos[j].JobInfo = jobMap[item.JobID]
		pos[j].InformerName = userInfoMap[item.Informer].Name
		pos[j].Avatar = userInfoMap[item.Informer].Avatar
	}
	return pos, total, nil
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
	req.UpdatedAt = time.Now().UnixMilli()
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

func (i *impl) ListAuditJobInfo(ctx *gin.Context, req *response.Paging) ([]*position.Job, int64, error) {
	db := i.mdb.Model(&position.Job{})
	if req.Search != "" {
		db = db.Where("reason LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	return nil, 0, nil
}
