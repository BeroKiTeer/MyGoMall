package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
	Redis "github.com/redis/go-redis/v9"
	"math/rand"
	"stock/biz/dal/mysql"
	"stock/biz/dal/redis"
	"stock/biz/model"
	"sync"
	"time"
)

type CheckItemService struct {
	ctx context.Context
} // NewCheckItemService new CheckItemService
func NewCheckItemService(ctx context.Context) *CheckItemService {
	return &CheckItemService{ctx: ctx}
}

// 新增本地锁 带清理机制的sync.Map
var mutexMap sync.Map

func getProductMutex(productID int64) *sync.Mutex {
	actual, _ := mutexMap.LoadOrStore(productID, &sync.Mutex{})
	return actual.(*sync.Mutex)
}

// Run create note info
func (s *CheckItemService) Run(req *stock.CheckItemReq) (resp *stock.CheckItemResp, err error) {

	productID := req.GetProductId()
	// 增加性能埋点
	startTime := time.Now()
	defer func() {
		klog.Infof("库存检查耗时 product:%d duration:%v", productID, time.Since(startTime))
	}()
	cacheKey := fmt.Sprintf("stock:%d", productID)
	// 双重检查锁: 保证线程安全的前提下 减少锁竞争开销
	mtx := getProductMutex(productID)
	mtx.Lock()
	defer mtx.Unlock()
	defer func() {
		// 异步清理30分钟未使用的锁
		go func(id int64) {
			time.Sleep(30 * time.Minute)
			mutexMap.Delete(id)
		}(productID)
	}()
	// 第一次缓存检查
	cachedProduct, err := redis.RedisClusterClient.Get(s.ctx, cacheKey).Result()
	if err == nil {
		// 如果缓存命中，直接返回缓存结果
		klog.Infof("首次缓存命中 product:%d", productID)
		var respData stock.CheckItemResp
		// 将 Redis 中的缓存数据反序列化到 respData
		err = json.Unmarshal([]byte(cachedProduct), &respData)
		if err != nil {
			klog.Errorf("failed to unmarshal cached product data: %v", err)
			return nil, fmt.Errorf("failed to unmarshal cached product data: %v", err)
		}
		return &respData, nil
	}
	// 第二次缓存检查（防止锁内重复查询）
	cachedProduct, err = redis.RedisClusterClient.Get(s.ctx, cacheKey).Result()
	if err == nil {
		if cachedProduct == "" {
			klog.Warnf("空缓存值 product:%d", productID)
			return nil, errors.New("empty cache value")
		}
		// 防止其他线程已更新缓存
		klog.Debugf("二次缓存命中 product:%d", productID)
		var respData stock.CheckItemResp
		// 将 Redis 中的缓存数据反序列化到 respData
		err = json.Unmarshal([]byte(cachedProduct), &respData)
		if err != nil {
			klog.Errorf("failed to unmarshal cached product data: %v", err)
			return nil, fmt.Errorf("failed to unmarshal cached product data: %v", err)
		}
		return &respData, nil
	}

	lockKey := fmt.Sprintf("lock:stock:%d", productID)
	lockValue := time.Now().UnixNano()
	defer func() {
		// 使用Lua脚本验证锁值后删除
		script := Redis.NewScript(`
		local val = redis.call("GET", KEYS[1])
		if val == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		end
		return 0
		`)
		script.Run(s.ctx, redis.RedisClusterClient, []string{lockKey}, lockValue)
	}()

	// 如果缓存没有命中，从数据库获取
	quantity, err := model.CheckQuantity(mysql.DB, req.ProductId)
	klog.Infof("数据库查询耗时 product:%d duration:%v",
		productID, time.Since(startTime).Round(time.Millisecond))
	if err != nil {
		klog.Errorf("Database query failed: %v", err)
		return nil, err
	}

	// 生成缓存数据
	cacheData := &stock.CheckItemResp{Quantity: quantity}
	if quantity < 0 {
		klog.Errorf("库存数量异常 productID:%d quantity:%d", productID, quantity)
		return nil, fmt.Errorf("invalid stock quantity")
	}
	jsonData, err := json.Marshal(cacheData)
	if err != nil {
		klog.Errorf("JSON序列化失败: %v", err)
		return nil, err
	}

	// 设置缓存
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = redis.RedisClusterClient.Set(s.ctx, cacheKey, jsonData,
			time.Duration(300+rand.Intn(100))*time.Second).Err()
		// 带随机抖动的缓存过期时间（300-400秒）防止缓存雪崩
		if err == nil {
			klog.Infof("缓存设置成功 product:%d duration:%v",
				productID, time.Since(startTime).Round(time.Millisecond))
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err != nil {
		klog.Errorf("缓存重试3次后仍失败 product:%d error:%v", productID, err)
	}

	// 设置锁（NX模式）
	result, err := redis.RedisClusterClient.SetNX(s.ctx,
		fmt.Sprintf("lock:stock:%d", productID),
		time.Now().UnixNano(),
		3*time.Second).Result()
	if err != nil {
		klog.Errorf("分布式锁操作异常 product:%d error:%v", productID, err)
	} else if !result {
		klog.Warnf("分布式锁竞争失败 product:%d", productID)
	}

	return cacheData, nil
}
