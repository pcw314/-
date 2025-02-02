package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
)

func (i *impl) GetUser(ctx *gin.Context, req *admin.LoginRequest) (*admin.User, error) {
	fmt.Println("我是登录业务层")
	var user admin.User
	err := i.mdb.Model(&admin.User{}).
		Where("username", req.Username).
		Where("role", req.Role).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (i *impl) Register(ctx *gin.Context, req *admin.LoginRequest) (string, error) {
	err := i.mdb.Model(&admin.LoginRequest{}).Create(&admin.LoginRequest{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}).Error
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	return "", nil
}

func (i *impl) ListArea(ctx *gin.Context, id string) ([]*admin.Areas, error) {
	var po []*admin.Areas
	db := i.mdb.Model(&admin.Areas{})
	if id != "0" {
		db.Where("parent_id = ?", id)
	}
	if utils.GetUserID(ctx) != 1 {
		db.Where("state = ?", 1)
	}
	err := db.Find(&po).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("po", po)
	// 使用 Map 缓存所有菜单
	areaMap := make(map[int]*admin.Areas)
	for _, item := range po {
		areaMap[item.ID] = item
	}

	// 构建树形结构
	var tree []*admin.Areas
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
