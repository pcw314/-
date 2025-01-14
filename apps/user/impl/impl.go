package impl

import (
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/db"
	"gitee.com/xygfm/authorization/ioc"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type impl struct {
	ioc.ObjectImpl
	user.ServiceImpl
	mdb *gorm.DB
	rdb *redis.Client
}

func init() {
	// 创建 impl 的实例
	implInstance := &impl{}

	// 将实例注册到 IOC 容器中
	err := ioc.RegisterController(implInstance)
	if err != nil {
		panic(err) // 如果注册失败，抛出 panic
	}
}

func (i *impl) Init() error {
	i.mdb = db.NewMysql().Gorm()
	i.rdb = db.NewRedis().Client()
	return nil
}

func (i *impl) Name() string {
	return user.AppName
}
