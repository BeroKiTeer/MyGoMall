package redis

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

var (
	ErrLocked  = errors.New("分布式锁竞争失败")
	ErrTimeOut = errors.New("超时！")
)

type Lock struct {
	lockKey    string
	lockValue  string
	lockttl    time.Duration
	watchDog   chan struct{}      // 看门狗通知通道
	stopCancel context.CancelFunc // 控制续期协程
}

func NewLock(key, value string, ttl time.Duration) *Lock {
	return &Lock{
		lockKey:   key,
		lockValue: value,
		lockttl:   ttl,
		watchDog:  make(chan struct{}),
	}
}

func (l *Lock) TryLock(ctx context.Context) error {

	// 2. 尝试获取分布式锁（使用 SETNX + EXPIRE）
	locked, err := RedisClusterClient.SetNX(ctx, l.lockKey, l.lockValue, l.lockttl).Result()
	if err != nil {
		klog.Error("获取分布式锁失败:", err)
		return err
	}
	if !locked {
		klog.Error("其他实例正在操作库存")
		return ErrLocked
	}
	go l.StartWatchDog() //启动看门狗
	return nil
}

func (l *Lock) UnLock(ctx context.Context) error {
	script := `
			if redis.call("get", KEYS[1]) == ARGV[1] then
				return redis.call("del", KEYS[1])
			else
				return 0
			end
			`
	err := RedisClusterClient.Eval(ctx, script, []string{l.lockKey}, l.lockValue).Err()
	close(l.watchDog)
	return err
}
func (l *Lock) StartWatchDog() {
	ticker := time.NewTicker(l.lockttl / 3)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 延长锁的过期时间
			ctx, cancel := context.WithTimeout(context.Background(), l.lockttl/3*2)
			ok, err := RedisClusterClient.Expire(ctx, l.lockKey, l.lockttl).Result()
			cancel()
			// 异常或锁已经不存在则不再续期
			if err != nil || !ok {
				klog.Warn("锁续期失败，可能已释放")
				return
			}
		case <-l.watchDog:
			// 已经解锁
			klog.Info("接收到锁释放信号，停止续期")
			return

		}
	}
}
