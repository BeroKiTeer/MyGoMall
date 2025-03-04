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
	"strconv"
	"strings"
	"time"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	// 参数检查
	if req.UserId == 0 {
		klog.Error("未输入用户id", err)
		return nil, errors.New("user id is required")
	}
	var userCart cart.Cart
	if conf.GetEnv() == "test" {
		// 看看 redis 里面有没有这个用户的购物车
		key := fmt.Sprintf("cart:%d", req.UserId)
		exist, err := redis.RedisClient.Exists(s.ctx, key).Result()

		// redis 里面有，尝试从 redis 中查找
		if exist == 1 {
			query, terr := redis.RedisClient.LRange(s.ctx, key, 0, -1).Result()
			err = terr
			for _, v := range query {
				str := strings.Split(v, ":")
				productID, terr := strconv.Atoi(str[0])
				quantity, terr := strconv.Atoi(str[1])
				if terr != nil {
					err = terr
					break
				}
				userCart.Items = append(userCart.Items, &cart.CartItem{ProductId: int64(productID), Quantity: int32(quantity)})
			}
		}

		// redis 里面没有，或者查找发生错误，从 mysql 中查找
		if exist == 0 || err != nil {
			err = nil
			userCart.UserId = req.UserId
			err = model.QueryItemsByUser(&userCart)

			if err != nil {
				return nil, errors.New("购物车数据查询时出现错误")
			}

			// 存进 redis
			for _, item := range userCart.Items {
				value := fmt.Sprintf("%d:%d", item.ProductId, item.Quantity)
				err = redis.RedisClient.LPush(s.ctx, key, value).Err()
				if err != nil {
					return nil, err
				}
			}
			err = redis.RedisClient.Expire(s.ctx, key, 3600*6*time.Second).Err()
			// 有 err 说明 redis 坏了，撤销缓存操作
			if err != nil {
				err = nil
				redis.RedisClient.Del(s.ctx, key)
			}
		}

	} else if conf.GetEnv() == "dev" {
		// 看看 redis 里面有没有这个用户的购物车
		key := fmt.Sprintf("cart:%d", req.UserId)
		exist, err := redis.RedisClusterClient.Exists(s.ctx, key).Result()

		// redis 里面有，尝试从 redis 中查找
		if exist == 1 {
			query, terr := redis.RedisClusterClient.LRange(s.ctx, key, 0, -1).Result()
			err = terr
			for _, v := range query {
				str := strings.Split(v, ":")
				productID, terr := strconv.Atoi(str[0])
				quantity, terr := strconv.Atoi(str[1])
				if terr != nil {
					err = terr
					break
				}
				userCart.Items = append(userCart.Items, &cart.CartItem{ProductId: int64(productID), Quantity: int32(quantity)})
			}
		}

		// redis 里面没有，或者查找发生错误，从 mysql 中查找
		if exist == 0 || err != nil {
			err = nil
			userCart.UserId = req.UserId
			err = model.QueryItemsByUser(&userCart)

			if err != nil {
				return nil, errors.New("购物车数据查询时出现错误")
			}

			// 存进 redis
			for _, item := range userCart.Items {
				value := fmt.Sprintf("%d:%d", item.ProductId, item.Quantity)
				err = redis.RedisClusterClient.LPush(s.ctx, key, value).Err()
				if err != nil {
					return nil, err
				}
			}
			err = redis.RedisClusterClient.Expire(s.ctx, key, 3600*6*time.Second).Err()
			// 有 err 说明 redis 坏了，撤销缓存操作
			if err != nil {
				err = nil
				redis.RedisClusterClient.Del(s.ctx, key)
			}
		}

	}

	return &cart.GetCartResp{Cart: &userCart}, nil
}
