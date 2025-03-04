package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
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
	row, err := model.GetUserById(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return
	}
	/**
	string address = 5;      // 地址
	  string role = 6;         // 角色（如"admin", "user"）
	  string status = 7;       // 账户状态（如"active", "banned", "pending"）
	  string created_at = 8;   // 账户创建时间
	  string updated_at = 9;   // 账户最近更新时间
	*/
	var address *model.Address
	if row.AddressId != 0 {
		address, err = model.GetAddressById(mysql.DB, s.ctx, row.AddressId)
		if err != nil {
			klog.Error(err)
			return nil, err
		}
	}
	if address == nil {
		address = &model.Address{
			Address: "",
		}
	}
	resp = &user.GetUserInfoResp{
		UserId:      int32(row.ID),
		Email:       row.Email,
		Username:    row.Username,
		PhoneNumber: row.PhoneNumber,
		Address:     address.Address,
		CreatedAt:   row.CreatedAt.Format(time.DateTime),
		UpdatedAt:   row.UpdatedAt.Format(time.DateTime),
	}
	// 0-普通用户, 1-管理员
	if row.Role == 0 {
		resp.Role = "user"
	} else {
		resp.Role = "admin"
	}
	// 0-正常, 1-禁用, 2-待审核
	if row.Status == 0 {
		resp.Status = "active"
	} else if row.Status == 1 {
		resp.Status = "banned"
	} else if row.Status == 2 {
		resp.Status = "pending"
	}
	return resp, err
}
