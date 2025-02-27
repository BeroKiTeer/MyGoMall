package service

import (
	"context"
	"errors"
	"fmt"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
	"stock/biz/dal/mysql"
	"stock/biz/dal/redis"
	"stock/biz/model"
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
	// 1. 幂等性检查: 查询Redis中是否有该订单的库存预扣信息
	key := fmt.Sprintf("predestock:%s:%d", req.GetOrderId(), req.GetProductId())
	exists, err := redis.RedisClient.Exists(s.ctx, key).Result()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if exists == 1 {
		return nil, errors.New("库存已预扣")
	}
	// 2. 查询库存是否充足
	tx := mysql.DB.Begin()
	quantity, err := model.CheckQuantity(tx, req.ProductId)
	if err != nil {
		klog.Error(err)
		tx.Rollback()
		return nil, err
	}
	if quantity < req.Quantity {
		klog.Error(err)
		tx.Rollback()
		return nil, err
	}
	// 3. 预扣库存，首先数据库中扣减库存 TODO: SQL 待修改
	if err = model.ReduceItem(tx, req.ProductId, req.Quantity); err != nil {
		klog.Error(err)
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		klog.Error(err)
		return nil, err
	}
	//4. 扣减的放到Redis里
	if err = redis.RedisClient.Set(s.ctx, key, req.Quantity, 15*time.Minute).Err(); err != nil {
		// TODO: 补偿机制(RabbitMQ)
		klog.Error("Redis 写入失败", err)
		return nil, err
	}

	return &stock.ReserveItemResp{
		Success: true,
	}, nil
}
