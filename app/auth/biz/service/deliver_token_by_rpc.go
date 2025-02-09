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

// 这里默认 redis 地址是 localhost:6379。未来可能需要更改
var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

func GenerateJWT(userID int32) (string, error) {

	secretKey := make([]byte, 32)
	// 随机造一个密钥
	_, err := rand.Read(secretKey)
	if err != nil {
		return "", err
	}

	// 把 userID 转为 string 类型 存到 key 里面，密钥是刚刚随机生成的
	err = rdb.Set(ctx, strconv.Itoa(int(userID)), secretKey, time.Hour).Err()
	if err != nil {
		return "", err
	}

	// 创建一个 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                           // 用户 ID
		"exp":     time.Now().Add(time.Hour).Unix(), // 过期时间 1 小时
		"iat":     time.Now().Unix(),                // 签发时间
		"iss":     "MyGoMall-AuthService",           // 签发者
	})

	// 使用密钥对 token 进行签名
	return token.SignedString(secretKey)
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	token, err := GenerateJWT(req.UserId)

	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{Token: token}, nil
}
