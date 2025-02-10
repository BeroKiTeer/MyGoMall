package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {

	// 把 token 分为好几段，其中第二段（parts[1]）是 payload
	parts := strings.Split(req.Token, ".")
	// 按照 base64 解码
	payloadString, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return &auth.VerifyResp{Res: false}, err
	}

	// 把解码的结果 转成 JSON 的 map形式 拿到 user_id
	var payloadJSON map[string]interface{}
	err = json.Unmarshal(payloadString, &payloadJSON)
	if err != nil {
		return &auth.VerifyResp{Res: false}, err
	}

	// 断言 user_id 是 float64，稍后强转为 string
	userID, ok := payloadJSON["user_id"].(float64)
	if !ok {
		return &auth.VerifyResp{Res: false}, fmt.Errorf("user_id 非法，token 无效。")
	}

	// 查询 userID 在 redis 里面是否存在，如果存在，获取密钥并验证
	secretKey, err := rdb.Get(ctx, strconv.Itoa(int(userID))).Result()

	if err != nil {
		return &auth.VerifyResp{Res: false}, err
	}

	// 按照查询到的密钥，解析并验证 Token
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}
		// 密钥只支持 []byte 类型，强转。
		return []byte(secretKey), nil
	})

	if err != nil {
		return &auth.VerifyResp{Res: false}, err
	}

	// 断言解析结果，返回是否对得上
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &auth.VerifyResp{Res: true}, nil
	} else {
		return &auth.VerifyResp{Res: false}, nil
	}

}
