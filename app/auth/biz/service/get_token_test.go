package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"testing"
)

func TestGetToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetTokenService(ctx)
	// init req and assert value

	req := &auth.GetTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
