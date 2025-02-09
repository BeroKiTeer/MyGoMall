package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

var jwtSecret = []byte("SECRET_KEY")

func GenerateJWT(userID int32) (string, error) {
	// 创建一个 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                           // 用户 ID
		"exp":     time.Now().Add(time.Hour).Unix(), // 过期时间 1 小时
		"iat":     time.Now().Unix(),                // 签发时间
		"iss":     "MyGoMall-AuthService",           // 签发者
	})

	// 使用密钥对 token 进行签名
	return token.SignedString(jwtSecret)
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	token, err := GenerateJWT(req.UserId)

	return &auth.DeliveryResp{Token: token}, err
}
