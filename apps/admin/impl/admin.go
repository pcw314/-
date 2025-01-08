package impl

import (
	"context"
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	service "gitee.com/xygfm/authorization/middleware"
)

func (i *impl) Login(ctx context.Context, req *admin.LoginRequest) (string, error) {
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
