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
		return nil, errors.New("Passwords do not match!")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	if err != nil {
		klog.Error("密码加密错误", err)
		return nil, err
	}

	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	err = model.CreateUser(mysql.DB, newUser)

	if err != nil {
		klog.Error("创建用户错误", err)
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, err
}
