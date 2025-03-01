package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/dal/mysql"
	"stock/biz/model"
	"stock/rpc"
)

type RecoverItemService struct {
	ctx context.Context
} // NewRecoverItemService new RecoverItemService
func NewRecoverItemService(ctx context.Context) *RecoverItemService {
	return &RecoverItemService{ctx: ctx}
}

// Run create note info
func (s *RecoverItemService) Run(req *stock.RecoverItemReq) (resp *stock.RecoverItemResp, err error) {

	// 根据 order_id 查所有的商品
	orderDetails, err := rpc.OrderClient.ShowOrderDetail(s.ctx, &order.ShowOrderDetailReq{OrderId: req.OrderId})
	if err != nil {
		return &stock.RecoverItemResp{Success: false}, errors.New("无法根据 order_查到对应商品！")
	}

	// 将每个商品填充到数据库中
	for _, item := range orderDetails.OrderItems {
		err = model.RecoverItem(mysql.DB, item.ProductId, int64(item.Quantity))
		if err != nil {
			return &stock.RecoverItemResp{Success: false}, err
		}
	}

	return &stock.RecoverItemResp{Success: true}, nil
}
