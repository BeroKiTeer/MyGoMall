package redis

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/redis/go-redis/v9"
	"product/conf"
)

var (
	RedisClient *redis.Client
	Nil         redis.Error = redis.Nil
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		klog.Fatal("redis连接失败: ", err)
		panic(err)
	}
}
