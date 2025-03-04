package service

import (
	mq "checkout/biz/dal/RabbitMQ"
	order_rpc "checkout/rpc/order"
	pdc_rpc "checkout/rpc/product"
	stock_rpc "checkout/rpc/stock"
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
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
	//获取订单号
	klog.Info("checkout:", req.UserId)
	placeOrderResp, err := orderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "RMB",
		Address:      req.Address,
		Email:        req.Email,
		OrderItems:   orderItems,
	},
	)
	if err != nil {
		klog.Error("获取订单号失败", err)
		return nil, errors.New("获取订单号失败！")
	}

	for _, val := range req.Items {
		//先验证是否有库存
		productResp, err := productClient.GetProduct(s.ctx, &product.GetProductReq{Id: val.ProductId})
		if err != nil {
			klog.Error("GetProductReq接口无响应，无法查询到库存", err)
			return nil, err
		}
		if int64(val.Quantity) > productResp.Product.Stock {
			klog.Info("GetProductReq接口相应成功，查询到库存")
			return nil, errors.New("库存不足！")
		}

		//库存预留操作
		_, err = stock_rpc.DefaultClient().ReserveItem(s.ctx, &stock.ReserveItemReq{
			OrderId: resp.OrderId,
		})
		if err != nil {
			klog.Error("ReserveItemReq接口无响应，无法预留库存", err)
			return nil, err
		}

		//将该商品加入订单
		orderItems = append(orderItems, val)
		//计算金额
		amount += productResp.Product.Price * int64(val.Quantity)
	}

	//发送支付请求到队列
	switch req.PaymentMethod {
	case "credit_card":
		{
			cardPaymentReq := &mq.CardPayment{
				OrderID:     placeOrderResp.Order.OrderId,
				Amount:      amount,
				CallbackURL: "...", //todo,关于url的生成仍未实现
			}
			err := mq.CardPaymentProducer.Send(cardPaymentReq)
			if err != nil {
				klog.Error("cardPaymentReq接口未响应，无法发送支付请求", err)
				return nil, err
			}
		}
	}

	resp = &checkout.CheckoutResp{
		OrderId: placeOrderResp.Order.OrderId,
		Amount:  amount,
	}
	return resp, nil
}
