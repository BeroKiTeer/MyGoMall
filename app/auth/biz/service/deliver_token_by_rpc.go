package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"crypto/rand"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

func GenerateJWT(userID int32) (string, error) {

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	err = rdb.Set(ctx, strconv.Itoa(int(userID)), key, time.Hour).Err()

	// 创建一个 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                           // 用户 ID
		"exp":     time.Now().Add(time.Hour).Unix(), // 过期时间 1 小时
		"iat":     time.Now().Unix(),                // 签发时间
		"iss":     "MyGoMall-AuthService",           // 签发者
	})

	// 使用密钥对 token 进行签名
	return token.SignedString(key)
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	token, err := GenerateJWT(req.UserId)

	return &auth.DeliveryResp{Token: token}, err
}
