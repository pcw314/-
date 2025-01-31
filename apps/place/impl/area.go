package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/place"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListArea(ctx *gin.Context, id int) ([]*place.Area, error) {
	var po []*place.Area
	db := i.mdb.Model(&place.Area{})
	if id != 0 {
		db.Where("parent_id = ?", id)
	}
	if utils.GetUserRole(ctx) != 3 {
		db.Where("state = ?", 1)
	}
	err := db.Find(&po).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("po", po)
	// 使用 Map 缓存所有
	areaMap := make(map[int]*place.Area)
	for _, item := range po {
		areaMap[item.ID] = item
	}

	// 构建树形结构
	var tree []*place.Area
	for _, item := range po {
		if item.ParentID == 1 {
			// 根节点
			tree = append(tree, item)
		} else {
			// 尝试找父节点
			parent, ok := areaMap[item.ParentID]
			if ok {
				// 找到父节点，添加为子节点
				parent.Children = append(parent.Children, item)
			} else {
				// 找不到父节点，作为顶层节点处理
				tree = append(tree, item)
			}
		}
	}

	return tree, nil
}

func (i *impl) UpdateArea(ctx *gin.Context, req place.Area) error {
	err := i.mdb.Model(&place.Area{}).Where("id = ?", req.ID).Updates(&req).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) CreateArea(ctx *gin.Context, req place.Area) (*place.Area, error) {
	err := i.mdb.Model(&place.Area{}).Create(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, nil
}
func (i *impl) DeleteArea(ctx *gin.Context, id int) error {
	err := i.mdb.Where("id = ?", id).Delete(&place.Area{}).Error
	if err != nil {
		return err
	}
	return nil
}
