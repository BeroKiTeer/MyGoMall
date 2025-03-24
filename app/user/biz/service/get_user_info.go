package service

import (
	"context"
	"encoding/json"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"user/biz/dal/mysql"
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
	//根据userid查询
	cachedUser, err := model.GetCachedUserById(s.ctx, mysql.DB, req.UserId)
	if err := json.Unmarshal([]byte(cachedUser), &User); err != nil {
		klog.Error("用户反序列化失败:", err)
		return nil, err
	}
	//address为nil怎么办
	if User.AddressId != 0 {
		addressStr, err = model.GetCachedAddressById(s.ctx, mysql.DB, User.AddressId)
		if err != nil {
			klog.Error(err)
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
