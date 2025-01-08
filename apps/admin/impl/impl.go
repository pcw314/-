package impl

import (
	"gitee.com/xygfm/authorization/apps/admin"
	"gitee.com/xygfm/authorization/db"
	"gitee.com/xygfm/authorization/ioc"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type impl struct {
	ioc.ObjectImpl
	admin.ServiceImpl
	mdb *gorm.DB
	rdb *redis.Client
}

func init() {
	err := ioc.RegisterController(&impl{})
	if err != nil {
		panic(err)
	}
}

func (i *impl) Init() error {
	i.mdb = db.NewMysql().Gorm()
	i.rdb = db.NewRedis().Client()
	return nil
}

func (i *impl) Name() string {
	return admin.AppName
}
