package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
	"time"
)

func (i *impl) ListJob(ctx *gin.Context, req *response.Paging, enterpriseID int) ([]*position.Job, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var po []*position.Job
	db := i.mdb.Model(&position.Job{}).
		Where("enterprise_id = ?", enterpriseID)
	if req.Search != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
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
func (i *impl) UpdateJob(ctx *gin.Context, req *position.Job) error {
	i.mdb.Model(&position.Job{}).Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"name":           req.Name,
			"type":           req.Type,
			"working_time":   req.WorkingTime,
			"place":          req.Place,
			"salary":         req.Salary,
			"post":           req.Post,
			"school_id":      req.SchoolID,
			"description":    req.Description,
			"require":        req.Require,
			"contact_name":   req.ContactName,
			"contact_number": req.ContactNumber,
			"state":          req.State,
			"updated_at":     time.Now().UnixMilli(),
		})
	return nil
}
func (i *impl) CreateJob(ctx *gin.Context, req *position.Job) (*position.Job, error) {
	req.CreatedAt = time.Now().UnixMilli()
	err := i.mdb.Model(&position.Job{}).Create(&req).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (i *impl) DeleteJob(ctx *gin.Context, id int) error {
	err := i.mdb.Delete(&position.Job{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) GetEnterpriseByUserID(ctx *gin.Context, userID int) (*user.Enterprise, error) {
	var po *user.Enterprise
	err := i.mdb.Model(&user.Enterprise{}).Where("user_id = ?", userID).First(&po).Error
	if err != nil {
		return nil, err
	}
	return po, nil
}

func (i *impl) ListJobBySchoolID(ctx *gin.Context, req *response.Paging, schoolID int) ([]*position.Job, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*position.Job
	db := i.mdb.Model(&position.Job{})
	if schoolID != 0 {
		db = db.Where("school_id = ?", schoolID)
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
	var enterpriseIDs []int
	var enterprise []*user.Enterprise
	var userIDs []int
	var users []*user.User
	for _, item := range pos {
		enterpriseIDs = append(enterpriseIDs, item.EnterpriseID)
	}
	err = i.mdb.Model(&user.Enterprise{}).
		Where("id IN (?)", enterpriseIDs).
		Find(&enterprise).Error
	if err != nil {
		return nil, 0, err
	}
	for _, item := range enterprise {
		userIDs = append(userIDs, item.UserID)
	}
	err = i.mdb.Model(&user.User{}).
		Where("id IN (?)", userIDs).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	enterpriseMAP := make(map[int]*user.Enterprise)
	userMAP := make(map[int]*user.User)
	for _, item := range enterprise {
		enterpriseMAP[item.ID] = item
	}
	for _, item := range users {
		userMAP[item.ID] = item
	}
	for j, item := range pos {
		pos[j].EnterpriseInfo = enterpriseMAP[item.EnterpriseID]
		pos[j].UserInfo = userMAP[enterpriseMAP[item.EnterpriseID].UserID]
	}

	return pos, total, nil
}
