package rds

import (
	"github.com/go-redis/redis"
	"go-iris-frame-demo/etc/config"
)

var Demo *redis.Client

func init() {
	Demo = redis.NewClient(&redis.Options{
		Addr:     config.Val("redis.addr"),
		Password: config.Val("redis.pwd"),
		PoolSize: config.Int("redis.pool"),
		DB:       config.Int("redis.db"),
	})

	_, err := Demo.Ping().Result()
	if err != nil {
		panic(err)
	}
}
