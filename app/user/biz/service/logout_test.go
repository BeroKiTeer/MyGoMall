package service

import (
	"context"
	"testing"
	"user/biz/dal/redis"
	user "user/kitex_gen/user"
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
