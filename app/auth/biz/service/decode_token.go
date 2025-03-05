package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

type DecodeTokenService struct {
	ctx context.Context
} // NewDecodeTokenService new EncodeTokenService
func NewDecodeTokenService(ctx context.Context) *DecodeTokenService {
	return &DecodeTokenService{ctx: ctx}
}

func GetUserIDFromToken(token string) (int32, error) {

	// 把 token 分为好几段，其中第二段（parts[1]）是 payload
	parts := strings.Split(token, ".")
	// 按照 base64 解码
	payloadString, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		klog.Error("解码错误", err)
		return 0, err
	}

	// 把解码的结果 转成 JSON 的 map形式 拿到 user_id
	var payloadJSON map[string]interface{}
	err = json.Unmarshal(payloadString, &payloadJSON)
	if err != nil {
		klog.Error(err)
		return 0, err
	}

	// 断言 user_id 是 float64，稍后强转为 string
	userID, ok := payloadJSON["user_id"].(float64)
	if !ok {
		klog.Error("userid = ", userID, "错误：", err)
		return 0, fmt.Errorf("user_id 非法，token 无效。")
	}

	return int32(userID), nil
}

// Run create note info
func (s *DecodeTokenService) Run(req *auth.DecodeTokenReq) (resp *auth.DecodeTokenResp, err error) {

	UserID, err := GetUserIDFromToken(req.Token)
	if err != nil {
		klog.Error(err)
		return &auth.DecodeTokenResp{}, err
	}
	return &auth.DecodeTokenResp{UserId: UserID}, err
}
