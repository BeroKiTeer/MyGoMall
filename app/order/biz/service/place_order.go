package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"order/biz/dal/mysql"
	"order/biz/model"
	"order/rpc"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	// 1. 验证库存 （库存服务RPC调用）
	for idx, item := range req.OrderItems {
		stk, err := rpc.StockClient.CheckItem(s.ctx, &stock.CheckItemReq{
			ProductId: item.ProductId,
		})
		if err != nil {
			klog.Warn("库存服务RPC调用失败 ", err)
			return nil, err
		}
		if stk.Quantity < int64(req.OrderItems[idx].Quantity) {
			klog.Warn("库存不足，下单失败。", err)
		}
	}

	// 2. 将购物车商品移动到订单明细表中
	// 2.1 获取购物车信息 （购物车服务RPC调用）
	carts, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Warn("购物车服务RPC调用失败 ", err)
	}
	// 2.2 将购物车商品移动到订单明细表中
	// 2.2.1 启动数据库事务
	tx := mysql.DB.Begin()
	if tx.Error != nil {
		klog.Warn("数据库事务开启失败", tx.Error)
		return nil, tx.Error
	}

	// 2.2.2 生成订单UUID
	orderUUID, err := uuid.NewUUID()
	klog.Info("生成订单UUID")
	if err != nil {
		klog.Warn("生成订单ID失败", err)
		return nil, err
	}

	// 2.2.3 创建空订单表记录
	model.CreateOrder(mysql.DB, &model.Order{
		Base: model.Base{
			ID: orderUUID.String(),
		},
		UserID:          int64(req.UserId),
		TotalPrice:      0,
		DiscountPrice:   0,
		ActualPrice:     0,
		OrderStatus:     0,
		PaymentStatus:   0,
		PaymentMethod:   "",
		ShippingAddress: "",
		RecipientName:   "",
		PhoneNumber:     "",
		ShippingStatus:  0,
		PaidAt:          nil,
		ShippedAt:       nil,
		CompletedAt:     nil,
		CanceledAt:      nil,
		Remark:          nil,
	})
	// 2.2.4 创建订单明细表记录
	for _, cartItem := range carts.Cart.Items {
		err := model.CreateOrderItem(tx, &model.OrderItem{
			OrderId:   orderUUID.String(),
			ProductId: cartItem.ProductId,
			Quantity:  int64(cartItem.Quantity),
		})
		if err != nil {
			klog.Warn("创建订单明细表记录失败", err)
			return nil, err
		}
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	// TODO: 3. 清空购物车 （购物车服务RPC调用）
	emptyCartResp, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Warnf(emptyCartResp.String(), err)
		return nil, err
	}

	// TODO: 4. 计算订单价格 (checkout RPC)

	return
}
