package db

import (
	"context"
	"gitee.com/xygfm/authorization/conf"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Redis struct {
	client *redis.Client
	lock   sync.Mutex
}

func NewRedis() *Redis {
	return &Redis{}
}

func (r *Redis) Client() *redis.Client {
	r.lock.Lock()
	defer r.lock.Unlock()
	redisConf := conf.GetConfig().Redis
	r.client = redis.NewClient(&redis.Options{
		Addr:     redisConf.Address(),
		DB:       redisConf.Database,
		Password: redisConf.Password,
	})
	ctx := context.Background()
	err := r.client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	return r.client
}
