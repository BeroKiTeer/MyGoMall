package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"testing"
	"user/biz/dal/redis"
)

func TestLogout_Run(t *testing.T) {
	redis.Init()
	ctx := context.Background()
	s := NewLogoutService(ctx)
	// init req and assert value

	req := &user.LogoutReq{UserId: 1234}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
