package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"testing"
)

func TestEncodeToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDecodeTokenService(ctx)
	// init req and assert value

	req := &auth.DecodeTokenReq{Token: ""}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
