package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"testing"
)

func TestEncodeToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDecodeTokenService(ctx)
	// init req and assert value

	req := &auth.DecodeTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
