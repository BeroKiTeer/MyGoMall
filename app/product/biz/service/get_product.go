package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"product/biz/dal/mysql"
	"product/biz/dal/redis"
	"product/biz/model"
	"strconv"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	klog.Infof("请求id：%+v\n", req.GetId())

	cachedProduct, err := redis.RedisClient.Get(s.ctx, fmt.Sprintf("product:%d", req.GetId())).Result()
	if errors.Is(err, redis.Nil) {
		// 如果缓存没有命中，从数据库获取
		p, categories, err := model.GetProductWithCategory(mysql.DB, req.GetId())
		if err != nil {
			klog.Error(err.Error())
			return nil, err
		}
		// 设置缓存，缓存过期时间为 1 小时
		cacheData := &product.GetProductResp{
			Product: &product.Product{
				Id:            p.ID,
				Name:          p.Name,
				Description:   p.Description,
				Images:        p.Images,
				Price:         p.Price,
				Categories:    categories,
				OriginalPrice: p.OriginalPrice,
				Stock:         p.Stock,
				Status:        uint32(p.Status),
			},
		}
		// 将结果缓存到 Redis
		cacheKey := req.Id // 使用产品ID作为缓存键
		// 设置缓存，设置过期时间（可以根据需求设置）
		err = redis.RedisClient.Set(s.ctx, strconv.Itoa(int(cacheKey)), cacheData, 3600).Err()
		if err != nil {
			klog.Warnf("Error setting Redis cache: %v", err)
		}
		return cacheData, nil
	} else if err != nil {
		return nil, err
	}
	// 如果缓存命中，直接返回缓存结果
	klog.Info("从 Redis 缓存中读取数据")
	var respData product.GetProductResp
	// 将 Redis 中的缓存数据反序列化到 respData
	err = json.Unmarshal([]byte(cachedProduct), &respData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached product data: %v", err)
	}

	return &respData, nil
}
