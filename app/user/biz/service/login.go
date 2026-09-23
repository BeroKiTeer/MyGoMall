package service

import (
	"context"
	"encoding/json"
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
	var User model.User
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	//先从redis中查询
	UserCached, err := model.GetCachedUserByEmail(s.ctx, mysql.DB, req.Email)
	if err = json.Unmarshal([]byte(UserCached), &User); err != nil {
		klog.Error("用户反序列化失败:", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.PasswordHashed), []byte(req.Password))
	if err != nil {
		klog.Error("密码错误", err)
		return nil, err
	}
	// 调用auth服务，生成 token
	token, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: int32(User.ID),
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
