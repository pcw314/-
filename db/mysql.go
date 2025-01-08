package db

import (
	"fmt"
	"gitee.com/xygfm/authorization/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Mysql struct {
	db   *gorm.DB
	lock sync.Mutex
}

func NewMysql() *Mysql {
	return &Mysql{}
}

func (m *Mysql) Gorm() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	c := conf.GetConfig().Mysql
	if m.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&multiStatements=true", c.Username, c.Password, c.Host, c.Port, c.Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db
	}
	return m.db
}
