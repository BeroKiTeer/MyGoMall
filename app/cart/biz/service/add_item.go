package service

import (
	"cart/biz/dal/redis"
	"cart/biz/model"
	"cart/conf"
	"cart/rpc"
	"context"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {

	log.Println("AddItemService Run")
	// 参数验证
	if req.UserId == 0 || req.Item == nil {
		klog.Error("未输入用户id和商品信息", err)
		return nil, errors.New("need user_id and item")
	}

	if req.Item.ProductId == 0 {
		klog.Error("未输入商品信息", err)
		return nil, errors.New("product_id is 0")
	}

	if req.Item.Quantity == 0 {
		klog.Error("未输入商品数量", err)
		return nil, errors.New("quantity is 0")
	}

	// 检查商品是否存在（RPC）
	ProductReq := product.GetProductReq{Id: req.Item.ProductId}

	productDetails, err := rpc.ProductClient.GetProduct(s.ctx, &ProductReq)
	if err != nil || productDetails == nil {
		klog.Error("商品不存在", err)
		log.Println("商品不存在", err)
		return nil, errors.New("the product does not exist")
	}

	// 检查商品库存是否足够
	if int32(productDetails.Product.Stock) < req.Item.Quantity {
		klog.Error("商品数量不足", err)
		log.Println("商品数量不足", err)
		return nil, errors.New("the stock is not enough")
	}
	// 检查商品是否已存在在购物车
	var targetItemQuantity int32 = -1
	err = model.CheckItemsByUserAndProduct(req.UserId, req.Item.ProductId, &targetItemQuantity)

	if err != nil {
		klog.Error("在购物车里未查询到商品", err)
		log.Println(err)
		return nil, err
	}

	// 清除 redis 中该用户的商品的缓存。
	key := fmt.Sprintf("cart:%d", req.UserId)
	if conf.GetEnv() == "test" {
		err = redis.RedisClient.Del(s.ctx, key).Err()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else if conf.GetEnv() == "dev" {
		err = redis.RedisClusterClient.Del(s.ctx, key).Err()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	// 将商品添加到购物车，持久化存储。（如果已存在，则修改原有的）
	if targetItemQuantity == -1 {
		klog.Info("添加购物车接口相应，添加到购物车成功")
		err = model.AddItem(req.UserId, req.Item.ProductId, req.Item.Quantity)
	} else {
		klog.Info("商品已存在，更新购物车成功")
		err = model.UpdateItem(req.UserId, req.Item.ProductId, req.Item.Quantity)
	}

	return &cart.AddItemResp{}, err
}
