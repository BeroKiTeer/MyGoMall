package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		//Email:    "123@qq.com",
		//Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
