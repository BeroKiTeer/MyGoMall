package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
	"stock/biz/model"
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
	// TODO: 查询库存是否充足
	quantity, err := model.CheckQuantity(req.ProductId)
	if quantity < req.Quantity {
		klog.Error(err)
		return nil, err
	}
	// TODO: 预扣库存
	err = model.ReduceItem(req.ProductId, req.Quantity)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	// TODO: 扣减的放到Redis里
	//res, err := redis.RedisClient.
	//	Get(s.ctx, fmt.Sprintf("predestock:%s:%d:", req.GetOrderId(), req.GetProductId())).
	//	Result()
	//if errors.Is(err, redis.Nil) {
	//	return res, err
	//}
	//if err != nil {
	//	klog.Error(err)
	//	return nil, err
	//}

	return
}
