package impl

import (
	"errors"
	"fmt"
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i *impl) ListCollect(ctx *gin.Context, req *response.Paging, userID int) ([]*position.Job, int64, error) {
	i.mdb.Model(&position.Collect{}).Where("user_id = ?", userID)
	limit := req.Size
	offset := (req.Page - 1) * limit
	var jobID []int
	var pos []*position.Job
	db := i.mdb.Model(&position.Collect{}).Select("job_id").
		Where("user_id = ?", userID)
	if req.Search != "" {
		db = db.Where("name LIKE ?", "%"+req.Search+"%")
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&jobID).Error
	if err != nil {
		return nil, 0, err
	}

	err = i.mdb.Model(&position.Job{}).Where("id in (?)", jobID).Find(&pos).Error
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
	//for j, item := range pos {
	//	pos[j].EnterpriseInfo = enterpriseMAP[item.EnterpriseID]
	//	pos[j].UserInfo = userMAP[enterpriseMAP[item.EnterpriseID].UserID]
	//}
	for j, item := range pos {
		if enterpriseInfo, ok := enterpriseMAP[item.EnterpriseID]; ok {
			pos[j].EnterpriseInfo = enterpriseInfo
			if userInfo, ok := userMAP[enterpriseInfo.UserID]; ok {
				pos[j].UserInfo = userInfo
			}
		}
	}
	return pos, total, nil
}

//	func (i *impl) CreateCollect(ctx *gin.Context, collect position.Collect) (*position.Collect, error) {
//		return nil, nil
//	}
//
//	func (i *impl) DeleteCollect(ctx *gin.Context, id int) error {
//		return nil
//	}
func (i *impl) CreateCollect(ctx *gin.Context, collect position.Collect) (*position.Collect, error) {
	// 检查是否已经存在相同的收藏记录
	var existingCollect position.Collect
	if err := i.mdb.Model(&position.Collect{}).
		Where("user_id = ? AND job_id = ?", collect.UserID, collect.JobID).
		First(&existingCollect).Error; err == nil {
		// 如果记录已存在，那就删除该收藏
		err = i.DeleteCollect(ctx, existingCollect.ID)
		if err != nil {
			return nil, err
		}
		return &existingCollect, fmt.Errorf("collect already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果发生其他数据库错误，返回错误
		return nil, err
	}

	// 创建新的收藏记录
	if err := i.mdb.Create(&collect).Error; err != nil {
		return nil, err
	}

	// 返回创建成功的收藏记录
	return &collect, nil
}

func (i *impl) DeleteCollect(ctx *gin.Context, id int) error {
	// 查询并删除指定 ID 的收藏记录
	if err := i.mdb.Model(&position.Collect{}).Where("id = ?", id).Delete(&position.Collect{}).Error; err != nil {
		return err
	}

	// 如果删除成功，返回 nil
	return nil
}
