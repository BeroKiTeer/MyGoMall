package service

import (
	"cart/biz/dal/redis"
	"cart/biz/model"
	"context"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"strconv"
	"strings"
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
		return nil, errors.New("user id is required")
	}

	// 看看 redis 里面有没有这个用户的购物车
	key := fmt.Sprintf("cart:%d", req.UserId)
	exist, err := redis.RedisClient.Exists(s.ctx, key).Result()
	var userCart cart.Cart

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
		userCart.UserId = req.UserId
		err = model.QueryItemsByUser(&userCart)
		// 存进 redis
		for _, item := range userCart.Items {
			value := fmt.Sprintf("%d:%d", item.ProductId, item.Quantity)
			err = redis.RedisClient.LPush(s.ctx, key, value).Err()
		}
		err = redis.RedisClient.Expire(s.ctx, key, 3600*6).Err()
	}

	return &cart.GetCartResp{Cart: &userCart}, err
}
