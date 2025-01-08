package global

import (
	"gitee.com/xygfm/authorization/db"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"sync"
)

var (
	Mdb   *gorm.DB
	mLock sync.Mutex
	rdb   *redis.Client
	rLock sync.Mutex
)

func Init() {
	Mdb = db.NewMysql().Gorm()
	rdb = db.NewRedis().Client()
}

func GetMysql() *gorm.DB {
	mLock.Lock()
	defer mLock.Unlock()
	return Mdb
}

func GetRedis() *redis.Client {
	rLock.Lock()
	defer rLock.Unlock()
	return rdb
}
