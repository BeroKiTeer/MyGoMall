package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"user/biz/dal/mysql"
	"user/biz/model"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Password != req.ConfirmPassword {
		klog.Error("Passwords do not match!")
		return nil, errors.New("passwords do not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	if err != nil {
		klog.Error("密码加密错误", err)
		return nil, err
	}

	err = model.CreateUser(mysql.DB, &model.User{
		PasswordHashed: string(passwordHashed),
		PhoneNumber:    "",
		Email:          req.Email,
		Status:         0,
		Role:           0,
	})
	klog.Info("注册用户:", req.Email)

	if err != nil {
		klog.Error("创建用户错误", err)
		return nil, err
	}

	return &user.RegisterResp{
		UserId: 1,
	}, err
}
