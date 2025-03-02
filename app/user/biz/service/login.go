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

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	klog.Info("LoginService Run")

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		klog.Error("查询用户失败", err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		klog.Error("密码错误", err)
		return nil, err
	}
	// 调用auth服务，生成 token
	token, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: int32(row.ID),
	})

	if err != nil {
		klog.Fatal(err)
		return nil, err
	}

	resp = &user.LoginResp{
		Token: token.Token,
	}

	return resp, nil
}
