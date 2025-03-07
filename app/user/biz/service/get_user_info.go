package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
	"user/biz/dal/mysql"
	"user/biz/dal/redis"
	"user/biz/model"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	// Finish your business logic.
	//查询用户是否存在
	var User model.User
	var addressStr string
	cachedUser, err := redis.RedisClusterClient.Get(s.ctx, fmt.Sprintf("UserId:%d", req.UserId)).Result()
	if errors.Is(err, redis.Nil) {
		// 如果缓存没有命中，从数据库获取
		row, err := model.GetUserById(mysql.DB, s.ctx, req.UserId)
		if err != nil {
			klog.Error(err)
			return nil, err
		}

		/**
		string address = 5;      // 地址
		  string role = 6;         // 角色（如"admin", "user"）
		  string status = 7;       // 账户状态（如"active", "banned", "pending"）
		  string created_at = 8;   // 账户创建时间
		  string updated_at = 9;   // 账户最近更新时间
		*/

		//将查询到的数据存入缓存
		cacheData := row
		//序列化后存储到redis里
		UserJson, _ := json.Marshal(cacheData)
		if err := redis.RedisClusterClient.Set(
			s.ctx,
			fmt.Sprintf("UserId:%d", row.ID),
			UserJson,
			1*time.Hour).Err(); err != nil {
			klog.Error("用户缓存失败:", err)
		}
		cachedUser = string(UserJson)
	} else if err != nil {
		klog.Error(err)
	}
	if err := json.Unmarshal([]byte(cachedUser), &User); err != nil {
		klog.Error("用户反序列化失败:", err)
		return nil, err
	}
	//address为nil怎么办
	if User.AddressId != 0 {
		addressStr, err := redis.RedisClusterClient.Get(s.ctx, fmt.Sprintf("AddressId:%d", User.AddressId)).Result()
		//若redis中不存在，则从mysql中获取
		if errors.Is(err, redis.Nil) {
			//从数据库中查询
			address, err := model.GetAddressById(mysql.DB, s.ctx, User.AddressId)
			if err != nil {
				klog.Error("查询失败:", err)
			}
			if address == nil {
				address = &model.Address{
					Address: "",
				}
			}
			// 序列化为JSON字符串存储(存储结构体时的策略)
			addressJSON, _ := json.Marshal(address)
			addressStr = string(addressJSON)
			if err := redis.RedisClusterClient.Set(
				s.ctx,
				fmt.Sprintf("AddressId:%d", User.AddressId),
				addressJSON, // 存储序列化后的JSON
				24*time.Hour).Err(); err != nil {
				klog.Error("地址缓存失败:", err)
			}
		} else if err != nil {
			klog.Error(err)
			return nil, err
		}
		var cachedAddress model.Address
		if err := json.Unmarshal([]byte(addressStr), &cachedAddress); err != nil {
			klog.Error("地址反序列化失败:", err)
			return nil, err
		}
	}

	// 0-普通用户, 1-管理员
	if User.Role == 0 {
		resp.Role = "user"
	} else {
		resp.Role = "admin"
	}
	// 0-正常, 1-禁用, 2-待审核
	if User.Status == 0 {
		resp.Status = "active"
	} else if User.Status == 1 {
		resp.Status = "banned"
	} else if User.Status == 2 {
		resp.Status = "pending"
	}
	resp = &user.GetUserInfoResp{
		UserId:      int32(User.ID),
		Email:       User.Email,
		Username:    User.Username,
		PhoneNumber: User.PhoneNumber,
		Address:     addressStr,
		Role:        strconv.Itoa(int(User.Role)),
		Status:      strconv.Itoa(int(User.Status)),
		CreatedAt:   User.CreatedAt.String(),
		UpdatedAt:   User.UpdatedAt.String(),
	}
	return resp, err
}
