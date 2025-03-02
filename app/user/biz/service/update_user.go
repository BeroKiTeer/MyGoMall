package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"user/biz/dal/mysql"
	"user/biz/model"
	"user/rpc"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info ✅
func (s *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	// Finish your business logic.

	resp = &user.UpdateUserResp{}

	// 1️⃣ 检查 token 是否有效
	_, err = rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{
		Token: req.Token,
	})

	if err != nil {
		klog.Fatal("token无效，请重新登录，", err)
		return nil, err
	}

	// 2️⃣ 获取用户 id
	r, err := rpc.AuthClient.DecodeToken(s.ctx, &auth.DecodeTokenReq{
		Token: req.Token,
	})
	row, err := model.UserExistsByID(mysql.DB, r.UserId)
	if row == false || err != nil {
		klog.Error("用户不存在")
		return nil, errors.New("用户不存在")
	}

	// 3️⃣ 更新用户信息（检查可选字段）
	updates := map[string]interface{}{}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Password != "" {
		updates["password_hashed"], err = hashPassword(req.Password) // 假设有哈希密码的方法
	}
	if req.PhoneNumber != "" {
		updates["phone_number"] = req.PhoneNumber
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}

	// 如果没有需要更新的字段，直接返回成功
	if len(updates) == 0 {
		resp.Success = true
		return resp, nil
	}

	// 4️⃣ 执行更新操作
	if err := model.UpdateUser(mysql.DB, s.ctx, r.UserId, updates); err != nil {
		klog.Error("更新用户SQL语句失败：", err)
		return nil, err
	}

	// 5️⃣ 更新成功，返回响应
	resp.Success = true
	return resp, nil
}

// ✅ 使用 bcrypt 进行密码哈希加密
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		klog.Error("密码加密错误", err)
		return "", err
	}
	return string(hashed), nil
}
