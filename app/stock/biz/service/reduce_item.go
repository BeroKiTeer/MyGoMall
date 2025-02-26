package service

import (
	"context"
	"fmt"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/dal/redis"
)

type ReduceItemService struct {
	ctx context.Context
} // NewReduceItemService new ReduceItemService
func NewReduceItemService(ctx context.Context) *ReduceItemService {
	return &ReduceItemService{ctx: ctx}
}

// Run create note info
func (s *ReduceItemService) Run(req *stock.ReduceItemReq) (resp *stock.ReduceItemResp, err error) {

	key := fmt.Sprintf("predestock:%s:%d", req.GetOrderId(), req.GetProductId())
	err = redis.RedisClient.Del(s.ctx, key).Err()
	if err != nil {
		return &stock.ReduceItemResp{Success: false}, err
	}

	return &stock.ReduceItemResp{Success: true}, nil
}
