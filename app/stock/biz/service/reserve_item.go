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
	"stock/rpc"
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
	for _, item := range items.OrderItems {
		// 1. 幂等性检查: 查询Redis中是否有该订单的库存预扣信息
		productId := item.GetProductId()
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
		if err = redis.RedisClient.Set(s.ctx, key, req.Quantity, 15*time.Minute).Err(); err != nil {
			// TODO: 补偿机制(RabbitMQ)
			klog.Error("Redis 写入失败", err)
			return nil, err
		}
	}
	if err = tx.Commit().Error; err != nil {
		// TODO: 事务提交失败的补偿，回滚Redis（删除Redis新加入的键）
		klog.Error(err)
		return nil, err
	}

	return &stock.ReserveItemResp{
		Success: true,
	}, nil
}
