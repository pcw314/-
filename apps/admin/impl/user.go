package impl

import (
	"context"
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	"gitee.com/xygfm/authorization/global"
)

func (i *impl) Register(ctx context.Context, req *admin.LoginRequest) (string, error) {
	fmt.Println("我是Register业务层")
	var name admin.LoginRequest
	err := global.Mdb.Model(&admin.LoginRequest{}).Find(&name).Error
	err = i.mdb.Model(&admin.LoginRequest{}).Find(&name).Error
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	fmt.Println("name")
	return "", nil
}
