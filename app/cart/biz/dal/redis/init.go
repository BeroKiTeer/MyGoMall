package redis

import (
	"context"

	"cart/conf"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient // 修改为集群客户端类型
)

func Init() {
	if conf.GetEnv() == "test" {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     conf.GetConf().Redis.Address,
			Username: conf.GetConf().Redis.Username,
			Password: conf.GetConf().Redis.Password,
			DB:       conf.GetConf().Redis.DB,
		})
	} else if conf.GetEnv() == "dev" {
		// 创建集群客户端
		RedisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    conf.GetConf().Redis.Addresses, // 需要改为复数形式，支持多个地址
			Username: conf.GetConf().Redis.Username,
			Password: conf.GetConf().Redis.Password,
			// 注意：集群模式通常不使用DB参数（Redis集群只支持DB 0）

			// 可根据需要添加集群专用配置
			MaxRedirects:   8,     // 最大重试次数
			ReadOnly:       false, // 是否开启只读模式
			RouteByLatency: false, // 是否开启就近路由
		})
	}
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
