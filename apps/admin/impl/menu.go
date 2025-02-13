package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListMenu(ctx *gin.Context, req []int) ([]*admin.Menu, error) {
	// 查询所有菜单数据
	var menus []*admin.Menu
	err := i.mdb.Model(&admin.Menu{}).Find(&menus).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// 使用 Map 缓存所有菜单
	menuMap := make(map[int]*admin.Menu)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 构建树形结构
	var tree []*admin.Menu
	for _, menu := range menus {
		if menu.ParentID == 0 {
			// 根节点
			tree = append(tree, menu)
		} else {
			// 找到父节点，将当前菜单添加到父节点的 Children 中
			if parent, ok := menuMap[menu.ParentID]; ok {
				if parent.Children == nil {
					parent.Children = []*admin.Menu{}
				}
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return tree, nil
}

func (i *impl) UpdateMenu(ctx *gin.Context, req *admin.Menu) error {
	err := i.mdb.Model(&admin.Menu{}).
		Where("id = ?", req.ID).
		Updates(&req).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) CreateMenu(ctx *gin.Context, req *admin.Menu) (*admin.Menu, error) {
	err := i.mdb.Create(&req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}
func (i *impl) DeleteMenu(ctx *gin.Context, id int) error {
	if err := i.mdb.Delete(&admin.Menu{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
