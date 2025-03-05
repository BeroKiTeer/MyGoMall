package service

import (
	"context"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
	"stock/biz/dal/redis"
	"stock/conf"
	"stock/rpc"
)

type ReduceItemService struct {
	ctx context.Context
} // NewReduceItemService new ReduceItemService
func NewReduceItemService(ctx context.Context) *ReduceItemService {
	return &ReduceItemService{ctx: ctx}
}

// Run create note info
func (s *ReduceItemService) Run(req *stock.ReduceItemReq) (resp *stock.ReduceItemResp, err error) {

	// TODO: 此处应为消息队列处理分布式事务，暂时使用rpc调用
	items, err := rpc.OrderClient.ShowOrderDetail(s.ctx, &order.ShowOrderDetailReq{OrderId: req.OrderId})
	if err != nil {
		klog.Errorf("show order detail failed, err: %v", err)
		return &stock.ReduceItemResp{Success: false}, err
	}
	for _, item := range items.OrderItems {
		key := fmt.Sprintf("predestock:%s:%d", req.GetOrderId(), item.ProductId)
		if conf.GetEnv() == "test" {
			err = redis.RedisClient.Del(s.ctx, key).Err()
			if err != nil {
				klog.Errorf("del predestock failed, err: %v", err)
				return &stock.ReduceItemResp{Success: false}, err
			}
		} else if conf.GetEnv() == "dev" {
			err = redis.RedisClusterClient.Del(s.ctx, key).Err()
			if err != nil {
				klog.Errorf("del predestock failed, err: %v", err)
				return &stock.ReduceItemResp{Success: false}, err
			}
		}
	}

	return &stock.ReduceItemResp{Success: true}, nil
}
