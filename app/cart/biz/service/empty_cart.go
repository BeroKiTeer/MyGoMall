package service

import (
	"cart/biz/dal/redis"
	"cart/biz/model"
	"cart/conf"
	"context"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {

	// 参数检查
	if req.UserId == 0 {
		klog.Error("未输入用户id", err)
		return nil, errors.New("empty user id")
	}

	// 检查商品是否已存在在购物车
	var targetItemQuantity int32 = -1
	err = model.CheckItemsByUser(req.UserId, &targetItemQuantity)

	if err != nil {
		klog.Error("未查询到商品", err)
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

	// 删除
	if targetItemQuantity != -1 {
		klog.Info("清空购物车接口响应成功")
		err = model.EmptyCart(req.UserId)
	}
	return &cart.EmptyCartResp{}, err
}
