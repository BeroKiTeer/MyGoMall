package redis

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"

	"github.com/redis/go-redis/v9"
	"stock/conf"
)

var (
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient // 修改为集群客户端类型
	Nil                redis.Error          = redis.Nil
	ErrLocked                               = errors.New("分布式锁竞争失败")
	ErrTimeOut                              = errors.New("超时！")
)

func Init() {

	if conf.GetEnv() == "test" {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     conf.GetConf().Redis.Address,
			Username: conf.GetConf().Redis.Username,
			Password: conf.GetConf().Redis.Password,
			DB:       conf.GetConf().Redis.DB,
		})
		if err := RedisClient.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}
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
		if err := RedisClusterClient.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}
	}

}

func TryLock(ctx context.Context, lockKey string, lockValue string) error {

	lockTimeout := 10 * time.Second // 锁超时时间

	// 2. 尝试获取分布式锁（使用 SETNX + EXPIRE）
	locked, err := RedisClusterClient.SetNX(ctx, lockKey, lockValue, lockTimeout).Result()
	if err != nil {
		klog.Error("获取分布式锁失败:", err)
		return err
	}
	if !locked {
		klog.Error("其他实例正在操作库存")
		return ErrLocked
	}
	return nil
}

func UnLock(ctx context.Context, Key string, Value string) error {
	script := `
			if redis.call("get", KEYS[1]) == ARGV[1] then
				return redis.call("del", KEYS[1])
			else
				return 0
			end
			`
	return RedisClusterClient.Eval(ctx, script, []string{Key}, Value).Err()
}
