package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func GetUserIDFromToken(token string) (int32, error) {

	// 把 token 分为好几段，其中第二段（parts[1]）是 payload
	parts := strings.Split(token, ".")
	// 按照 base64 解码
	payloadString, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, err
	}

	// 把解码的结果 转成 JSON 的 map形式 拿到 user_id
	var payloadJSON map[string]interface{}
	err = json.Unmarshal(payloadString, &payloadJSON)
	if err != nil {
		return 0, err
	}

	// 断言 user_id 是 float64，稍后强转为 string
	userID, ok := payloadJSON["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id 非法，token 无效。")
	}

	return int32(userID), nil
}
