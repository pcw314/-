package utils

import (
	"errors"
	"gitee.com/xygfm/authorization/conf"
	"gitee.com/xygfm/authorization/global"
	"gorm.io/gorm"
)

func GetApiByType(Type string) *conf.Api {
	api := &conf.Api{}
	err := global.Mdb.Model(&conf.Api{}).Where("type = ?", Type).First(api).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
	return api
}
