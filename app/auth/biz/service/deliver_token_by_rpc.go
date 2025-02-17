package service

import (
	"auth/biz/dal/redis"
	auth "auth/kitex_gen/auth"
	"context"
	"crypto/rand"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

var rdb = redis.RedisClient

func GenerateJWT(userID int32, seconds int32, ctx context.Context) (string, error) {

	// 这个变量控制 token 生效时间
	duration := time.Duration(seconds) * time.Second

	// 随机造一个密钥，这里的 32 可以改大改小，具体根据 安全性与速度的平衡 做决策
	secretKey := make([]byte, 32)
	_, err := rand.Read(secretKey)
	if err != nil {
		return "", err
	}

	if seconds == 0 {
		redis.RedisClient.Expire(ctx, strconv.Itoa(int(userID)), 0)
		return "", nil
	}

	// 把 userID 转为 string 类型 存到 key 里面，密钥是刚刚随机生成的
	err = redis.RedisClient.Set(ctx, strconv.Itoa(int(userID)), secretKey, duration).Err()
	if err != nil {
		return "", err
	}

	// 创建一个 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                          // 用户 ID
		"exp":     time.Now().Add(duration).Unix(), // 过期时间 1 小时
		"iat":     time.Now().Unix(),               // 签发时间
		"iss":     "MyGoMall-AuthService",          // 签发者
	})

	// 使用密钥对 token 进行签名
	return token.SignedString(secretKey)
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	token, err := GenerateJWT(req.UserId, 3600, s.ctx)

	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{Token: token}, nil
}
