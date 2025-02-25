package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	order_rpc "github.com/BeroKiTeer/MyGoMall/common/rpc/order"
	pdc_rpc "github.com/BeroKiTeer/MyGoMall/common/rpc/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

var (
	productClient = pdc_rpc.DefaultClient()
	orderClient   = order_rpc.DefaultClient()
)

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	//计算所购买的货物总价格
	var amount int64 = 0
	var orderItems []*order.OrderItem
	for _, val := range req.Items {
		//先验证是否有库存
		productResp, err := productClient.GetProduct(s.ctx, &product.GetProductReq{Id: val.ProductId})
		if err != nil {
			return nil, err
		}
		if int64(val.Quantity) > productResp.Product.Stock {
			return nil, errors.New("库存不足！")
		}

		//库存预留操作 ToDo

		//将该商品加入订单
		orderItems = append(orderItems, val)
		//计算金额
		amount += productResp.Product.Price * int64(val.Quantity)
	}
	//获取订单号
	placeOrderResp, err := orderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "RMB",
		Address:      req.Address,
		Email:        req.Email,
		OrderItems:   orderItems,
	},
	)
	//发送支付请求到队列

	if err != nil {
		return nil, errors.New("获取订单号失败！")
	}

	resp = &checkout.CheckoutResp{
		OrderId: placeOrderResp.Order.OrderId,
		Amount:  amount,
	}
	return resp, nil
}
