package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"log"
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

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	// 调用auth服务，生成 token
	token, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: int32(row.ID),
	})

	if err != nil {
		log.Fatal(err)
	}

	resp = &user.LoginResp{
		Token: token.Token,
	}

	return resp, nil
}
