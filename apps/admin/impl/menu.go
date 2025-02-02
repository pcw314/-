package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListMenu(ctx *gin.Context, req []int) ([]*admin.MenuRequest, error) {
	// 查询所有菜单数据
	var menus []*admin.MenuRequest
	err := i.mdb.Model(&admin.MenuRequest{}).Where("id in ?", req).Find(&menus).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// 使用 Map 缓存所有菜单
	menuMap := make(map[int]*admin.MenuRequest)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 构建树形结构
	var tree []*admin.MenuRequest
	for _, menu := range menus {
		if menu.PID == 0 {
			// 根节点
			tree = append(tree, menu)
		} else {
			// 找到父节点，将当前菜单添加到父节点的 Children 中
			if parent, ok := menuMap[menu.PID]; ok {
				if parent.Children == nil {
					parent.Children = []*admin.MenuRequest{}
				}
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return tree, nil
}

func (i *impl) UpdateMenu(ctx *gin.Context, req *admin.MenuRequest) error {
	err := i.mdb.Model(&admin.MenuRequest{}).
		Where("id = ?", req.ID).
		Updates(&req).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) CreateMenu(ctx *gin.Context, req *admin.MenuRequest) (*admin.MenuRequest, error) {
	err := i.mdb.Create(&req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}
func (i *impl) DeleteMenu(ctx *gin.Context, id int) error {
	if err := i.mdb.Delete(&admin.MenuRequest{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
