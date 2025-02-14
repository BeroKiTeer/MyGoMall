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

	req := &auth.DecodeTokenReq{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NTY3ODkwLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.FPaOaNsXYOFHMSfWpKxI456V0znpJawcvo2UJOZG9mA"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
