package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/place"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListSchool(ctx *gin.Context, req *response.Paging, area int) ([]*place.School, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var school []*place.School
	db := i.mdb.Model(&place.School{})
	if req.Search != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	if area != 0 {
		db = db.Where("province_id = ? OR city_id = ? OR area_id = ?", area, area, area)
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&school).Error
	if err != nil {
		return nil, 0, err
	}
	var areaIDs []int
	var areas []*place.Area
	areaMap := make(map[int]string)
	for _, item := range school {
		areaIDs = append(areaIDs, item.AreaID)
		areaIDs = append(areaIDs, item.ProvinceID)
		areaIDs = append(areaIDs, item.CityID)
	}
	fmt.Println("areaIDs:", areaIDs)
	err = i.mdb.Model(&place.Area{}).Select("id", "name").Where("id IN (?)", areaIDs).Find(&areas).Error
	if err != nil {
		return nil, 0, err
	}
	for _, item := range areas {
		areaMap[item.ID] = item.Name
	}
	for j, item := range school {
		school[j].Province = areaMap[item.ProvinceID]
		school[j].City = areaMap[item.CityID]
		school[j].Area = areaMap[item.AreaID]
	}
	return school, total, nil
}
func (i *impl) UpdateSchool(ctx *gin.Context, req place.School) error {
	err := i.mdb.Model(&place.School{}).Where("id = ?", req.ID).Updates(&req).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) CreateSchool(ctx *gin.Context, req place.School) (*place.School, error) {
	err := i.mdb.Model(&place.School{}).Create(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, nil
}
func (i *impl) DeleteSchool(ctx *gin.Context, id int) error {
	err := i.mdb.Delete(&place.School{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) ListHotSchool(ctx *gin.Context, req *response.Paging) ([]*place.School, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var school []*place.School
	db := i.mdb.Model(&place.School{})
	if req.Search != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err := db.Order("visit_num desc").Find(&school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}
