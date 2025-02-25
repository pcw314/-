package impl

import (
	"errors"
	"fmt"
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"time"
)

func (i *impl) ListJob(ctx *gin.Context, req *response.Paging, enterpriseID int) ([]*position.Job, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*position.Job
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
	err = db.Order("id desc").Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	var enterpriseIDs []int
	var enterprise []*user.Enterprise

	for _, item := range pos {
		enterpriseIDs = append(enterpriseIDs, item.EnterpriseID)
	}
	err = i.mdb.Model(&user.Enterprise{}).
		Where("id IN (?)", enterpriseIDs).
		Find(&enterprise).Error
	if err != nil {
		return nil, 0, err
	}

	enterpriseMAP := make(map[int]*user.Enterprise)
	for _, item := range enterprise {
		enterpriseMAP[item.ID] = item
	}

	for j, item := range pos {
		if enterpriseInfo, ok := enterpriseMAP[item.EnterpriseID]; ok {
			pos[j].EnterpriseInfo = enterpriseInfo
		}
	}

	return pos, total, nil
}
func (i *impl) UpdateJob(ctx *gin.Context, req *position.Job) error {
	i.mdb.Model(&position.Job{}).Where("id = ?", req.ID).
		Updates(map[string]interface{}{
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
	req.State = -1
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
func (i *impl) GetJobByID(ctx *gin.Context, id int) (*position.Job, error) {
	var po *position.Job
	err := i.mdb.Model(&position.Job{}).Where("id = ?", id).First(&po).Error
	if err != nil {
		return nil, err
	}
	//var enterpriseInfo *user.Enterprise
	//err = i.mdb.Model(&user.Enterprise{}).
	//	Where("id = ?", po.EnterpriseID).
	//	First(&enterpriseInfo).Error
	//if err != nil {
	//	return nil, err
	//}
	//po.EnterpriseInfo = enterpriseInfo
	return po, nil
}

func (i *impl) SetJobState(ctx *gin.Context, id int) (string, error) {
	var po *position.Job
	err := i.mdb.Model(&position.Job{}).Where("id = ?", id).First(&po).Error
	if err != nil {
		return "", err
	}
	if po.State == 1 {
		err = i.mdb.Model(&position.Job{}).Where("id = ?", id).Update("state", -1).Error
		if err != nil {
			return "", err
		}
		return "下架成功", nil
	} else if po.State == -1 {
		err = i.mdb.Model(&position.Job{}).Where("id = ?", id).Update("state", 0).Error
		if err != nil {
			return "", err
		}
		return "上架成功", nil
	} else if po.State == 0 {
		return "当前处于待审核", errors.New("等待审核")
	}
	return "", nil
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
	if req.Search != "" {
		db = db.Where("name LIKE ? OR post LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
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
	var jobIDs []int
	var enterpriseIDs []int
	var enterprise []*user.Enterprise
	//var userIDs []int
	//var users []*user.User
	for _, item := range pos {
		enterpriseIDs = append(enterpriseIDs, item.EnterpriseID)
		jobIDs = append(jobIDs, item.ID)
	}

	var collects []*position.Collect
	err = i.mdb.Model(&position.Collect{}).
		Where("user_id = ?", utils.GetUserID(ctx)).
		Where("job_id in ?", jobIDs).
		Find(&collects).Error

	collectMAP := make(map[int]int)
	for _, item := range collects {
		collectMAP[item.JobID] = 1
	}

	err = i.mdb.Model(&user.Enterprise{}).
		Where("id IN (?)", enterpriseIDs).
		Find(&enterprise).Error
	if err != nil {
		return nil, 0, err
	}
	//for _, item := range enterprise {
	//	userIDs = append(userIDs, item.UserID)
	//}
	//err = i.mdb.Model(&user.User{}).
	//	Where("id IN (?)", userIDs).
	//	Find(&users).Error
	//if err != nil {
	//	return nil, 0, err
	//}
	enterpriseMAP := make(map[int]*user.Enterprise)
	//userMAP := make(map[int]*user.User)
	for _, item := range enterprise {
		enterpriseMAP[item.ID] = item
	}
	//for _, item := range users {
	//	userMAP[item.ID] = item
	//}
	//for j, item := range pos {
	//	pos[j].EnterpriseInfo = enterpriseMAP[item.EnterpriseID]
	//	pos[j].UserInfo = userMAP[enterpriseMAP[item.EnterpriseID].UserID]
	//}
	for j, item := range pos {
		if enterpriseInfo, ok := enterpriseMAP[item.EnterpriseID]; ok {
			pos[j].EnterpriseInfo = enterpriseInfo
			//if userInfo, ok := userMAP[enterpriseInfo.UserID]; ok {
			//	pos[j].UserInfo = userInfo
			//}
		}
		if is, ok := collectMAP[item.ID]; ok {
			pos[j].IsCollect = is
		}
	}

	return pos, total, nil
}
