package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"testing"
	"user/biz/dal/redis"
)

func TestDeleteUser_Run(t *testing.T) {
	redis.Init()
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	req := &user.DeleteUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
