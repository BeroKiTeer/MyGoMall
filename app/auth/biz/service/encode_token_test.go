package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"testing"
)

func TestEncodeToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEncodeTokenService(ctx)
	// init req and assert value

	req := &auth.EncodeTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
