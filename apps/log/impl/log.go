package impl

import (
	"gitee.com/xygfm/authorization/apps/log"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListLog(ctx *gin.Context, req *response.Paging) ([]*log.Log, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*log.Log
	db := i.mdb.Model(&log.Log{})
	if req.Search != "" {
		db = db.Where("LOWER(path) LIKE LOWER(?) OR LOWER(method) LIKE LOWER(?) OR LOWER(query) LIKE LOWER(?) OR ip LIKE ?",
			"%"+req.Search+"%",
			"%"+req.Search+"%",
			"%"+req.Search+"%",
			"%"+req.Search+"%")
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

func (i *impl) DeleteLog(ctx *gin.Context, id int) error {
	err := i.mdb.Delete(&log.Log{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
