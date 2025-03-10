package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
	"stock/biz/dal/mysql"
	"stock/biz/dal/redis"
	"stock/biz/model"
	"stock/conf"
	"stock/rpc"
	"sync"
	"time"
)

type ReserveItemService struct {
	ctx context.Context
} // NewReserveItemService new ReserveItemService
func NewReserveItemService(ctx context.Context) *ReserveItemService {
	return &ReserveItemService{ctx: ctx}
}

// Run create note info
func (s *ReserveItemService) Run(req *stock.ReserveItemReq) (resp *stock.ReserveItemResp, err error) {
	// Finish your business logic.

	items, err := rpc.OrderClient.ShowOrderDetail(s.ctx, &order.ShowOrderDetailReq{OrderId: req.OrderId})
	if err != nil {
		return &stock.ReserveItemResp{Success: false}, err
	}
	tx := mysql.DB.Begin()
	//先存储加入到redis中的预扣商品的键值
	var preDelStockKeys []string
	for _, item := range items.OrderItems {
		// 1. 幂等性检查: 查询Redis中是否有该订单的库存预扣信息
		productId := item.GetProductId()
		//上本地锁
		mtx := getProductMutex(productId)
		mtx.Lock()
		defer mtx.Unlock()
		key := fmt.Sprintf("predestock:%s:%d", req.GetOrderId(), productId)
		exists, err := redis.RedisClient.Exists(s.ctx, key).Result()
		if err != nil {
			klog.Error(err)
			return nil, err
		}
		if exists == 1 {
			klog.Error("库存已预扣")
			return nil, errors.New("库存已预扣")
		}
		//查询之前加分布式锁
		// 1. 生成分布式锁的 Key 和随机值（防止误删）
		lockKey := fmt.Sprintf("stock_lock:%d", productId)
		lockValue := fmt.Sprintf("%d:%s", productId, time.Now().String()) // 随机值
		lockTimeout := 10 * time.Second                                   // 锁超时时间

		// 2. 尝试获取分布式锁（使用 SETNX + EXPIRE）
		locked, err := redis.RedisClient.SetNX(s.ctx, lockKey, lockValue, lockTimeout).Result()
		if err != nil {
			klog.Error("获取分布式锁失败:", err)
			tx.Rollback()
			return nil, errors.New("系统繁忙，请重试")
		}
		if !locked {
			klog.Error("其他实例正在操作库存")
			tx.Rollback()
			return nil, errors.New("操作冲突，请稍后重试")
		}

		// 3. 确保最终释放锁（使用 defer + Lua 脚本保证原子性）
		defer func() {
			script := `
			if redis.call("get", KEYS[1]) == ARGV[1] then
				return redis.call("del", KEYS[1])
			else
				return 0
			end
			`
			redis.RedisClient.Eval(s.ctx, script, []string{lockKey}, lockValue).Result()
		}()
		// 2. 查询库存是否充足
		quantity, err := model.CheckQuantity(tx, productId)
		if err != nil {
			klog.Error(err)
			tx.Rollback()
			return nil, err
		}
		if quantity < int64(req.Quantity) {
			klog.Error(err)
			tx.Rollback()
			return nil, err
		}

		// 3. 预扣库存，首先数据库中扣减库存 TODO: SQL 待修改
		if err = model.ReduceItem(tx, productId, int64(req.Quantity)); err != nil {
			klog.Error(err)
			tx.Rollback()
			return nil, err
		}
		//4. 扣减的放到Redis里
		if conf.GetEnv() == "test" {
			if err = redis.RedisClient.Set(s.ctx, key, req.Quantity, 15*time.Minute).Err(); err != nil {
				// TODO: 补偿机制
				redis.RedisClient.Del(s.ctx, key)
				klog.Error("Redis 写入失败", err)
				return nil, err
			} else {
				preDelStockKeys = append(preDelStockKeys, lockKey)
			}
		} else if conf.GetEnv() == "dev" {
			if err = redis.RedisClusterClient.Set(s.ctx, key, req.Quantity, 15*time.Minute).Err(); err != nil {
				// TODO: 补偿机制
				redis.RedisClient.Del(s.ctx, key)
				klog.Error("Redis 写入失败", err)
				return nil, err
			} else {
				preDelStockKeys = append(preDelStockKeys, lockKey)
			}
		}
	}
	if err = tx.Commit().Error; err != nil {
		// 并行删除所有关联的Redis键
		var wg sync.WaitGroup
		for _, key := range preDelStockKeys {
			wg.Add(1)
			go func(k string) {
				defer wg.Done()
				if conf.GetEnv() == "test" {
					if err := redis.RedisClient.Del(s.ctx, k).Err(); err != nil {
						klog.Errorf("Redis键删除失败 key:%s error:%v", k, err)
					}
				} else {
					if err := redis.RedisClusterClient.Del(s.ctx, k).Err(); err != nil {
						klog.Errorf("Redis键删除失败 key:%s error:%v", k, err)
					}
				}
			}(key)
		}
		wg.Wait()
		klog.Error(err)
		return nil, err
	}

	return &stock.ReserveItemResp{
		Success: true,
	}, nil
}
