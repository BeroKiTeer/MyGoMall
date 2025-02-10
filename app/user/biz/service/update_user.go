package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/biz/dal/mysql"
	"user/biz/model"
	user "user/kitex_gen/user"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	// Finish your business logic.

	resp = &user.UpdateUserResp{}

	// 1️⃣ 检查 user_id 是否有效
	if req.UserId <= 0 {
		return nil, errors.New("无效的用户 ID")
	}

	// 2️⃣ 查询用户是否存在
	row, err := model.UserExistsByID(mysql.DB, req.UserId)
	if row == false || err != nil {
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

	// TODO:4️⃣ 执行更新操作
	if err := model.UpdateUser(mysql.DB, s.ctx, req.UserId, updates); err != nil {
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
		return "", err
	}
	return string(hashed), nil
}
