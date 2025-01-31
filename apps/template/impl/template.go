package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/template"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListName(ctx *gin.Context, req *response.Paging) ([]*template.Template, int64, error) {
	fmt.Println("业务层")
	limit := req.Size
	offset := (req.Page - 1) * limit

	var name []*template.Template
	db := i.mdb.Model(&template.Template{})
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
	err = db.Order("id desc").Find(&name).Error
	if err != nil {
		return nil, 0, err
	}

	return name, total, nil
}
