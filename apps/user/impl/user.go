package impl

import (
	"context"
	"fmt"
	"gitee.com/xygfm/authorization/apps/user"
	service "gitee.com/xygfm/authorization/middleware"
)

func (i *impl) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
	fmt.Println("我是登录业务层")
	//var name admin.LoginRequest
	//err := global.Mdb.Model(&admin.LoginRequest{}).Find(&name).Error
	//err = i.mdb.Model(&admin.LoginRequest{}).Find(&name).Error
	token, err := service.MakeToken(2, "13979875415", 2)
	if err != nil {
		return "", err
	}
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	fmt.Println("name")
	return token, nil
}

func (i *impl) Register(ctx context.Context, req *user.LoginRequest) (string, error) {
	fmt.Println("我是Register业务层")
	var name user.LoginRequest
	err := i.mdb.Model(&user.LoginRequest{}).Find(&name).Error
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	fmt.Println("name")
	return "", nil
}
func (i *impl) ListMenu(ctx context.Context, req *[]int) ([]*user.MenuRequest, error) {
	fmt.Println("进入获取菜单的业务层")
	// 查询所有菜单数据
	var menus []*user.MenuRequest
	err := i.mdb.Model(&user.MenuRequest{}).Find(&menus).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// 使用 Map 缓存所有菜单
	menuMap := make(map[int]*user.MenuRequest)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 构建树形结构
	var tree []*user.MenuRequest
	for _, menu := range menus {
		if menu.PID == 0 {
			// 根节点
			tree = append(tree, menu)
		} else {
			// 找到父节点，将当前菜单添加到父节点的 Children 中
			if parent, ok := menuMap[menu.PID]; ok {
				if parent.Children == nil {
					parent.Children = []*user.MenuRequest{}
				}
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return tree, nil
}
