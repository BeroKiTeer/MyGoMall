package service

import (
	"cart/biz/model"
	cart "cart/kitex_gen/cart"
	"cart/rpc"
	"context"
	"errors"
	"product/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {

	// 参数验证
	if req.UserId == 0 || req.Item == nil {
		return nil, errors.New("need user_id and item")
	}

	if req.Item.ProductId == 0 {
		return nil, errors.New("product_id is 0")
	}

	if req.Item.Quantity == 0 {
		return nil, errors.New("quantity is 0")
	}

	// 检查商品是否存在（RPC）
	ProductReq := product.GetProductReq{Id: req.Item.ProductId}
	productDetails, err := rpc.ProductClient.GetProduct(s.ctx, &ProductReq)
	if err != nil || productDetails == nil {
		return nil, errors.New("the product does not exist")
	}

	// 检查商品库存是否足够
	if int32(productDetails.Product.Stock) < req.Item.Quantity {
		return nil, errors.New("the stock is not enough")
	}

	// 检查商品是否已存在在购物车
	var targetItemQuantity int32 = -1
	model.CheckItemsByUserAndProduct(req.UserId, req.Item.ProductId, &targetItemQuantity)

	// 将商品添加到购物车，持久化存储。（如果已存在，则修改原有的）
	if targetItemQuantity == -1 {
		model.AddItem(req.UserId, req.Item.ProductId, req.Item.Quantity)
	} else {
		model.UpdateItem(req.UserId, req.Item.ProductId, req.Item.Quantity)
	}

	return
}
