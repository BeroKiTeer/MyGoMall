package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"order/biz/dal/redis"
	"order/biz/model"
	order "order/kitex_gen/order"
	"order/rpc"
	"product/biz/dal/mysql"
	"product/kitex_gen/product"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	// 参数验证
	if req.OrderId == "" {
		return nil, err
	}
	// 获取订单中的商品信息
	productIds := model.GetProductIdsFromOrder(mysql.DB, req.OrderId)

	for _, productID := range productIds {
		r, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: productID})
		if err != nil {
			klog.Error("get product error: ", err)
		}
	}

	//TODO: 1. 库存减少或锁定
	for _, productID := range productIds {
		if !reduceStock(s.ctx, productID, quantity) {
			log.Println("库存不足，秒杀失败！")
			return
		}
	}

	//TODO: 2. 订单状态改为已支付
	return
}

// reduceStock lua脚本减少库存
func reduceStock(ctx context.Context, productID int64, quantity int) bool {
	luaScript := `
		local stock_key = KEYS[1]
		local amount = tonumber(ARGV[1])
		local current_stock = tonumber(redis.call('GET', stock_key))
		if current_stock >= amount then
			redis.call('INCRBY', stock_key, -amount)
			return 1
		else
			return 0
		end
	`
	res, err := redis.RedisClient.Eval(ctx, luaScript, []string{productID}, quantity).Result()
	if err != nil {
		klog.Warn("reduceStock error: ", err)
		return false
	}
	return res.(int64) == 1
}

// compensate 补偿操作：恢复库存
func compensate(ctx context.Context, productID string, quantity int) {
	// 恢复库存
	redis.RedisClient.IncrBy(ctx, productID, int64(quantity))
}
